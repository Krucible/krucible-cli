package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// getClustersCmd represents the get clusters command
var getClustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Get all clusters within the account",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		clusters, err := client.GetClusters()
		if err != nil {
			panic(err)
		}

		jsonBytes, err := json.MarshalIndent(clusters, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonBytes))
	},
}

func init() {
	getCmd.AddCommand(getClustersCmd)
}
