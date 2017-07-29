package main

import (
	"context"
	"fmt"
	"log"

	plugin "github.com/appcelerator/amp/cluster/plugin/local"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

const (
	defaultURL     = "unix:///var/run/docker.sock"
	defaultVersion = "1.30"
)

var (
	// Version is set with a linker flag (see Makefile)
	Version string
	// Build is set with a linker flag (see Makefile)
	Build               string
	dockerClientVersion string
	engineURL           string
	dockerClient        *client.Client
	defaultLabels       = map[string]string{"amp.type.api": "true", "amp.type.route": "true", "amp.type.search": "true", "amp.type.kv": "true", "amp.type.mq": "true", "amp.type.metrics": "true", "amp.type.core": "true", "amp.type.user": "true"}
	opts                = &plugin.RequestOptions{InitRequest: swarm.InitRequest{}, Labels: defaultLabels}
)

func initClient(cmd *cobra.Command, args []string) (err error) {
	dockerClient, err = client.NewClient(defaultURL, defaultVersion, nil, nil)
	return
}

func create(cmd *cobra.Command, args []string) {
	// docker swarn init --advertise-addr $interface
	ctx := context.Background()

	if err := plugin.EnsureSwarmExists(ctx, dockerClient, opts); err != nil {
		log.Fatal(err)
	}
	if err := plugin.LabelNode(ctx, dockerClient, opts); err != nil {
		log.Fatal(err)
	}
	if err := plugin.RunAgent(ctx, dockerClient, "", opts); err != nil {
		log.Fatal(err)
	}
	// use the info command to print json cluster info to stdout
	info(cmd, args)
}

func update(cmd *cobra.Command, args []string) {
	// nothing to do
}

func delete(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	if err := plugin.RunAgent(ctx, dockerClient, "uninstall", opts); err != nil {
		log.Fatal(err)
	}
	if err := plugin.DeleteCluster(ctx, dockerClient, opts); err != nil {
		log.Fatal(err)
	}

	log.Println("cluster deleted")
}

func info(cmd *cobra.Command, args []string) {
	// docker node inspect self -f '{{.Status.State}}'
	ctx := context.Background()
	swarmResp, err := plugin.InfoCluster(ctx, dockerClient)
	if err != nil {
		log.Fatal(err)
	}
	nodeResp, err := plugin.InfoNode(ctx, dockerClient)
	if err != nil {
		log.Fatal(err)
	}

	j, err := plugin.InfoToJSON(swarmResp, nodeResp)
	if err != nil {
		log.Fatal(err)
	}
	// print json result to stdout
	fmt.Print(j)
}

func main() {
	rootCmd := &cobra.Command{
		Use:               "localplugin",
		Short:             "init/update/destroy an local cluster in Docker swarm mode",
		PersistentPreRunE: initClient,
	}
	rootCmd.PersistentFlags().StringVarP(&opts.Tag, "tag", "t", "latest", "Tag (version) to deploy")
	rootCmd.PersistentFlags().StringVarP(&opts.Registration, "registration", "r", "email", "registration mode")
	rootCmd.PersistentFlags().StringVarP(&opts.Notifications, "notifications", "n", "none", "notifications mode")
	rootCmd.PersistentFlags().StringVarP(&opts.InitRequest.ListenAddr, "listen-addr", "l", "0.0.0.0:2377", "Listen address")
	rootCmd.PersistentFlags().StringVarP(&opts.InitRequest.AdvertiseAddr, "advertise-addr", "a", "eth0", "Advertise address")
	rootCmd.PersistentFlags().BoolVarP(&opts.InitRequest.ForceNewCluster, "force-new-cluster", "", false, "force initialization of a new swarm")
	rootCmd.PersistentFlags().BoolVarP(&opts.ForceLeave, "force-leave", "", false, "force leave the swarm")

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "init cluster in swarm mode",
		Run:   create,
	}

	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "get information about the cluster",
		Run:   info,
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update the cluster",
		Run:   update,
	}

	destroyCmd := &cobra.Command{
		Use:   "destroy",
		Short: "destroy the cluster",
		Run:   delete,
	}

	rootCmd.AddCommand(initCmd, infoCmd, updateCmd, destroyCmd)

	_ = rootCmd.Execute()
}
