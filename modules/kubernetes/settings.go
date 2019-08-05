package kubernetes

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const defaultTitle = "Kubernetes"

type Settings struct {
	common *cfg.Common

	objects    []string `help:"Kubernetes objects to show. Options are pods, deployments."`
	title      string   `help:"Override the title of widget."`
	kubeconfig string   `help:"Location of a kubeconfig file."`
	namespaces []string `help:"List of namespaces to watch. If blank, defaults to all namespaces."`
}

func NewSettingsFromYAML(name string, moduleConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, moduleConfig, globalConfig),

		objects:    utils.ToStrs(moduleConfig.UList("objects")),
		title:      moduleConfig.UString("title"),
		kubeconfig: moduleConfig.UString("kubeconfig"),
		namespaces: utils.ToStrs(moduleConfig.UList("namespaces")),
	}

	return &settings
}
