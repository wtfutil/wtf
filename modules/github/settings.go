package github

import (
	"fmt"
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

const defaultTitle = "GitHub"

type Settings struct {
	common *cfg.Common

	apiKey        string        `help:"Your GitHub API token."`
	baseURL       string        `help:"Your GitHub Enterprise API URL." optional:"true"`
	customQueries []customQuery `help:"Custom queries allow you to filter pull requests and issues however you like. Give the query a title and a filter. Filters can be copied directly from GitHub’s UI." optional:"true"`
	enableStatus  bool          `help:"Display pull request mergeability status (‘dirty’, ‘clean’, ‘unstable’, ‘blocked’)." optional:"true"`
	repositories  []string      `help:"A list of github repositories." values:"Example: wtfutil/wtf"`
	uploadURL     string        `help:"Your GitHub Enterprise upload URL (often the same as API URL). optional:"true"`
	username      string        `help:"Your GitHub username. Used to figure out which review requests you’ve been added to."`
}

type customQuery struct {
	title   string `help:"Display title for this query"`
	filter  string `help:"Github query filter"`
	perPage int    `help:"Number of issues to show"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:       ymlConfig.UString("apiKey", os.Getenv("WTF_GITHUB_TOKEN")),
		baseURL:      ymlConfig.UString("baseURL", os.Getenv("WTF_GITHUB_BASE_URL")),
		enableStatus: ymlConfig.UBool("enableStatus", false),
		uploadURL:    ymlConfig.UString("uploadURL", os.Getenv("WTF_GITHUB_UPLOAD_URL")),
		username:     ymlConfig.UString("username"),
	}
	settings.repositories = parseRepositories(ymlConfig)
	settings.customQueries = parseCustomQueries(ymlConfig)

	return &settings
}

func parseRepositories(ymlConfig *config.Config) []string {

	result := []string{}
	repositories, err := ymlConfig.Map("repositories")
	if err == nil {
		for key, value := range repositories {
			result = append(result, fmt.Sprintf("%s/%s", value, key))
		}
		return result
	}

	result = wtf.ToStrs(ymlConfig.UList("repositories"))
	return result
}

func parseCustomQueries(ymlConfig *config.Config) []customQuery {
	result := []customQuery{}
	if customQueries, err := ymlConfig.Map("customQueries"); err == nil {
		for _, query := range customQueries {
			c := customQuery{}
			for key, value := range query.(map[string]interface{}) {
				switch key {
				case "title":
					c.title = value.(string)
				case "filter":
					c.filter = value.(string)
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
