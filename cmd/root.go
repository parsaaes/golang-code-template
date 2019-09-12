package cmd

import (
	"os"

	"github.com/rashadansari/golang-code-template/buildinfo"
	"github.com/rashadansari/golang-code-template/cmd/migrate"
	"github.com/rashadansari/golang-code-template/cmd/server"
	"github.com/rashadansari/golang-code-template/config"
	"github.com/rashadansari/golang-code-template/log"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// ExitFailure status code
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := buildinfo.Validate(); err != nil {
		logrus.Error(err.Error())
	}

	cfg := config.Init("config.yaml")

	log.SetupLogger(cfg.Logger)

	var root = &cobra.Command{
		Use:   "template",
		Short: "Template Server",
	}

	server.Register(root, cfg)
	migrate.Register(root, cfg)

	if err := root.Execute(); err != nil {
		logrus.Errorf("failed to execute root command: %s", err.Error())
		os.Exit(ExitFailure)
	}
}
