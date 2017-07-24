package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/appcelerator/amp/cluster/agent/admin"
	"github.com/appcelerator/amp/cluster/agent/pkg/docker"
	"github.com/appcelerator/amp/cluster/agent/pkg/docker/stack"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/compose/convert"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/term"
	"github.com/spf13/cobra"
)

const (
	TARGET_SINGLE  = "single"
	TARGET_CLUSTER = "cluster"
)

func NewInstallCommand() *cobra.Command {
	installCmd := &cobra.Command{
		Use:   "install",
		Short: "Set up amp services in swarm environment",
		RunE:  install,
	}
	return installCmd
}

func install(cmd *cobra.Command, args []string) error {
	stdin, stdout, stderr := term.StdStreams()
	dockerCli := docker.NewDockerCli(stdin, stdout, stderr)

	namespace := "amp"
	if len(args) > 0 {
		namespace = args[0]
	}

	etcdClusterMode, err := serviceDeploymentMode(dockerCli.Client(), "amp.type.kv", "true")
	if err != nil {
		return err
	}
	elasticsearchClusterMode, err := serviceDeploymentMode(dockerCli.Client(), "amp.type.search", "true")
	if err != nil {
		return err
	}
	clusterMode := map[string]string{"elasticsearch": elasticsearchClusterMode, "etcd": etcdClusterMode}
	files, err := getStackFiles("./stacks", clusterMode)
	if err != nil {
		return err
	}

	for _, f := range files {
		log.Println(f)
		if strings.Contains(f, "test") {
			err := deployTest(dockerCli, f, "test", 60 /* timeout in seconds */)
			stack.Remove(dockerCli, stack.RemoveOptions{Namespaces: []string{"test"}})
			if err != nil {
				return err
			}
		} else {
			err := deploy(dockerCli, f, namespace)
			if err != nil {
				return err
			}
			time.Sleep(10 * time.Second)
		}
	}
	return nil
}

// returns the deployment mode
// based on the number of nodes with the label passed as argument
// if number of nodes > 2, mode = cluster, else mode = single
func serviceDeploymentMode(c client.APIClient, labelKey string, labelValue string) (string, error) {
	// unfortunately filtering labels on NodeList won't work as expected, Cf. https://github.com/moby/moby/issues/27231
	nodes, err := c.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		return "", err
	}
	matchingNodes := 0
	for _, node := range nodes {
		// node is a swarm.Node
		for k, v := range node.Spec.Labels {
			if k == labelKey {
				if labelValue == "" || labelValue == v {
					matchingNodes++
				}
			}
		}
	}
	switch matchingNodes {
	case 0:
		return "", fmt.Errorf("can't find a node with label %s", labelKey)
	case 1:
		fallthrough
	case 2:
		return TARGET_SINGLE, nil
	default:
		return TARGET_CLUSTER, nil
	}
}

// returns sorted list of yaml file pathnames
func getStackFiles(path string, clusterMode map[string]string) ([]string, error) {
	if path == "" {
		path = "./stacks"
	}

	// a bit more work but we can't just use filepath.Glob
	// since we need to match both *.yml and *.yaml
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	stackfiles := []string{}
	for _, f := range files {
		name := f.Name()
		// not compiling regex since only expecting less than a dozen stackfiles
		matched, err := regexp.MatchString("\\.ya?ml$", name)
		if err != nil {
			log.Println(err)
		} else if matched {
			// looking for the service name, in case there's an indication for the cluster mode (single vs cluster)
			// expecting a file with a name NN-SERVICENAME-mode.*
			split := strings.Split(name, "-")
			if len(split) == 3 {
				serviceName := split[1]
				if strings.Contains(name, TARGET_SINGLE) && clusterMode[serviceName] != TARGET_SINGLE {
					continue
				}
				if strings.Contains(name, TARGET_CLUSTER) && clusterMode[serviceName] != TARGET_CLUSTER {
					continue
				}
			}
			stackfiles = append(stackfiles, filepath.Join(path, name))
		}
	}
	return stackfiles, nil
}

func deploy(d *command.DockerCli, stackfile string, namespace string) error {
	return deployExpectingState(d, stackfile, namespace, swarm.TaskStateRunning)
}

func deployExpectingState(d *command.DockerCli, stackfile string, namespace string, expectedState swarm.TaskState) error {
	if namespace == "" {
		// use the stackfile basename as the default stack namespace
		namespace = filepath.Base(stackfile)
		namespace = strings.TrimSuffix(namespace, filepath.Ext(namespace))
	}

	opts := stack.DeployOptions{
		Namespace:        namespace,
		Composefile:      stackfile,
		ResolveImage:     stack.ResolveImageNever,
		SendRegistryAuth: false,
		Prune:            false,
		ExpectedState:    expectedState,
	}

	return stack.Deploy(context.Background(), d, opts)
}

func deployTest(d *command.DockerCli, stackfile string, namespace string, timeout int) error {
	// Deploy the test stack
	if err := deployExpectingState(d, stackfile, namespace, swarm.TaskStateComplete); err != nil {
		return err
	}

	// Create a docker client
	c, err := client.NewClient(admin.DefaultURL, admin.DefaultVersion, nil, nil)
	if err != nil {
		return err
	}

	// List stack tasks
	options := types.TaskListOptions{Filters: filters.NewArgs()}
	options.Filters.Add("label", convert.LabelNamespace+"="+namespace)
	tasks, err := stack.ListTasks(context.Background(), c, options)
	if err != nil {
		return err
	}

	// Assert we have at least one task
	if len(tasks) == 0 {
		return fmt.Errorf("no task for test")
	}

	// Assert we have only one task
	if len(tasks) != 1 {
		return fmt.Errorf("too many tasks for test: %d", len(tasks))
	}

	// If the task has an error, the test has failed
	task := tasks[0]
	if task.Status.Err != "" {
		return fmt.Errorf("test failed with status: %s", task.Status.Err)
	}

	log.Println("Test successful")
	return nil
}
