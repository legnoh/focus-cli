package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version string
	debug   bool
	log     *logrus.Logger
)

var rootCmd = &cobra.Command{
	Use:     "focus",
	Version: version,
	Short:   "Control Focus mode on macOS",
	Long:    `Control Focus mode on macOS via CLI`,
}

func Execute() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	if rootCmd.Execute() != nil {
		log.Fatal("Root execute is failed... exit")
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "print debug log")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	debug, err := rootCmd.PersistentFlags().GetBool("debug")
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		log.SetLevel(logrus.DebugLevel)
	}
}
