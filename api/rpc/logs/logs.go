package logs

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/appcelerator/amp/api/rpc/cluster/constants"
	"github.com/appcelerator/amp/pkg/elasticsearch"
	"github.com/appcelerator/amp/pkg/nats-streaming"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/olivere/elastic.v5"
)

const (
	NumberOfEntries = 100
)

// Server is used to implement log.LogServer
type Server struct {
	ES *elasticsearch.Elasticsearch
	NS *ns.NatsStreaming
}

// Get implements logs.LogsServer
func (s *Server) LogsGet(ctx context.Context, in *GetRequest) (*GetReply, error) {
	if err := s.ES.Connect(); err != nil {
		return nil, errors.New("unable to connect to elasticsearch service")
	}
	log.Infoln("rpc-logs: Get", in.String())

	// Prepares indices
	indices := []string{}
	date := time.Now().UTC()
	since := int(in.Since)
	if since < 2 {
		since = 2
	}
	for i := 0; i < since; i++ {
		indices = append(indices, "ampbeat-"+date.Format("2006.01.02"))
		date = date.AddDate(0, 0, -1)
	}

	// Prepare request to elasticsearch
	request := s.ES.GetClient().Search().Index(indices...).IgnoreUnavailable(true)
	request.Type("logs")
	request.Sort("time_id", false)
	request.Size(NumberOfEntries)
	if in.Size != 0 {
		request.Size(int(in.Size))
	}
	// Pagination
	if in.From != "" {
		request.SearchAfter(in.From)
	}

	masterQuery := elastic.NewBoolQuery()
	if in.Container != "" {
		boolQuery := elastic.NewBoolQuery()
		masterQuery.Filter(
			boolQuery.Should(elastic.NewPrefixQuery("container_id", in.Container)),
			boolQuery.Should(elastic.NewPrefixQuery("container_name", in.Container)),
		)
	}
	if in.Service != "" {
		boolQuery := elastic.NewBoolQuery()
		masterQuery.Filter(
			boolQuery.Should(elastic.NewPrefixQuery("service_id", in.Service)),
			boolQuery.Should(elastic.NewPrefixQuery("service_name", in.Service)),
		)
	}
	if in.Stack != "" {
		boolQuery := elastic.NewBoolQuery()
		masterQuery.Filter(
			boolQuery.Should(elastic.NewPrefixQuery("stack_id", in.Stack)),
			boolQuery.Should(elastic.NewPrefixQuery("stack_name", in.Stack)),
		)
	}
	if in.Task != "" {
		masterQuery.Filter(elastic.NewPrefixQuery("task_id", in.Task))
	}
	if in.Node != "" {
		masterQuery.Filter(elastic.NewPrefixQuery("node_id", in.Node))
	}
	if in.Message != "" {
		if in.Regexp {
			masterQuery.Filter(elastic.NewRegexpQuery("msg", in.Message))
		} else {
			queryString := elastic.NewSimpleQueryStringQuery(in.Message)
			queryString.Field("msg")
			masterQuery.Filter(queryString)
		}
	}
	if !in.IncludeAmpLogs {
		masterQuery.MustNot(elastic.NewExistsQuery(dockerToEsLabel(constants.LabelKeyRole)))
	}

	// Perform ES request
	searchResult, err := request.Query(masterQuery).Do(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, elasticsearch.FormatError(err))
	}

	// Build reply
	size := len(searchResult.Hits.Hits)
	reply := GetReply{
		Size:  int32(size),
		Total: int32(searchResult.Hits.TotalHits),
	}
	reply.Entries = make([]*LogEntry, size)
	for i, hit := range searchResult.Hits.Hits {
		entry := &LogEntry{}
		if err := s.unmarshal(*hit.Source, entry); err != nil {
			return nil, status.Errorf(codes.Internal, "%v", err)
		}
		// Pagination
		if len(hit.Sort) > 0 {
			if fromID, ok := hit.Sort[0].(string); ok {
				entry.FromId = fromID
			} else {
				log.Warnln("rpc-logs: Unable to cast fromID to string. Got value:", hit.Sort[0])
			}
		}
		reply.Entries[i] = entry

		// Convert ES labels to Docker labels
		labels := make(map[string]string)
		for k, v := range entry.Labels {
			labels[esToDockerLabel(k)] = v
		}
		reply.Entries[i].Labels = labels
	}

	// Reverse entries
	for i, j := 0, len(reply.Entries)-1; i < j; i, j = i+1, j-1 {
		reply.Entries[i], reply.Entries[j] = reply.Entries[j], reply.Entries[i]
	}
	log.Infof("rpc-logs: Get successful, returned %d entries\n", len(reply.Entries))
	return &reply, nil
}

