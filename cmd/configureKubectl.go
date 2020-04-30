package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

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

func configureKubectl(clusterID, server, ca string) {
	filePath := path.Join(getConfigDirOrDie(), clusterID+"-cert.pem")
	ioutil.WriteFile(filePath, []byte(ca), 0644)
	runKubectlCommand("config",
		"set-cluster", clusterID,
		"--server", server,
		"--certificate-authority", filePath,
	)
	runKubectlCommand("config", "set-credentials", "krucible-"+clusterID, "--token", "krucible")
	runKubectlCommand("config", "set-context", clusterID, "--cluster", clusterID, "--user", "krucible")
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

		if cluster.State != "running" {
			fmt.Fprintln(os.Stderr, "Cluster not in running state")
			os.Exit(1)
		}

		configureKubectl(
			cluster.ID,
			cluster.ConnectionDetails.Server,
			cluster.ConnectionDetails.CertificateAuthority,
		)
	},
}

func init() {
	rootCmd.AddCommand(configureKubectlCmd)
}
