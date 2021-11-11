package main

import (
	"fmt"
	"path/filepath"
	"flag"
	"context"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	serviceNames := flag.String("sn", "", "service names separated by comma")
	nameSpace := flag.String("ns", "", "namespace")
	wildCard := flag.String("wc", "", "wildcard to match portion of the key")
	varEnable := flag.Bool("var", false, "check configmap vars as well")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	for _, service := range strings.Split(*serviceNames, ",") {
		secret, err := clientset.CoreV1().Secrets(*nameSpace).Get(context.TODO(), service, metav1.GetOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n",service)
		for key, value := range secret.Data {
			if strings.Contains(key,*wildCard) {
				fmt.Printf("    %s: %s\n", key, value)
				}
		}
		if *varEnable {
			configmap, err := clientset.CoreV1().ConfigMaps(*nameSpace).Get(context.TODO(), service, metav1.GetOptions{})
			if err != nil {
				panic(err)
			}
			for key, value := range configmap.Data {
				if strings.Contains(key,*wildCard) {
					fmt.Printf("    %s: %s\n", key, value)
					}
				}
		}
	}
}