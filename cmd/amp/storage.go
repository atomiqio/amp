package main

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/appcelerator/amp/api/rpc/storage"
	"github.com/appcelerator/amp/cmd/amp/cli"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// StorageCmd is the main command for attaching storage subcommands.
var StorageCmd = &cobra.Command{
	Use:   "kv",
	Short: "Storage operations",
	Long:  `KV command manages all key-value based storage operations in the ETCD.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return AMP.Connect()
	},
}

var (
	// storagePutCmd represents the creation of storage key-value pair
	storagePutCmd = &cobra.Command{
		Use:   "put KEY VALUE",
		Short: "Assign specified value with specified key",
		Long: `The put command creates a storage object with the key-value input if the key does not already exist.
Else, it updates the existing key with the new input value.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return storagePut(AMP, cmd, args)
		},
	}
	// storageGetCmd represents the retrieval of storage value based on key
	storageGetCmd = &cobra.Command{
		Use:   "get KEY",
		Short: "Retrieve a storage object",
		Long:  `The get command retrieves a key-value pair based on the specified input key.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return storageGet(AMP, cmd, args)
		},
	}
	// storageDeleteCmd represents the deletion of storage value based on key
	storageDeleteCmd = &cobra.Command{
		Use:     "rm KEY or del KEY",
		Short:   "Remove a storage object (alias: del)",
		Long:    `The remove command deletes the key-value pair in storage based on the specified input key.`,
		Aliases: []string{"del"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return storageDelete(AMP, cmd, args)
		},
	}
	// storageListCmd represents the list of storage key-value pair
	storageListCmd = &cobra.Command{
		Use:   "ls",
		Short: "List all storage objects",
		Long:  `The list command returns a list of all the key-value pair in storage.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return storageList(AMP, cmd, args)
		},
	}
)

func init() {
	RootCmd.AddCommand(StorageCmd)
	StorageCmd.AddCommand(storagePutCmd)
	StorageCmd.AddCommand(storageGetCmd)
	StorageCmd.AddCommand(storageDeleteCmd)
	StorageCmd.AddCommand(storageListCmd)
}

// storagePut validates the input command line arguments and creates or updates storage key-value pair
// by invoking the corresponding rpc/storage method
func storagePut(amp *cli.AMP, cmd *cobra.Command, args []string) (err error) {
	switch len(args) {
	case 0:
		return errors.New("must specify storage key and storage value")
	case 1:
		return errors.New("must specify storage value")
	case 2:
		// OK
	default:
		return errors.New("too many arguments")
	}

	k := args[0]
	v := args[1]
	request := &storage.PutStorage{Key: k, Val: v}
	client := storage.NewStorageClient(amp.Conn)
	reply, er := client.Put(context.Background(), request)
	if er != nil {
		manager.fatalf(grpc.ErrorDesc(er))
		return
	}
	fmt.Println(reply.Val)
	return nil
}

// storageGet validates the input command line arguments and retrieves storage key-value pair
//by invoking the corresponding rpc/storage method
func storageGet(amp *cli.AMP, cmd *cobra.Command, args []string) (err error) {
	if len(args) > 1 {
		return errors.New("too many arguments - check again")
	} else if len(args) == 0 {
		return errors.New("must specify storage key")
	}
	k := args[0]
	if k == "" {
		return errors.New("must specify storage key")
	}

	request := &storage.GetStorage{Key: k}

	client := storage.NewStorageClient(amp.Conn)
	reply, er := client.Get(context.Background(), request)
	if er != nil {
		manager.fatalf(grpc.ErrorDesc(er))
		return
	}
	fmt.Println(reply.Val)
	return nil
}

// storageDelete validates the input command line arguments and deletes storage key-value pair
// by invoking the corresponding rpc/storage method
func storageDelete(amp *cli.AMP, cmd *cobra.Command, args []string) (err error) {
	if len(args) > 1 {
		return errors.New("too many arguments - check again")
	} else if len(args) == 0 {
		return errors.New("must specify storage key")
	}
	k := args[0]
	if k == "" {
		return errors.New("must specify storage key")
	}

	request := &storage.DeleteStorage{Key: k}

	client := storage.NewStorageClient(amp.Conn)
	reply, er := client.Delete(context.Background(), request)
	if er != nil {
		manager.fatalf(grpc.ErrorDesc(er))
		return
	}
	fmt.Println(reply.Val)
	return nil
}

// storageList validates the input command line arguments and lists all the storage
// key-value pairs by invoking the corresponding rpc/storage method
func storageList(amp *cli.AMP, cmd *cobra.Command, args []string) (err error) {
	if len(args) > 0 {
		return errors.New("too many arguments - check again")
	}
	request := &storage.ListStorage{}
	client := storage.NewStorageClient(amp.Conn)
	reply, er := client.List(context.Background(), request)
	if er != nil {
		manager.fatalf(grpc.ErrorDesc(er))
		return
	}
	if reply == nil || len(reply.List) == 0 {
		fmt.Println("No storage object is available")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, "KEY\tVALUE\t")
	fmt.Fprintln(w, "---\t-----\t")
	for _, info := range reply.List {
		fmt.Fprintf(w, "%s\t%s\t\n", info.Key, info.Val)
	}
	w.Flush()
	return nil
}
