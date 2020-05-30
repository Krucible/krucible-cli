package binaryfetcher

import (
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
)

func GetKubectlBinary(configDir string) string {
	downloadLocation := map[string]string{
		"windows": "https://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/windows/amd64/kubectl.exe",
		"linux":   "https://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/linux/amd64/kubectl",
		"darwin":  "https://storage.googleapis.com/kubernetes-release/release/v1.18.0/bin/darwin/amd64/kubectl",
	}[runtime.GOOS]

	binaryLocation := path.Join(configDir, "kubectl")
	if _, err := os.Stat(binaryLocation); os.IsNotExist(err) {
		binary, err := os.OpenFile(binaryLocation, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0770)
		if err != nil {
			panic(err)
		}

		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
		s.Prefix = "Downloading kubectl... "
		s.Start()
		resp, err := http.Get(downloadLocation)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		_, err = io.Copy(binary, resp.Body)
		if err != nil {
			panic(err)
		}
		s.Stop()
	}

	return binaryLocation
}
