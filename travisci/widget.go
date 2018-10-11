package travisci

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	"strings"
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
	wtf.TextWidget

	builds   *Builds
	selected int
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget(app, "TravisCI", "travisci", true),
	}

	widget.HelpfulWidget.SetView(widget.View)
	widget.unselect()

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	builds, err := BuildsFor()

	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.Name)
		widget.View.SetText(err.Error())
	} else {
		widget.builds = builds
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	if widget.builds == nil {
		return
	}

	widget.View.SetWrap(false)

	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s - Builds", widget.Name)))
	widget.View.SetText(widget.contentFrom(widget.builds))
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
		return wtf.DefaultFocussedRowColor()
	}
	return "White"
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
		travisHost := TRAVIS_HOSTS[wtf.Config.UBool("wtf.mods.travisci.pro", false)]
		wtf.OpenFile(fmt.Sprintf("https://%s/%s/%s/%d", travisHost, build.Repository.Slug, "builds", build.ID))
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
	widget.display()
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
	case "j":
		widget.next()
		return nil
	case "k":
		widget.prev()
		return nil
	case "r":
		widget.Refresh()
		return nil
	}

	switch event.Key() {
	case tcell.KeyDown:
		widget.next()
		return nil
	case tcell.KeyEnter:
		widget.openBuild()
		return nil
	case tcell.KeyEsc:
		widget.unselect()
		return event
	case tcell.KeyUp:
		widget.prev()
		widget.display()
		return nil
	default:
		return event
	}
}
