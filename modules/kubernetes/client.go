package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	// Includes authentication modules for various Kubernetes providers
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

type clientInstance struct {
	Client kubernetes.Interface
}

// getInstance returns a Kubernetes interface for a clientset
func (widget *Widget) getInstance() (*clientInstance, error) {
	var err error

	widget.clientOnce.Do(func() {
		widget.client = &clientInstance{}
		widget.client.Client, err = widget.getKubeClient()
	})

	return widget.client, err
}

// getKubeClient returns a kubernetes clientset for the kubeconfig provided
func (widget *Widget) getKubeClient() (kubernetes.Interface, error) {
	var overrides *clientcmd.ConfigOverrides
	if widget.context != "" {
		overrides = &clientcmd.ConfigOverrides{
			CurrentContext: widget.context,
		}
	}

	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: widget.kubeconfig},
		overrides).ClientConfig()

	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
