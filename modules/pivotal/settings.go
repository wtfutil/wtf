package pivotal

import (
	"fmt"
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"os"
)

const (
	defaultFocusable = true
	defaultTitle     = "Pivotal"
)

type customQuery struct {
	title   string `help:"Display title for this query"`
	filter  string `help:"Pivotal search query filter"`
	perPage int    `help:"Number of issues to show"`
	project string `help:"Pivotal project id"`
}

type Settings struct {
	*cfg.Common

	filter        string
	projectId     string
	apiToken      string
	status        string
	customQueries []customQuery `help:"Custom queries allow you to filter pull requests and issues however you like. Give the query a title and a filter. Filters can be copied directly from GitHub’s UI." optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		filter:    ymlConfig.UString("filter", ymlConfig.UString("filter")),
		projectId: ymlConfig.UString("projectId", ymlConfig.UString("projectId", os.Getenv("PIVOTALTRACKER_PROJECT"))),
		apiToken:  ymlConfig.UString("apiToken", ymlConfig.UString("apiToken", os.Getenv("PIVOTALTRACKER_TOKEN"))),
		status:    ymlConfig.UString("status"),
	}

	settings.customQueries = parseCustomQueries(ymlConfig)

	cfg.ModuleSecret(name, globalConfig, &settings.apiToken).Load()

	return &settings
}

func parseCustomQueries(ymlConfig *config.Config) []customQuery {
	var result []customQuery

	if customQueries, err := ymlConfig.Map("customQueries"); err == nil {
		for _, query := range customQueries {
			c := customQuery{}
			for key, value := range query.(map[string]interface{}) {
				switch key {
				case "title":
					c.title = value.(string)
				case "filter":
					c.filter = value.(string)
				case "project":
					switch value := value.(type) {
					case bool, float64, int:
						c.project = fmt.Sprint(value)
					case string:
						c.project = value
					}
				case "perPage":
					c.perPage = value.(int)
				}
			}

			if c.title != "" && c.filter != "" {
				result = append(result, c)
			}
		}
	}
	return result
}
