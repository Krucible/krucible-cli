package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/Krucible/krucible-cli/pkg/binaryfetcher"
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
		c := getClientOrDie()
		cluster, err := c.GetCluster(ClusterID)
		if cluster.State != "running" {
			fmt.Fprintln(os.Stderr, "Cluster not in running state")
			os.Exit(1)
		}

		kubeconfig, err := c.GetClusterKubeConfig(ClusterID)
		if err != nil {
			panic(err)
		}

		filePath := path.Join(getConfigDirOrDie(), ClusterID+"-config.json")
		if err := ioutil.WriteFile(filePath, []byte(kubeconfig), 0644); err != nil {
			panic(err)
		}

		kubectlBinary := binaryfetcher.GetKubectlBinary(getConfigDirOrDie())
		kargs := append([]string{"--kubeconfig", filePath}, args...)
		kcmd := exec.Command(kubectlBinary, kargs...)
		kcmd.Stdout = os.Stdout
		kcmd.Stderr = os.Stderr
		kcmd.Stdin = os.Stdin
		kcmd.Start()
		kcmd.Wait()
	},
}

func init() {
	rootCmd.AddCommand(kubectlCmd)

	kubectlCmd.Flags().StringVarP(&ClusterID, "cluster", "c", "", "The ID of the cluster upon which these commands should be run")
	kubectlCmd.MarkFlagRequired("cluster")
}
