package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Krucible/krucible-go-client/krucible"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var DisplayName string

// clusterCmd represents the cluster command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Create a new cluster",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Prefix = "Creating cluster... "
		s.Start()
		newClusterResult, err := client.CreateCluster(krucible.CreateClusterConfig{
			DisplayName: DisplayName,
		})
		s.Stop()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "Cluster "+newClusterResult.Cluster.ID+" created and ready to use!")
	},
}

func init() {
	createCmd.AddCommand(clusterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	clusterCmd.Flags().StringVarP(&DisplayName, "display-name", "n", "", "Desired display name for the cluster")
	clusterCmd.MarkFlagRequired("display-name")
}
