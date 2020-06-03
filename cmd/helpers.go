package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/Krucible/krucible-cli/pkg/binaryfetcher"
	"github.com/Krucible/krucible-go-client/krucible"
)

func createClusterConfigFile(client *krucible.Client, clusterID string) (configFilePath string) {
	cluster, err := client.GetCluster(ClusterID)
	if cluster.State != "running" {
		fmt.Fprintln(os.Stderr, "Cluster not in running state")
		os.Exit(1)
	}

	kubeconfig, err := client.GetClusterKubeConfig(ClusterID)
	if err != nil {
		panic(err)
	}

	configFilePath = path.Join(getConfigDirOrDie(), ClusterID+"-config.json")
	if err := ioutil.WriteFile(configFilePath, []byte(kubeconfig), 0644); err != nil {
		panic(err)
	}
	return configFilePath
}

func RunBinary(binaryName, clusterID string, args []string) {
	c := getClientOrDie()
	configFilePath := createClusterConfigFile(c, clusterID)
	bin := binaryfetcher.GetBinary(binaryName, getConfigDirOrDie())
	kcmd := exec.Command(bin, args...)
	kcmd.Env = []string{"KUBECONFIG=" + configFilePath}
	kcmd.Stdout = os.Stdout
	kcmd.Stderr = os.Stderr
	kcmd.Stdin = os.Stdin
	if err := kcmd.Start(); err != nil {
		panic(err)
	}
	if err := kcmd.Wait(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		panic(err)
	}
}
