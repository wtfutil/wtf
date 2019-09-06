package kubernetes

import (
	"sync"

	"k8s.io/client-go/kubernetes"
	// Includes authentication modules for various Kubernetes providers
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeClient *clientInstance
var clientOnce sync.Once

type clientInstance struct {
	Client kubernetes.Interface
}

// getInstance returns a Kubernetes interface for a clientset
func (widget *Widget) getInstance() *clientInstance {
	clientOnce.Do(func() {
		if kubeClient == nil {
			kubeClient = &clientInstance{
				Client: widget.getKubeClient(),
			}
		}
	})
	return kubeClient
}

// getKubeClient returns a kubernetes clientset for the kubeconfig provided
func (widget *Widget) getKubeClient() kubernetes.Interface {
	config, err := clientcmd.BuildConfigFromFlags("", widget.kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}
