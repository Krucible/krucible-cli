package cmd

import (
	"github.com/spf13/cobra"
)

// helmCmd represents the helm command
var helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Run kubebox for the given cluster",
	Run: func(cmd *cobra.Command, args []string) {
		RunBinary("helm", ClusterID, args)
	},
}

func init() {
	rootCmd.AddCommand(helmCmd)

	helmCmd.Flags().StringVarP(&ClusterID, "cluster", "c", "", "The ID of the cluster upon which these commands should be run")
	helmCmd.MarkFlagRequired("cluster")
}
