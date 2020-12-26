package travisci

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.ScrollableWidget

	builds   *Builds
	settings *Settings
	err      error
}

func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, pages, settings.Common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	builds, err := BuildsFor(widget.settings)

	if err != nil {
		widget.err = err
		widget.builds = nil
		widget.SetItemCount(0)
	} else {
		widget.err = nil
		widget.builds = builds
		widget.SetItemCount(len(builds.Builds))
	}
	widget.Render()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("%s - Builds", widget.CommonSettings().Title)
	var str string
	if widget.err != nil {
		str = widget.err.Error()
	} else {
		var rowFormat = "[%s] [%s] %s-%s (%s) [%s]%s - [blue]%s"
		if !widget.settings.compact {
			rowFormat += "\n"
		}

		for idx, build := range widget.builds.Builds {
			row := fmt.Sprintf(
				rowFormat,
				widget.RowColor(idx),
				buildColor(build),
				build.Repository.Name,
				build.Number,
				build.Branch.Name,
				widget.RowColor(idx),
				strings.Split(build.Commit.Message, "\n")[0],
				build.CreatedBy.Login,
			)
			str += utils.HighlightableHelper(widget.View, row, idx, len(build.Branch.Name))
		}
	}

	return title, str, false
}

func buildColor(build Build) string {
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
