package travisci

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.KeyboardWidget
	view.ScrollableWidget

	builds   *Builds
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:   view.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: view.NewScrollableWidget(app, settings.common, true),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	builds, err := BuildsFor(widget.settings.apiKey, widget.settings.pro)

	if err != nil {
		widget.Redraw(widget.CommonSettings().Title, err.Error(), true)
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

	title := fmt.Sprintf("%s - Builds", widget.CommonSettings().Title)
	widget.Redraw(title, widget.contentFrom(widget.builds), false)
}

func (widget *Widget) contentFrom(builds *Builds) string {
	var str string
	for idx, build := range builds.Builds {

		row := fmt.Sprintf(
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
		str += utils.HighlightableHelper(widget.View, row, idx, len(build.Branch.Name))
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
		utils.OpenFile(fmt.Sprintf("https://%s/%s/%s/%d", travisHost, build.Repository.Slug, "builds", build.ID))
	}
}
