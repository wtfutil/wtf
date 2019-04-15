package jira

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type colors struct {
	rows struct {
		even string
		odd  string
	}
}

type Settings struct {
	colors
	common *cfg.Common

	apiKey                  string
	domain                  string
	email                   string
	jql                     string
	projects                []string
	username                string
	verifyServerCertificate bool
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.jira")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		apiKey:                  localConfig.UString("apiKey", os.Getenv("WTF_JIRA_API_KEY")),
		domain:                  localConfig.UString("domain"),
		email:                   localConfig.UString("email"),
		jql:                     localConfig.UString("jql"),
		username:                localConfig.UString("username"),
		verifyServerCertificate: localConfig.UBool("verifyServerCertificate", true),
	}

	settings.colors.rows.even = localConfig.UString("colors.even", "lightblue")
	settings.colors.rows.odd = localConfig.UString("colors.odd", "white")

	settings.projects = settings.arrayifyProjects(localConfig)

	return &settings
}

/* -------------------- Unexported functions -------------------- */

// arrayifyProjects figures out if we're dealing with a single project or an array of projects
func (settings *Settings) arrayifyProjects(localConfig *config.Config) []string {
	projects := []string{}

	// Single project
	project, err := localConfig.String("project")
	if err == nil {
		projects = append(projects, project)
		return projects
	}

	// Array of projects
	projectList := localConfig.UList("project")
	for _, projectName := range projectList {
		if project, ok := projectName.(string); ok {
			projects = append(projects, project)
		}
	}

	return projects
}
