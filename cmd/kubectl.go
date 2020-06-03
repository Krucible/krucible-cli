package cmd

import (
	"github.com/spf13/cobra"
)

// kubectlCmd represents the kubectl command
var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "Run kubectl commands on a given cluster",
	Long: `Run any kubectl command on a given krucible cluster.
Use -- to separate krucible arguments from kubectl arguments
For example: krucible kubectl --cluster $CLUSTER_ID -- get pods`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		RunBinary("kubectl", ClusterID, args)
	},
}

func init() {
	rootCmd.AddCommand(kubectlCmd)

	kubectlCmd.Flags().StringVarP(&ClusterID, "cluster", "c", "", "The ID of the cluster upon which these commands should be run")
	kubectlCmd.MarkFlagRequired("cluster")
}
