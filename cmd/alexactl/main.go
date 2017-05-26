package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/PI-Victor/alexandria/pkg/cli"
)

var rootCmd = &cobra.Command{
	Use:   "alexactl",
	Short: "alexactl - CLI control Alexandria libary images",
	Example: `alexactl provides functionality of manipulating images stored in the Alexandria image
library.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(cli.PullImages)
	rootCmd.AddCommand(cli.ListImages)
	rootCmd.AddCommand(cli.ImportImages)
	rootCmd.AddCommand(cli.Image)
	rootCmd.Execute()
}

func init() {
	// NOTE: using $HOME might not be a good idea.
	logrus.SetLevel(logrus.DebugLevel)

	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath("/home/vpalade/.alexandria")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logrus.Errorf("Fatal error config file: %s \n", err)
	}
}
