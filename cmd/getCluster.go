package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// getClusterCmd represents the getCluster command
var getClusterCmd = &cobra.Command{
	Use:   "cluster",
	Args:  cobra.ExactArgs(1),
	Short: "Get Krucible cluster metadata",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		cluster, err := client.GetCluster(args[0])
		if err != nil {
			panic(err)
		}

		j, err := json.Marshal(cluster)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(j))
	},
}

func init() {
	getCmd.AddCommand(getClusterCmd)
}
