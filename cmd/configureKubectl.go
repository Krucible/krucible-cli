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
	if kOutput, err := kctlExec.CombinedOutput(); err != nil {
		fmt.Fprintln(os.Stderr, "kubectl command errored. This is the output:")
		fmt.Fprintf(os.Stderr, string(kOutput))
		panic(err)
	}
}

func configureKubectl(clusterID, server, ca, authToken string) {
	filePath := path.Join(getConfigDirOrDie(), clusterID+"-cert.pem")
	ioutil.WriteFile(filePath, []byte(ca), 0644)
	runKubectlCommand("config",
		"set-cluster", clusterID,
		"--server", server,
		"--certificate-authority", filePath,
		"--embed-certs",
	)
	runKubectlCommand("config", "set-credentials", "krucible-"+clusterID, "--token", authToken)
	runKubectlCommand("config", "set-context", clusterID, "--cluster", clusterID, "--user", "krucible-"+clusterID)
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
			cluster.ConnectionDetails.ClusterAuthToken,
		)
	},
}

func init() {
	rootCmd.AddCommand(configureKubectlCmd)
}
