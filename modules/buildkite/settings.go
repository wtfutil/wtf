package buildkite

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
	"os"
)

type PipelineSettings struct {
	slug     string
	branches []string
}

type Settings struct {
	common    *cfg.Common
	apiKey    string             `help:"Your Buildkite API Token"`
	orgSlug   string             `help:"Organization Slug"`
	pipelines []PipelineSettings `help:"An array of pipelines to get data from"`
}

const defaultTitle = "Buildkite"
const defaultFocusable = true

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:    cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		apiKey:    ymlConfig.UString("apiKey", os.Getenv("WTF_BUILDKITE_TOKEN")),
		orgSlug:   ymlConfig.UString("organizationSlug"),
		pipelines: buildPipelineSettings(ymlConfig),
	}

	return &settings
}

func buildPipelineSettings(ymlConfig *config.Config) []PipelineSettings {
	pipelines := []PipelineSettings{}

	for slug, _ := range ymlConfig.UMap("pipelines") {
		branches := utils.ToStrs(ymlConfig.UList("pipelines." + slug + ".branches"))
		if len(branches) == 0 {
			branches = []string{"master"}
		}

		pipeline := PipelineSettings{
			slug:     slug,
			branches: branches,
		}

		pipelines = append(pipelines, pipeline)
	}

	return pipelines
}
