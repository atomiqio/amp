package main

import (
	"github.com/spf13/cobra"
	"os"
)

// PlatformStop is the main command for attaching platform subcommands.
var PlatformStop = &cobra.Command{
	Use:   "stop",
	Short: "Stop platform",
	Long:  `Stop all AMP platform services.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopAMP(cmd, args)
	},
}

func init() {
	PlatformStop.Flags().BoolP("silence", "s", false, "no console output at all")
	PlatformCmd.AddCommand(PlatformStop)
}

func stopAMP(cmd *cobra.Command, args []string) error {
	manager := &ampManager{}
	if cmd.Flag("silence").Value.String() == "true" {
		manager.silence = true
	}
	if cmd.Flag("verbose").Value.String() == "true" {
		manager.verbose = true
	}
	if err := manager.init("Stopping AMP platform"); err != nil {
		manager.printf(colError, "Start error: %v\n", err)
		os.Exit(1)
	}
	if cmd.Flag("server").Value.String() != "" {
		manager.printf(colWarn, "Error: --server has no effect for stop command\n")
		os.Exit(1)
	}
	stack := getAMPInfrastructureStack(manager)
	manager.computeStatus(stack)
	if manager.status == "stopped" {
		manager.printf(colMagenta, "AMP platform already stopped\n")
		return nil
	}
	if err := manager.stop(stack); err != nil {
		manager.printf(colError, "Stop error: %v\n", err)
		os.Exit(1)
	}
	manager.printf(colMagenta, "AMP platform stopped\n")
	return nil
}
