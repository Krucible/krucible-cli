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

// configureKubectlCmd represents the configure-kubectl command
var configureKubectlCmd = &cobra.Command{
	Use:   "configure-kubectl [krucible cluster ID]",
	Args:  cobra.ExactArgs(1),
	Short: "Configure your kubectl context to connect to the given Krucible cluster",
	Run: func(cmd *cobra.Command, args []string) {
		client := getClientOrDie()
		cluster, err := client.GetCluster(args[0])
		if err != nil {
			panic(err)
		}

		runKubectlCommand("config", "set-cluster", cluster.ID, "--server", cluster.ConnectionDetails.Server)
		runKubectlCommand("config", "set-context", cluster.ID, "--cluster", cluster.ID)
		runKubectlCommand("config", "use-context", cluster.ID)
	},
}

func init() {
	rootCmd.AddCommand(configureKubectlCmd)
}
