package asana

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
	defaultTitle     = "Asana"
)

type Settings struct {
	*cfg.Common

	projectId string `help:"The Asana Project ID. If the mode is 'project' or 'project_sections' this is required to known which Asana Project to pull your tasks from" values:"A valid Asana Project ID string" optional:"true"`

	workspaceId string `help:"The Asana Workspace ID. If mode is 'workspace' this is required" values:"A valid Asana Workspace ID string" optional:"true"`

	sections []string `help:"The Asana Section Labels to fetch from the Project. Required if the mode is 'project_sections'" values:"An array of Asana Section Label strings" optional:"true"`

	allUsers bool `help:"Fetch tasks for all users, defaults to false" values:"bool" optional:"true"`

	mode string `help:"What mode to query Asana, 'project', 'project_sections', 'workspace'" values:"A string with either 'project', 'project_sections' or 'workspace'"`

	hideComplete bool `help:"Hide completed tasks, defaults to false" values:"bool" optional:"true"`

	apiKey string `help:"Your Asana Personal Access Token. Leave this blank to use the WTF_ASANA_TOKEN environment variable." values:"Your Asana Personal Access Token as a string" optional:"true"`

	token string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig,
			globalConfig),
		projectId:    ymlConfig.UString("projectId", ""),
		apiKey:       ymlConfig.UString("apiKey", ""),
		workspaceId:  ymlConfig.UString("workspaceId", ""),
		sections:     utils.ToStrs(ymlConfig.UList("sections")),
		allUsers:     ymlConfig.UBool("allUsers", false),
		mode:         ymlConfig.UString("mode", ""),
		hideComplete: ymlConfig.UBool("hideComplete", false),
	}

	return &settings
}
