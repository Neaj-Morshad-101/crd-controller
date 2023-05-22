package main

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	_ "k8s.io/code-generator"
	"log"
	"path/filepath"
)

func main() {
	log.Println("Configuring KubeConfig...")

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	log.Printf("kubeconfig : %s", *kubeconfig)

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	//log.Printf("config : %s", config)

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	println(kubeClient)

}
