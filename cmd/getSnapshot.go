package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// getSnapshotCmd represents the getSnapshot command
var getSnapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Get the metadata for a given snapshot",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		cluster, err := client.GetSnapshot(args[0])
		if err != nil {
			panic(err)
		}

		j, err := json.MarshalIndent(cluster, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(j))
	},
}

func init() {
	getCmd.AddCommand(getSnapshotCmd)
}
