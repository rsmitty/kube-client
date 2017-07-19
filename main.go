package main

import (
	"os"

	"github.com/rsmitty/kube-client/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
