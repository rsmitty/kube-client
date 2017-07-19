/*
This file is here to put commonly used Kubernetes functions into a single spot.
*/

package utils

import (
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//GCP package imported to allow for auth with GKE clusters if that's your current context
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

//GetClientSet creates a Kubernetes config from the config file passed in.
//Returns a pointer to a clientset that can be used for API requests.
func GetClientSet(configPath string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return clientset
}
