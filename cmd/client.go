package cmd

import (
	"fmt"
	"os"

	"github.com/Krucible/krucible-go-client/krucible"
	"github.com/spf13/viper"
)

func getClientOrDie() *krucible.Client {
	if !viper.IsSet("apiKeyId") || !viper.IsSet("apiKeySecret") || !viper.IsSet("accountId") {
		fmt.Fprintln(os.Stderr, "Error: authentication not configured. Please run krucible set-config")
		os.Exit(1)
	}

	var baseUrl string
	if viper.IsSet("baseUrl") {
		baseUrl = viper.GetString("baseUrl")
	} else {
		baseUrl = "https://usekrucible.com/api"
	}

	return krucible.NewClient(krucible.ClientConfig{
		BaseURL:      baseUrl,
		APIKeyId:     viper.GetString("apiKeyId"),
		APIKeySecret: viper.GetString("apiKeySecret"),
		AccountID:    viper.GetString("accountId"),
	})
}
