/*
This file implements the pods delete subcommand for cobra
and defines the pods list and pods delete subcommand.
*/

package cmd

import (
	"fmt"

	"github.com/rsmitty/kube-client/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

var podDelete = &cobra.Command{
	Use:   "delete NAME",
	Short: "Deletes pods",
	Long:  "Deletes pods in a Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
		} else {
			kubeConf, err := cmd.Flags().GetString("kubeconfig")
			if err != nil {
				panic(err)
			}

			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				panic(err)
			}
			deletePod(kubeConf, namespace, args[0])
		}
	},
}

func deletePod(configPath string, namespace string, podName string) {
	clientset := utils.GetClientSet(configPath)

	err := clientset.CoreV1().Pods(namespace).Delete(podName, &v1.DeleteOptions{})
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Printf("Successfully submitted deletion for pod %s\n", podName)
	}
}
