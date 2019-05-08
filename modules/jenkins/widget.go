package jenkins

import (
	"fmt"
	"strconv"

	"regexp"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

const HelpText = `
 Keyboard commands for Jenkins:

   /: Show/hide this help window
   j: Select the next job in the list
   k: Select the previous job in the list
   r: Refresh the data

   arrow down: Select the next job in the list
   arrow up:   Select the previous job in the list

   return: Open the selected job in a browser
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.TextWidget

	app *tview.Application

	selected int
	settings *Settings
	view     *View
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:  wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget: wtf.NewKeyboardWidget(),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		app:      app,
		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.unselect()

	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)

	widget.HelpfulWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	view, err := widget.Create(
		widget.settings.url,
		widget.settings.user,
		widget.settings.apiKey,
	)
	widget.view = view

	if err != nil {
		widget.Redraw(widget.CommonSettings.Title, err.Error(), true)
		return
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	if widget.view == nil {
		return
	}

	title := fmt.Sprintf("%s: [red]%s", widget.CommonSettings.Title, widget.view.Name)

	widget.Redraw(title, widget.contentFrom(widget.view), false)
	widget.app.QueueUpdateDraw(func() {
		widget.View.Highlight(strconv.Itoa(widget.selected)).ScrollToHighlight()
	})
}

func (widget *Widget) contentFrom(view *View) string {
	var str string
	for idx, job := range view.Jobs {
		var validID = regexp.MustCompile(widget.settings.jobNameRegex)

		if validID.MatchString(job.Name) {
			str = str + fmt.Sprintf(
				`["%d"][%s] [%s]%-6s[white][""]`,
				idx,
				widget.rowColor(idx),
				widget.jobColor(&job),
				job.Name,
			)

			str = str + "\n"
		}
	}

	return str
}

func (widget *Widget) rowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		return widget.settings.common.DefaultFocussedRowColor()
	}

	return widget.settings.common.RowColor(idx)
}

func (widget *Widget) jobColor(job *Job) string {
	switch job.Color {
	case "blue":
		// Override color if successBallColor boolean param provided in config
		return widget.settings.successBallColor
	case "red":
		return "red"
	default:
		return "white"
	}
}

func (widget *Widget) next() {
	widget.selected++
	if widget.view != nil && widget.selected >= len(widget.view.Jobs) {
		widget.selected = 0
	}

	widget.display()
}

func (widget *Widget) prev() {
	widget.selected--
	if widget.selected < 0 && widget.view != nil {
		widget.selected = len(widget.view.Jobs) - 1
	}

	widget.display()
}

func (widget *Widget) openJob() {
	sel := widget.selected
	if sel >= 0 && widget.view != nil && sel < len(widget.view.Jobs) {
		job := &widget.view.Jobs[widget.selected]
		wtf.OpenFile(job.Url)
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
	widget.display()
}
