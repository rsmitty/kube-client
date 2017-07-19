/*
This file implements the pods subcommand for cobra
and defines the pods list and pods delete subcommand.

It is also responsible for adding the kubeconfig and
namespace flags to each subcommand.
*/

package cmd

import (
	"os/user"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var pod = &cobra.Command{
	Use:   "pod SUBCOMMAND",
	Short: "Allows for interaction with pods",
	Long:  "Allows for listing and deleting pods in a Kubernetes cluster",
}

func init() {
	//Grab user's home directory to use in defaults below
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := user.HomeDir

	//Iterate through the subcommands we want to add, as well as provide default flags for kubeconfig and namespace
	subCommands := []*cobra.Command{podList, podDelete}
	for _, subCommand := range subCommands {
		subCommand.Flags().String("kubeconfig", home+"/.kube/config", "A path to your kubeconfig file.")
		subCommand.Flags().String("namespace", "default", "A kubernetes namespace you would like to interact with.")
		pod.AddCommand(subCommand)
	}

}
