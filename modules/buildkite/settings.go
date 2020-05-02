package buildkite

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultTitle     = "Buildkite"
	defaultFocusable = true
)

// PipelineSettings defines the configuration properties for a pipeline
type PipelineSettings struct {
	slug     string
	branches []string
}

// Settings defines the configuration properties for this module
type Settings struct {
	common    *cfg.Common
	apiKey    string             `help:"Your Buildkite API Token"`
	orgSlug   string             `help:"Organization Slug"`
	pipelines []PipelineSettings `help:"An array of pipelines to get data from"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:    cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		apiKey:    ymlConfig.UString("apiKey", os.Getenv("WTF_BUILDKITE_TOKEN")),
		orgSlug:   ymlConfig.UString("organizationSlug"),
		pipelines: buildPipelineSettings(ymlConfig),
	}

	cfg.ConfigureSecret(
		globalConfig,
		"",
		name,
		&settings.orgSlug,
		&settings.apiKey,
	)

	return &settings
}

/* -------------------- Unexported Functions -------------------- */

func buildPipelineSettings(ymlConfig *config.Config) []PipelineSettings {
	pipelines := []PipelineSettings{}

	for slug := range ymlConfig.UMap("pipelines") {
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
