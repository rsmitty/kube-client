/*
This file implements the pods list subcommand for cobra.
It calls the Kubernetes API to retrieve a list of pods
from a given namespace.
*/

package cmd

import (
	"fmt"

	"github.com/rsmitty/kube-client/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

var podList = &cobra.Command{
	Use:   "list",
	Short: "Lists pods",
	Long:  "Lists pods in a Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			cmd.Help()
		} else {
			kubeConf, err := cmd.Flags().GetString("kubeconfig")
			if err != nil {
				log.Fatal(err)
			}

			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				log.Fatal(err)
			}
			listPods(kubeConf, namespace)
		}
	},
}

func listPods(configPath string, namespace string) {
	clientset := utils.GetClientSet(configPath)
	pods, err := clientset.CoreV1().Pods(namespace).List(v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Here are the pods in %s namespace:\n", namespace)
	for i := 0; i < len(pods.Items); i++ {
		fmt.Printf("#%d: %s\n", i, pods.Items[i].Name)
	}
}
