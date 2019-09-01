package jenkins

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"net/url"
)

type Widget struct {
	view.KeyboardWidget
	view.ScrollableWidget

	settings *Settings
	view     *View
	err      error
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

	view, err := widget.Create(
		widget.settings.url,
		widget.settings.user,
		widget.settings.apiKey,
	)
	widget.view = view

	if err != nil {
		widget.err = err
		widget.SetItemCount(0)
	} else {
		widget.SetItemCount(len(widget.view.Jobs))
	}

	widget.Render()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("%s: [red]%s", widget.CommonSettings().Title, widget.view.Name)
	if widget.err != nil {
		return title, widget.err.Error(), true
	}
	if widget.view == nil || len(widget.view.Jobs) == 0 {
		return title, "No content to display", false
	}

	var str string
	jobs := widget.view.Jobs
	for idx, job := range jobs {
		jobName, _ := url.QueryUnescape(job.Name)

		row := fmt.Sprintf(
			`[%s] [%s]%-6s[white]`,
			widget.RowColor(idx),
			widget.jobColor(&job),
			jobName,
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(job.Name))
	}

	return title, str, false
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

func (widget *Widget) openJob() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.view != nil && sel < len(widget.view.Jobs) {
		job := &widget.view.Jobs[sel]
		utils.OpenFile(job.Url)
	}
}
