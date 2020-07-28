package binaryfetcher

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/Krucible/krucible-go-client/krucible"
	"github.com/briandowns/spinner"
)

var binaryDownloadLocations = map[string]map[string]string{
	"kubectl": map[string]string{
		"windows": "https://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/windows/amd64/kubectl.exe",
		"linux":   "https://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/linux/amd64/kubectl",
		"darwin":  "https://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/darwin/amd64/kubectl",
	},
	"kubebox": map[string]string{
		"windows": "https://github.com/astefanutti/kubebox/releases/download/v0.8.0/kubebox-windows.exe",
		"linux":   "https://github.com/astefanutti/kubebox/releases/download/v0.8.0/kubebox-linux",
		"darwin":  "https://github.com/astefanutti/kubebox/releases/download/v0.8.0/kubebox-macos",
	},
	"helm": map[string]string{
		"windows": "https://krucible-cli-binaries.s3-eu-west-1.amazonaws.com/helm/helm-v3.2.1-windows-amd64.exe",
		"linux":   "https://krucible-cli-binaries.s3-eu-west-1.amazonaws.com/helm/helm-v3.2.1-linux-amd64",
		"darwin":  "https://krucible-cli-binaries.s3-eu-west-1.amazonaws.com/helm/helm-v3.2.1-darwin-amd64",
	},
}

func createClusterConfigFile(client *krucible.Client, clusterID, configDir string) (configFilePath string) {
	cluster, err := client.GetCluster(clusterID)
	if cluster.State != "running" {
		fmt.Fprintln(os.Stderr, "Cluster not in running state")
		os.Exit(1)
	}

	kubeconfig, err := client.GetClusterKubeConfig(clusterID)
	if err != nil {
		panic(err)
	}

	configFilePath = path.Join(configDir, clusterID+"-config.json")
	if err := ioutil.WriteFile(configFilePath, []byte(kubeconfig), 0644); err != nil {
		panic(err)
	}
	return configFilePath
}

func GetBinary(binaryName, configDir string) string {
	os.MkdirAll(path.Join(configDir, "bin"), 0755)
	downloadLocation := binaryDownloadLocations[binaryName][runtime.GOOS]
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}
	binaryLocation := path.Join(configDir, "bin", binaryName)
	if _, err := os.Stat(binaryLocation); os.IsNotExist(err) {
		binary, err := os.OpenFile(binaryLocation, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0770)
		if err != nil {
			panic(err)
		}

		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
		s.Prefix = "Downloading " + binaryName + "... "
		s.Start()
		resp, err := http.Get(downloadLocation)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != 200 {
			panic("Received the following status code: " + string(resp.StatusCode))
		}
		defer resp.Body.Close()
		defer binary.Sync()
		defer binary.Close()

		_, err = io.Copy(binary, resp.Body)
		if err != nil {
			panic(err)
		}
		s.Stop()
	}

	return binaryLocation
}
