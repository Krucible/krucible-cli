package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var APIKeyID string
var APIKeySecret string
var AccountID string
var BaseURL string

// setConfigCmd represents the setConfig command
var setConfigCmd = &cobra.Command{
	Use:   "set-config",
	Short: "Set up the connection to the Krucible API",
	Long: `Configure the connection details for the Krucible API.

API keys can be created at https://usekrucible.com/api-keys`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("apiKeyId", APIKeyID)
		viper.Set("apiKeySecret", APIKeySecret)
		viper.Set("accountId", AccountID)
		if BaseURL != "" {
			viper.Set("baseUrl", BaseURL)
		}

		if viper.ConfigFileUsed() == "" {
			configFilePath := path.Join(getConfigDirOrDie(), "config.json")
			f, err := os.Create(configFilePath)
			if err != nil {
				panic(err)
			}
			f.Close()
		}
		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setConfigCmd)
	setConfigCmd.Flags().StringVar(&APIKeyID, "api-key-id", "", "Your API key ID")
	setConfigCmd.Flags().StringVar(&APIKeySecret, "api-key-secret", "", "Your API key secret")
	setConfigCmd.Flags().StringVar(&AccountID, "account-id", "", "Your account ID")
	setConfigCmd.Flags().StringVar(&BaseURL, "base-url", "", "The API base URL (you should not need to set this)")
	setConfigCmd.MarkFlagRequired("api-key-id")
	setConfigCmd.MarkFlagRequired("api-key-secret")
	setConfigCmd.MarkFlagRequired("account-id")
}
