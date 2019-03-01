package cmd

import (
	"KeyValueCache/kvcache"
	"github.com/spf13/cobra"
)

func Run() error {
	runner := &kvcCommandRunner{
		kvcache.NewSimpleKVCache(),
	}

	var RootCmd = &cobra.Command{Use: "kvc"}

	var putCmd = &cobra.Command{
		Use:   "put <key> <value>",
		Short: "store <value> under <key>",
		Long:  "store <value> under <key>",
		RunE: runner.Put,
	}

	RootCmd.AddCommand(putCmd)
	return RootCmd.Execute()
}

type kvcCommandRunner struct {
	cache kvcache.KeyValueCache
}

func (runner kvcCommandRunner) Put(cmd *cobra.Command, args []string) error {
	//fmt.Println("put", args)

	// Logic here
	return nil
}