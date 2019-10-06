package buildkite

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
	"strings"
)

const HelpText = `
 Keyboard commands for Buildkite:
`

type Widget struct {
	view.KeyboardWidget
	view.TextWidget
	settings *Settings

	builds []Build
	err    error
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget: view.NewKeyboardWidget(app, pages, settings.common),
		TextWidget:     view.NewTextWidget(app, settings.common),
		settings:       settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)
	widget.View.SetScrollable(true)
	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	builds, err := widget.getBuilds()

	if err != nil {
		widget.err = err
		widget.builds = nil
	} else {
		widget.builds = builds
		widget.err = nil
	}

	// The last call should always be to the display function
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("%s - [green]%s", widget.CommonSettings().Title, widget.settings.orgSlug)

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	pipelineData := groupByPipeline(widget.builds)
	maxPipelineLength := getLongestPipelineLength(widget.builds)

	str := ""
	for pipeline, builds := range pipelineData {
		str += fmt.Sprintf("[white]%s", padRight(pipeline, maxPipelineLength))
		for _, build := range builds {
			str += fmt.Sprintf("  [%s]%s[white]", buildColor(build.State), build.Branch)
		}
		str += "\n"
	}

	return title, str, false
}

func groupByPipeline(builds []Build) map[string][]Build {
	grouped := make(map[string][]Build)

	for _, build := range builds {
		if _, ok := grouped[build.Pipeline.Slug]; ok {
			grouped[build.Pipeline.Slug] = append(grouped[build.Pipeline.Slug], build)
		} else {
			grouped[build.Pipeline.Slug] = []Build{}
			grouped[build.Pipeline.Slug] = append(grouped[build.Pipeline.Slug], build)
		}
	}

	return grouped
}

func getLongestPipelineLength(builds []Build) int {
	maxPipelineLength := 0

	for _, build := range builds {
		if len(build.Pipeline.Slug) > maxPipelineLength {
			maxPipelineLength = len(build.Pipeline.Slug)
		}
	}

	return maxPipelineLength
}

func padRight(text string, length int) string {
	padLength := length - len(text)

	if padLength <= 0 {
		return text[:length]
	}

	return text + strings.Repeat(" ", padLength)
}

func buildColor(state string) string {
	switch state {
	case "passed":
		return "green"
	case "failed":
		return "red"
	default:
		return "yellow"
	}
}
