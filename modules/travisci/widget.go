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
	wtf.TextWidget

	builds   *Builds
	selected int
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:  wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget: wtf.NewKeyboardWidget(),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.unselect()

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
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
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
			widget.rowColor(idx),
			buildColor(&build),
			build.Repository.Name,
			build.Number,
			build.Branch.Name,
			widget.rowColor(idx),
			strings.Split(build.Commit.Message, "\n")[0],
			build.CreatedBy.Login,
		)
	}

	return str
}

func (widget *Widget) rowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		return widget.settings.common.DefaultFocussedRowColor()
	}

	return widget.settings.common.RowColor(idx)
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

func (widget *Widget) next() {
	widget.selected++
	if widget.builds != nil && widget.selected >= len(widget.builds.Builds) {
		widget.selected = 0
	}

	widget.display()
}

func (widget *Widget) prev() {
	widget.selected--
	if widget.selected < 0 && widget.builds != nil {
		widget.selected = len(widget.builds.Builds) - 1
	}

	widget.display()
}

func (widget *Widget) openBuild() {
	sel := widget.selected
	if sel >= 0 && widget.builds != nil && sel < len(widget.builds.Builds) {
		build := &widget.builds.Builds[widget.selected]
		travisHost := TRAVIS_HOSTS[widget.settings.pro]
		wtf.OpenFile(fmt.Sprintf("https://%s/%s/%s/%d", travisHost, build.Repository.Slug, "builds", build.ID))
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
	widget.display()
}
