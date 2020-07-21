package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// getSnapshotsCmd represents the getSnapshots command
var getSnapshotsCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "Get all snapshots in the account",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		snapshots, err := client.GetSnapshots()
		if err != nil {
			panic(err)
		}

		jsonBytes, err := json.MarshalIndent(snapshots, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonBytes))
	},
}

//func init() {
//getCmd.AddCommand(getSnapshotsCmd)
//}
