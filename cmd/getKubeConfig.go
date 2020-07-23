package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getKubeConfigCmd represents the getKubeConfig command
var getKubeConfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Get the kubeconfig of a Krucible cluster",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		kubeconfig, err := client.GetClusterKubeConfig(args[0])
		if err != nil {
			panic(err)
		}

		fmt.Println(string(kubeconfig))
	},
}

func init() {
	getCmd.AddCommand(getKubeConfigCmd)
}
