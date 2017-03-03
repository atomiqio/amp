package main

import (
	"errors"
	"fmt"
	"github.com/appcelerator/amp/api/rpc/function"
	"github.com/appcelerator/amp/cmd/amp/cli"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

var (
	functionCmd = &cobra.Command{
		Use:     "function",
		Short:   "Function operations",
		Long:    `Function command manages all function-related operations.`,
		Aliases: []string{"fn"},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return AMP.Connect()
		},
	}

	createFunctionCmd = &cobra.Command{
		Use:   "create FUNC-NAME IMAGE",
		Short: "Create a function",
		Long: `The create command registers a function with the specified name and image.
If successful, a function id is returned.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return createFunction(AMP, cmd, args)
		},
	}

	listFunctionCmd = &cobra.Command{
		Use:   "ls [OPTION]",
		Short: "List functions",
		Long:  `The list command displays all registered functions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listFunction(AMP, cmd, args)
		},
	}

	removeFunctionCmd = &cobra.Command{
		Use:   "rm FUNC-ID",
		Short: "Remove a function",
		Long:  `The remove command unregisters the specified function.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return removeFunction(AMP, cmd, args)
		},
	}
)

func init() {
	listFunctionCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")

	functionCmd.AddCommand(createFunctionCmd)
	functionCmd.AddCommand(listFunctionCmd)
	functionCmd.AddCommand(removeFunctionCmd)
	RootCmd.AddCommand(functionCmd)
}

func createFunction(amp *cli.AMP, cmd *cobra.Command, args []string) (err error) {
	switch len(args) {
	case 0:
		return errors.New("must specify function name and docker image")
	case 1:
		return errors.New("must specify docker image")
	case 2:
	// OK
	default:
		return errors.New("too many arguments")
	}

	name, image := strings.TrimSpace(args[0]), strings.TrimSpace(args[1])
	if name == "" {
		return errors.New("function name cannot be empty")
	}
	if image == "" {
		return errors.New("docker image cannot be empty")
	}

	// Create function
	request := &function.CreateRequest{Function: &function.FunctionEntry{
		Name:  name,
		Image: image,
	}}
	reply, er := function.NewFunctionClient(amp.Conn).Create(context.Background(), request)
	if er != nil {
		manager.fatalf(grpc.ErrorDesc(er))
		return
	}

	fmt.Println(reply.Function.Id)
	return nil
}

func listFunction(amp *cli.AMP, cmd *cobra.Command, args []string) (err error) {
	// List functions
	request := &function.ListRequest{}
	reply, er := function.NewFunctionClient(amp.Conn).List(context.Background(), request)
	if er != nil {
		manager.fatalf(grpc.ErrorDesc(er))
		return
	}

	// --quiet only display IDs
	if quiet, err := strconv.ParseBool(cmd.Flag("quiet").Value.String()); err != nil {
		return fmt.Errorf("Unable to convert quiet parameter: %v", cmd.Flag("f").Value.String())
	} else if quiet {
		for _, fn := range reply.Functions {
			fmt.Println(fn.Id)
		}
		return nil
	}

	// Table view
	w := tabwriter.NewWriter(os.Stdout, 0, 0, tablePadding, ' ', 0)
	fmt.Fprintln(w, "ID\tName\tImage\tOwner")
	for _, fn := range reply.Functions {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", fn.Id, fn.Name, fn.Image, fn.Owner)
	}
	w.Flush()

	return nil
}

func removeFunction(amp *cli.AMP, cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		return errors.New("rm requires at least one argument")
	}

	client := function.NewFunctionClient(amp.Conn)
	for _, arg := range args {
		if arg == "" {
			continue
		}

		request := &function.DeleteRequest{Id: arg}
		_, er := client.Delete(context.Background(), request)
		if er != nil {
			manager.fatalf(grpc.ErrorDesc(er))
			return
		} else {
			fmt.Println(arg)
		}
	}

	return nil
}
