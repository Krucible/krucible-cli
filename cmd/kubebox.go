package cmd

import (
	"github.com/spf13/cobra"
)

// kubeboxCmd represents the kubebox command
var kubeboxCmd = &cobra.Command{
	Use:   "kubebox",
	Short: "Run kubebox for the given cluster",
	Long:  "kubebox is an interactive dashboard for your Kubernetes cluster, allowing you to view logs and metrics.",
	Run: func(cmd *cobra.Command, args []string) {
		RunBinary("kubebox", ClusterID, args)
	},
}

func init() {
	rootCmd.AddCommand(kubeboxCmd)
	kubeboxCmd.Flags().StringVarP(&ClusterID, "cluster", "c", "", "The ID of the cluster upon which these commands should be run")
	kubeboxCmd.MarkFlagRequired("cluster")
}
