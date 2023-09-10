package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile = "./friingo/config.toml"
	debug      = false
)

var rootCmd = &cobra.Command{
	Use:   "friingo",
	Short: "Friingo is an open-source all-in-one server management solution.",
	Long:  "An open-source all-in-one server management solution, built with love by yolocat in Go. Documentation is available at https://friingo.gq/docs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello, world!\nDebug mode: %t\nConfig file: %s\n", debug, configFile)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Friingo",
	Long:  "All software has versions. This is Friingo's",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Get actual version
		fmt.Println("Friingo v0.0.1")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", configFile, fmt.Sprintf("config file (default is %s)", configFile))
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", debug, "enable debug mode")

	viper.SetDefault("debug", false)
	viper.SetDefault("hello", "world")

	rootCmd.AddCommand(versionCmd)
}

func initConfig() {
	err := os.MkdirAll(path.Dir(configFile), 0755)
	cobra.CheckErr(err)

	viper.AddConfigPath(configFile)
	viper.SetConfigType("toml")
	viper.SetConfigName(strings.Split(path.Base(configFile), ".")[0])

	viper.SafeWriteConfig()

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
