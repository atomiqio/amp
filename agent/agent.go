package core

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/appcelerator/amp/api/rpc/logs"
	"github.com/appcelerator/amp/api/rpc/stats"
	"github.com/appcelerator/amp/pkg/docker"
	"github.com/appcelerator/amp/pkg/nats-streaming"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

const (
	containersDataDir = "/containers"
)

// Agent data
type Agent struct {
	dock                *docker.Docker
	containers          map[string]*ContainerData
	eventStreamReading  bool
	logsSavedDatePeriod int
	natsStreaming       *ns.NatsStreaming
	nbLogs              int
	nbMetrics           int
	metricsBuffer       *stats.StatsReply
	metricsBufferMutex  *sync.Mutex
	logsBuffer          *logs.GetReply
	logsBufferMutex     *sync.Mutex
}

// AgentInit Connect to docker engine, get initial containers list and start the agent
func AgentInit(version, build string) error {
	agent := Agent{
		logsSavedDatePeriod: 60,
		metricsBuffer:       &stats.StatsReply{},
		metricsBufferMutex:  &sync.Mutex{},
		logsBuffer:          &logs.GetReply{},
		logsBufferMutex:     &sync.Mutex{},
	}
	agent.trapSignal()
	conf.init(version, build)

	// containers dir creation
	if err := os.MkdirAll(containersDataDir, 0666); err != nil {
		return fmt.Errorf("Unable to create container data directory: %s", err)
	}

	// NATS Connect
	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("Unable to get hostname: %s", err)
	}
	agent.natsStreaming = ns.NewClient(ns.DefaultURL, ns.ClusterID, os.Args[0]+"-"+hostname, time.Minute)
	if err = agent.natsStreaming.Connect(); err != nil {
		return err
	}

	// Connection to Docker
	agent.dock = docker.NewClient(conf.dockerEngine, docker.DefaultVersion)
	if err = agent.dock.Connect(); err != nil {
		_ = agent.natsStreaming.Close()
		return err
	}
	log.Println("Connected to Docker-engine")

	log.Println("Extracting containers list...")
	agent.containers = make(map[string]*ContainerData)
	ContainerListOptions := types.ContainerListOptions{All: true}
	containers, err := agent.dock.GetClient().ContainerList(context.Background(), ContainerListOptions)
	if err != nil {
		_ = agent.natsStreaming.Close()
		return err
	}
	for _, cont := range containers {
		agent.addContainer(cont.ID)
	}
	log.Println("done")
	agent.start()
	return nil
}

// Main agent loop, verify if events and logs stream are started if not start them
func (a *Agent) start() {
	a.initAPI()
	nb := 0
	a.startMetricsBufferSender()
	a.startLogsBufferSender()
	for {
		a.updateStreams()
		nb++
		if nb == 10 {
			log.Printf("Sent %d logs and %d metrics on the last %d seconds\n", a.nbLogs, a.nbMetrics, nb*conf.period)
			nb = 0
			a.nbLogs = 0
			a.nbMetrics = 0
		}
		time.Sleep(time.Duration(conf.period) * time.Second)
	}
}

func (a *Agent) startMetricsBufferSender() {
	if conf.metricsBufferPeriod == 0 || conf.metricsBufferSize == 0 {
		a.metricsBuffer.Entries = make([]*stats.MetricsEntry, 1)
		return
	}
	go func() {
		time.Sleep(time.Second * time.Duration(conf.metricsBufferPeriod))
		if len(a.metricsBuffer.Entries) > 0 {
			a.metricsBufferMutex.Lock()
			a.sendMetricsBuffer()
			a.metricsBufferMutex.Unlock()
		}
	}()
}

func (a *Agent) startLogsBufferSender() {
	if conf.logsBufferPeriod == 0 || conf.logsBufferSize == 0 {
		a.logsBuffer.Entries = make([]*logs.LogEntry, 1)
		return
	}
	go func() {
		time.Sleep(time.Second * time.Duration(conf.logsBufferPeriod))
		if len(a.logsBuffer.Entries) > 0 {
			a.logsBufferMutex.Lock()
			a.sendLogsBuffer()
			a.logsBufferMutex.Unlock()
		}
	}()
}

// Starts logs and metrics stream of eech new started container
func (a *Agent) updateStreams() {
	a.updateLogsStream()
	a.updateMetricsStreams()
	a.updateEventsStream()
}

// Close AgentInit resources
func (a *Agent) stop() {
	a.closeLogsStreams()
	a.closeMetricsStreams()
	err := a.dock.GetClient().Close()
	if err != nil {
		log.Println("error closing connection to docker client: ", err)
	}
	err = a.natsStreaming.Close()
	if err != nil {
		log.Println("error closing connection to NATS: ", err)
	}
}

// Launch a routine to catch SIGTERM Signal
func (a *Agent) trapSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		<-ch
		log.Println("\nagent received SIGTERM signal")
		a.stop()
		os.Exit(1)
	}()
}
