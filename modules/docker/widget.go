package docker

import (

	// "github.com/docker/docker/client"
	"github.com/docker/docker/client"
	"github.com/olebedev/config"
	"github.com/pkg/errors"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
)

const defaultTitle = "docker"

type Settings struct {
	common     *cfg.Common
	labelColor string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common:     cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),
		labelColor: ymlConfig.UString("labelColor", "labelColor"),
	}

	return &settings
}

type Widget struct {
	view.TextWidget
	cli           *client.Client
	settings      *Settings
	displayBuffer string
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common, false),
		settings:   settings,
	}

	widget.View.SetScrollable(true)
	// widget.View.SetRegions(true)

	// os.Setenv("DOCKER_HOST", "unix:///var/run/docker.sock")
	// os.Setenv("DOCKER_API_VERSION", "1.23")

	cli, err := client.NewEnvClient()
	if err != nil {
		widget.displayBuffer = errors.Wrap(err, "could not create client").Error()
	} else {
		widget.cli = cli
	}

	widget.refreshDisplayBuffer()

	return &widget
}

func (widget *Widget) Refresh() {
	widget.refreshDisplayBuffer()
	widget.Redraw(widget.display)
}

func (widget *Widget) display() (string, string, bool) {
	return widget.CommonSettings().Title, widget.displayBuffer, true
}

func (widget *Widget) refreshDisplayBuffer() {

	if widget.cli == nil {
		return
	}

	widget.displayBuffer = ""

	widget.displayBuffer += "[" + widget.settings.labelColor + "::bul]system\n"
	widget.displayBuffer += widget.getSystemInfo()

	widget.displayBuffer += "\n"

	widget.displayBuffer += "[" + widget.settings.labelColor + "::bul]containers\n"
	widget.displayBuffer += widget.getContainerStates()

}