// custom unmarshal for @timestamp
func (s *Server) unmarshal(data []byte, entry *LogEntry) error {
	type Alias LogEntry
	aux := &struct {
		TimestampTmp string `json:"@timestamp"`
		*Alias
	}{
		Alias: (*Alias)(entry),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	entry.Timestamp = aux.TimestampTmp
	return nil
}

// GetStream implements log.LogServer
func (s *Server) LogsGetStream(in *GetRequest, stream Logs_LogsGetStreamServer) error {
	if err := s.NS.Connect(); err != nil {
		return errors.New("unable to connect to nats service")
	}
	log.Infoln("rpc-logs: GetStream", in.String())

	sub, err := s.NS.GetClient().Subscribe(ns.LogsSubject, func(msg *stan.Msg) {
		entries := &GetReply{}
		if err := proto.Unmarshal(msg.Data, entries); err != nil {
			log.Errorln("error unmarshalling message", err)
			return
		}
		for _, entry := range entries.Entries {
			if match(entry, in) {
				stream.Send(entry)
			}
		}
	})
	if err != nil {
		return status.Errorf(codes.Internal, "%v", err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		}
	}
}

func match(entry *LogEntry, in *GetRequest) bool {
	match := true
	if in.Container != "" {
		prefix := strings.ToLower(in.Container)
		containerID := strings.ToLower(entry.ContainerId)
		containerName := strings.ToLower(entry.ContainerName)
		match = match && (strings.HasPrefix(containerID, prefix) || strings.HasPrefix(containerName, prefix))
	}
	if in.Service != "" {
		prefix := strings.ToLower(in.Service)
		serviceID := strings.ToLower(entry.ServiceId)
		serviceName := strings.ToLower(entry.ServiceName)
		match = match && (strings.HasPrefix(serviceID, prefix) || strings.HasPrefix(serviceName, prefix))
	}
	if in.Stack != "" {
		prefix := strings.ToLower(in.Stack)
		stackID := strings.ToLower(entry.StackId)
		stackName := strings.ToLower(entry.StackName)
		match = match && (strings.HasPrefix(stackID, prefix) || strings.HasPrefix(stackName, prefix))
	}
	if in.Task != "" {
		match = match && strings.HasPrefix(strings.ToLower(entry.TaskId), strings.ToLower(in.Task))
	}
	if in.Node != "" {
		match = match && strings.HasPrefix(strings.ToLower(entry.NodeId), strings.ToLower(in.Node))
	}
	if in.Message != "" {
		if in.Regexp {
			matched, _ := regexp.MatchString(in.Message, entry.Msg)
			match = match && matched
		} else {
			match = match && strings.Contains(strings.ToLower(entry.Msg), strings.ToLower(in.Message))
		}

	}
	if !in.IncludeAmpLogs {
		_, ampLogs := entry.Labels[constants.LabelKeyRole]
		match = match && !ampLogs
	}
	return match
}

func dockerToEsLabel(name string) string { return "labels." + strings.Replace(name, ".", "-", -1) }
func esToDockerLabel(name string) string { return strings.Replace(name, "-", ".", -1) }
