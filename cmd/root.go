/*
This file implements the root command for cobra and adds the pods subcommand.
*/

package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kube-client",
	Short: "A toy kubernetes client in golang",
	Long:  "This is a toy kubernetes client implemented in golang. It will support listing pods and deleting them.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logLevel, err := cmd.Flags().GetString("log-level")
		if err != nil {
			log.Fatal(err)
		}
		switch logLevel {
		case "debug":
			log.SetLevel(log.DebugLevel)
		case "error":
			log.SetLevel(log.ErrorLevel)
		default:
			log.SetLevel(log.InfoLevel)
		}
	},
}

func init() {
	RootCmd.PersistentFlags().String("log-level", "info", "A desired logging level. Supported vals: debug, info, error")
	RootCmd.AddCommand(pod)
}
