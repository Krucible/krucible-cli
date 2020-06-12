package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "krucible",
	Short: "The official command line client for Krucible",
	Long: `The official command line client for Krucible.
Krucible is a platform for creating Kubernetes clusters optimised for testing
and development. See https://usekrucible.com for more.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

func getConfigDirOrDie() string {
	rootConfigDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: could not get user config dir")
		panic(err)
	}
	configDir := path.Join(rootConfigDir, "krucible")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		panic(err)
	}
	return configDir
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with name ".krucible" (without extension).
		viper.AddConfigPath(getConfigDirOrDie())
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("krucible")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	viper.ReadInConfig()
}
