package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func runKubectlCommand(args ...string) {
	kubectlPath, err := exec.LookPath("kubectl")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: kubectl is not on the $PATH. Please install it.")
		os.Exit(1)
	}
	kctlExec := exec.Command(kubectlPath, args...)
	if err = kctlExec.Run(); err != nil {
		panic(err)
	}
}

func configureKubectl(clusterID, server string) {
	runKubectlCommand("config", "set-cluster", clusterID, "--server", server)
	runKubectlCommand("config", "set-context", clusterID, "--cluster", clusterID)
	runKubectlCommand("config", "use-context", clusterID)
}

// configureKubectlCmd represents the configure-kubectl command
var configureKubectlCmd = &cobra.Command{
	Use:   "configure-kubectl [krucible cluster ID]",
	Args:  cobra.ExactArgs(1),
	Short: "Configure your kubectl context to connect to the given Krucible cluster",
	Run: func(cmd *cobra.Command, args []string) {
		clusterID := args[0]
		client := getClientOrDie()
		cluster, err := client.GetCluster(clusterID)
		if err != nil {
			panic(err)
		}

		configureKubectl(cluster.ID, cluster.ConnectionDetails.Server)
	},
}

func init() {
	rootCmd.AddCommand(configureKubectlCmd)
}
