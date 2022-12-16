/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rofi-go/cli/config/cli"
	"github.com/rofi-go/cli/internal/cli/domain/ports/input"
	"github.com/rofi-go/cli/internal/cli/domain/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var appSvc input.AppUseCase

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		appSvc.StartApp()
	},
}

var daemon bool

var _log *logrus.Logger
var v *viper.Viper
var configs *cli.AppConfig

func init() {

	_log = logrus.New()
	if os.Getenv("ENV") == "dev" {
		_log.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	} else {
		_log.SetFormatter(&logrus.JSONFormatter{})
		_log.SetOutput(os.Stdout)
	}
	_log.Out = os.Stdout

	v = viper.New()
	_configs := cli.NewConfig(v)

	env := os.Getenv("ENV")
	appConfig, err := _configs.GetConfig(env)
	configs = appConfig
	if err != nil {
		return
	}

	appSvc = service.NewAppService(_log.WithFields(logrus.Fields{"service": "app"}), configs)
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().BoolVarP(&daemon, "daemon", "d", false, "run as daemon")

}
