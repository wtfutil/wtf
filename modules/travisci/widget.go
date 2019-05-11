package travisci

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

const HelpText = `
 Keyboard commands for Travis CI:

   /: Show/hide this help window
   j: Select the next build in the list
   k: Select the previous build in the list
   r: Refresh the data

   arrow down: Select the next build in the list
   arrow up:   Select the previous build in the list

   return: Open the selected build in a browser
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.ScrollableWidget

	builds   *Builds
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:    wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget:   wtf.NewKeyboardWidget(),
		ScrollableWidget: wtf.NewScrollableWidget(app, settings.common, true),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.HelpfulWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	builds, err := BuildsFor(widget.settings.apiKey, widget.settings.pro)

	if err != nil {
		widget.Redraw(widget.CommonSettings.Title, err.Error(), true)
		return
	}
	widget.builds = builds
	widget.SetItemCount(len(builds.Builds))
	widget.Render()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) Render() {
	if widget.builds == nil {
		return
	}

	title := fmt.Sprintf("%s - Builds", widget.CommonSettings.Title)
	widget.Redraw(title, widget.contentFrom(widget.builds), false)
}

func (widget *Widget) contentFrom(builds *Builds) string {
	var str string
	for idx, build := range builds.Builds {

		str = str + fmt.Sprintf(
			"[%s] [%s] %s-%s (%s) [%s]%s - [blue]%s\n",
			widget.RowColor(idx),
			buildColor(&build),
			build.Repository.Name,
			build.Number,
			build.Branch.Name,
			widget.RowColor(idx),
			strings.Split(build.Commit.Message, "\n")[0],
			build.CreatedBy.Login,
		)
	}

	return str
}

func buildColor(build *Build) string {
	switch build.State {
	case "broken":
		return "red"
	case "failed":
		return "red"
	case "failing":
		return "red"
	case "pending":
		return "yellow"
	case "started":
		return "yellow"
	case "fixed":
		return "green"
	case "passed":
		return "green"
	default:
		return "white"
	}
}

func (widget *Widget) openBuild() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.builds != nil && sel < len(widget.builds.Builds) {
		build := &widget.builds.Builds[sel]
		travisHost := TRAVIS_HOSTS[widget.settings.pro]
		wtf.OpenFile(fmt.Sprintf("https://%s/%s/%s/%d", travisHost, build.Repository.Slug, "builds", build.ID))
	}
}
