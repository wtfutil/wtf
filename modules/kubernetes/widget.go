package kubernetes

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Widget contains all the config for the widget
type Widget struct {
	view.TextWidget

	objects    []string
	title      string
	kubeconfig string
	namespaces []string
	context    string
	settings   *Settings
}

// NewWidget creates a new instance of the widget
func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		objects:    settings.objects,
		title:      settings.title,
		kubeconfig: settings.kubeconfig,
		namespaces: settings.namespaces,
		settings:   settings,
		context:    settings.context,
	}

	widget.View.SetWrap(true)

	return &widget
}

// Refresh executes the command and updates the view with the results
func (widget *Widget) Refresh() {
	title := widget.generateTitle()
	client, err := widget.getInstance()

	if err != nil {
		widget.Redraw(func() (string, string, bool) { return title, err.Error(), true })
		return
	}

	var content string

	if utils.Includes(widget.objects, "nodes") {
		nodeList, nodeError := client.getNodes()
		if nodeError != nil {
			widget.Redraw(func() (string, string, bool) { return title, "[red] Error getting node data [white]\n", true })
			return
		}
		content += fmt.Sprintf("[%s]Nodes[white]\n", widget.settings.Colors.Subheading)
		for _, node := range nodeList {
			content += fmt.Sprintf("%s\n", node)
		}
		content += "\n"
	}

	if utils.Includes(widget.objects, "deployments") {
		deploymentList, deploymentError := client.getDeployments(widget.namespaces)
		if deploymentError != nil {
			widget.Redraw(func() (string, string, bool) { return title, "[red] Error getting deployment data [white]\n", true })
			return
		}
		content += fmt.Sprintf("[%s]Deployments[white]\n", widget.settings.Colors.Subheading)
		for _, deployment := range deploymentList {
			content += fmt.Sprintf("%s\n", deployment)
		}
		content += "\n"
	}

	if utils.Includes(widget.objects, "pods") {
		podList, podError := client.getPods(widget.namespaces)
		if podError != nil {
			widget.Redraw(func() (string, string, bool) { return title, "[red] Error getting pod data [white]\n", false })
			return
		}
		content += fmt.Sprintf("[%s]Pods[white]\n", widget.settings.Colors.Subheading)
		for _, pod := range podList {
			content += fmt.Sprintf("%s\n", pod)
		}
		content += "\n"
	}

	widget.Redraw(func() (string, string, bool) { return title, content, false })
}

/* -------------------- Unexported Functions -------------------- */

// generateTitle generates a title for the widget
func (widget *Widget) generateTitle() string {
	if len(widget.title) != 0 {
		return widget.title
	}
	title := "Kube"

	if widget.context != "" {
		title = fmt.Sprintf("%s (%s)", title, widget.context)
	}

	if len(widget.namespaces) == 1 {
		title += fmt.Sprintf(" - Namespace: %s", widget.namespaces[0])
	} else if len(widget.namespaces) > 1 {
		title += fmt.Sprintf(" - Namespaces: %q", widget.namespaces)
	}
	return title
}

// getPods returns a slice of pod strings
func (client *clientInstance) getPods(namespaces []string) ([]string, error) {
	var podList []string
	if len(namespaces) != 0 {
		for _, namespace := range namespaces {
			pods, err := client.Client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
			if err != nil {
				return nil, err
			}

			for _, pod := range pods.Items {
				var podString string
				status := pod.Status.Phase
				name := pod.ObjectMeta.Name
				if len(namespaces) == 1 {
					podString = fmt.Sprintf("%-50s %s", name, status)
				} else {
					podString = fmt.Sprintf("%-20s %-50s %s", namespace, name, status)
				}
				podList = append(podList, podString)
			}
		}
	} else {
		pods, err := client.Client.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		for _, pod := range pods.Items {
			podString := fmt.Sprintf("%-20s %-50s %s", pod.ObjectMeta.Namespace, pod.ObjectMeta.Name, pod.Status.Phase)
			podList = append(podList, podString)
		}
	}

	return podList, nil
}

// get Deployments returns a string slice of pod strings
func (client *clientInstance) getDeployments(namespaces []string) ([]string, error) {
	var deploymentList []string
	if len(namespaces) != 0 {
		for _, namespace := range namespaces {
			deployments, err := client.Client.AppsV1().Deployments(namespace).List(metav1.ListOptions{})
			if err != nil {
				return nil, err
			}

			for _, deployment := range deployments.Items {
				var deployString string
				if len(namespaces) == 1 {
					deployString = fmt.Sprintf("%-50s", deployment.ObjectMeta.Name)
				} else {
					deployString = fmt.Sprintf("%-20s %-50s", deployment.ObjectMeta.Namespace, deployment.ObjectMeta.Name)
				}
				deploymentList = append(deploymentList, deployString)
			}
		}
	} else {
		deployments, err := client.Client.AppsV1().Deployments("").List(metav1.ListOptions{})
		if err != nil {
			return nil, err
		}

		for _, deployment := range deployments.Items {
			deployString := fmt.Sprintf("%-20s %-50s", deployment.ObjectMeta.Namespace, deployment.ObjectMeta.Name)
			deploymentList = append(deploymentList, deployString)
		}
	}
	return deploymentList, nil
}

// getNodes returns a string slice of nodes
func (client *clientInstance) getNodes() ([]string, error) {
	var nodeList []string

	nodes, err := client.Client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, node := range nodes.Items {
		var nodeStatus string
		for _, condition := range node.Status.Conditions {
			if condition.Reason == "KubeletReady" {
				switch {
				case condition.Status == "True":
					nodeStatus = "Ready"
				case condition.Reason == "False":
					nodeStatus = "NotReady"
				default:
					nodeStatus = "Unknown"
				}
			}
		}
		nodeString := fmt.Sprintf("%-50s %s", node.ObjectMeta.Name, nodeStatus)
		nodeList = append(nodeList, nodeString)
	}
	return nodeList, nil
}
