package urlcheck

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

// makeConfig creates a *config.Config from a YAML string for testing.
func makeConfig(yaml string) (*config.Config, error) {
	return config.ParseYaml(yaml)
}

// makeMinimalCommon creates a minimal Common settings struct for tests that need one.
func makeMinimalCommon() *cfg.Common {
	ymlCfg, _ := config.ParseYaml("enabled: true\n")
	globalCfg, _ := config.ParseYaml("wtf:\n  colors:\n    border:\n      focusable: darkslateblue\n      focused: orange\n      normal: gray\n")
	return cfg.NewCommonSettingsFromModule("urlcheck", "URLcheck", false, ymlCfg, globalCfg)
}
