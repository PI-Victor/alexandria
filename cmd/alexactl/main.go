package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/PI-Victor/alexandria/pkg/cli"
)

var rootCmd = &cobra.Command{
	Use:   "alexctl",
	Short: "alexctl - CLI control alexandria libary images",
	Example: ` alexctl provides functionality of manipulating images stored in the Alexandria library.
  `,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(cli.PullImages)
	rootCmd.AddCommand(cli.ListImages)
	rootCmd.Execute()
}

func init() {
	// NOTE: using $HOME might not be a good idea.
	viper.AddConfigPath("$HOME/.alexandria")
	viper.SetConfigFile("config")
	logrus.SetLevel(logrus.WarnLevel)
}
