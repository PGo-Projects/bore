package cmd

import (
	"os"

	"github.com/PGo-Projects/bore/internal/allitebooks/utils"
	tm "github.com/buger/goterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string
var RootCmd = &cobra.Command{Use: "bore"}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		utils.DisplayMessage(err.Error(), tm.RED)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		utils.DisplayMessage("Can't read config: "+err.Error(), tm.RED)
		os.Exit(1)
	}
}
