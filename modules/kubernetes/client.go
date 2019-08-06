package kubernetes

import (
	"sync"

	"k8s.io/client-go/kubernetes"
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
	config, _ := clientcmd.BuildConfigFromFlags("", widget.kubeconfig)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error)
	}
	return clientset
}
