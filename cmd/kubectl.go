package cmd

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

// kubectlCmd represents the kubectl command
var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "Run kubectl commands on a given cluster",
	Long: `Run any kubectl command on a given krucible cluster.
Use -- to separate krucible arguments from kubectl arguments
For example: krucible kubectl --cluster $CLUSTER_ID -- get pods`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		c := getClientOrDie()
		kubeconfig, err := c.GetClusterKubeConfig(ClusterID)
		if err != nil {
			panic(err)
		}

		filePath := path.Join(getConfigDirOrDie(), ClusterID+"-config.json")
		if err := ioutil.WriteFile(filePath, []byte(kubeconfig), 0644); err != nil {
			panic(err)
		}

		kargs := append([]string{"--kubeconfig", filePath}, args...)
		kcmd := exec.Command("kubectl", kargs...)
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
