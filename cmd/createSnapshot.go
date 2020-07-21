package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/Krucible/krucible-go-client/krucible"
	"github.com/spf13/cobra"
)

var ClusterID string

// createSnapshotCmd represents the createSnapshot command
var createSnapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Create a snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		snapshot, err := client.CreateSnapshot(krucible.CreateSnapshotConfig{
			ClusterID: ClusterID,
		})
		if err != nil {
			panic(err)
		}

		jsonBytes, err := json.MarshalIndent(snapshot, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonBytes))
	},
}

//func init() {
//createCmd.AddCommand(createSnapshotCmd)
//createSnapshotCmd.Flags().StringVarP(&ClusterID, "cluster", "c", "", "The ID of the cluster to snapshot")
//createSnapshotCmd.MarkFlagRequired("cluster")
//}
