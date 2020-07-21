package cmd

import (
	"github.com/spf13/cobra"
)

// deleteClusterCmd represents the deleteCluster command
var deleteClusterCmd = &cobra.Command{
	Use:   "cluster [cluster ID]",
	Short: "Delete a given Krucible cluster",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		if err := client.DeleteCluster(args[0]); err != nil {
			panic(err)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteClusterCmd)
}
