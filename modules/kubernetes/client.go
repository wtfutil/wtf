package kubernetes

import (
	"sync"

	"k8s.io/client-go/kubernetes"
	// Includes authentication modules for various Kubernetes providers
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeClient *clientInstance
var kubeError error
var clientOnce sync.Once

type clientInstance struct {
	Client kubernetes.Interface
}

// getInstance returns a Kubernetes interface for a clientset
func (widget *Widget) getInstance() (*clientInstance, error) {
	clientOnce.Do(func() {
		if kubeClient == nil {
			client, err := widget.getKubeClient()
			if err != nil {
				kubeError = err
			}
			kubeClient = &clientInstance{
				Client: client,
			}
		}
	})
	return kubeClient, kubeError
}

// getKubeClient returns a kubernetes clientset for the kubeconfig provided
func (widget *Widget) getKubeClient() (kubernetes.Interface, error) {
	config, err := clientcmd.BuildConfigFromFlags("", widget.kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
