package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/olebedev/config"
	"github.com/rivo/tview"

	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/pluggable"
	"github.com/wtfutil/wtf/view"
	"github.com/wtfutil/wtf/wtf"
)

const (
	defaultFocusable = false
	defaultTitle     = "Harambe"
)

type Settings struct {
	*cfg.Common
}

func newSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
	}

	return &settings
}

type ExamplePluggable struct {
	view.TextWidget

	settings *Settings
}

func newWidget(tviewApp *tview.Application, settings *Settings) *ExamplePluggable {
	widget := ExamplePluggable{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Initialize is specifically exported to satisfy the wtf.pluggable.ModulePluggable interface
// and make it loadable by said package.
// This is hacky and I'm so sorry.
func (widget *ExamplePluggable) Initialize(
	name string,
	ymlConfig *config.Config,
	globalConfig *config.Config,
	tviewApp *tview.Application,
	_ *tview.Pages,
) wtf.Wtfable {
	settings := newSettingsFromYAML(name, ymlConfig, globalConfig)
	w := newWidget(tviewApp, settings)
	*widget = *w

	return widget
}

func (widget *ExamplePluggable) Refresh() {
	widget.Redraw(widget.getText)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *ExamplePluggable) getText() (string, string, bool) {
	text := "Pour one out for Harambe"

	return widget.CommonSettings().Title, text, false
}

/* -------------------- net/rpc hackery -------------------- */

type pluggableRPCServer struct{}

func (p *pluggableRPCServer) Module() (pluggable.ExternalModule, error) {
	return &ExamplePluggable{}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		Plugins: map[string]plugin.Plugin{
			pluggable.RPCPluginName: &pluggable.RPCPluggablePlugin{Impl: &pluggableRPCServer{}},
		},
	})
}
