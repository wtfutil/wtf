package cdsfavorites

import (
	"fmt"

	"github.com/ovh/cds/sdk"
)

func (widget *Widget) display() {
	widget.TextWidget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	if len(widget.View.GetHighlights()) > 0 {
		widget.View.ScrollToHighlight()
	} else {
		widget.View.ScrollToBeginning()
	}

	widget.Items = make([]int64, 0)

	workflow := widget.currentCDSWorkflow()
	if workflow == nil {
		return "", " Workflow not selected ", false
	}

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.common.SigilStr(len(widget.workflows), widget.Idx, width) + "\n"
	title := fmt.Sprintf("%s - %s", widget.CommonSettings().Title, widget.title(workflow))

	str += widget.displayWorkflowRuns(workflow)

	return title, str, false
}

func (widget *Widget) title(workflow *sdk.Workflow) string {
	return fmt.Sprintf(
		"[%s]%s/%s[white]",
		widget.settings.common.Colors.TextTheme.Title,
		workflow.ProjectKey, workflow.Name,
	)
}

func (widget *Widget) displayWorkflowRuns(workflow *sdk.Workflow) string {
	runs, _ := widget.client.WorkflowRunList(workflow.ProjectKey, workflow.Name, 0, 16)

	widget.SetItemCount(len(runs))

	if len(runs) == 0 {
		return " [grey]none[white]\n"
	}

	content := ""
	for idx, run := range runs {
		var tags string
		for _, tag := range run.Tags {
			toadd := true
			for _, v := range widget.settings.hideTags {
				if v == tag.Tag {
					toadd = false
					break
				}
			}
			if toadd {
				tags = fmt.Sprintf("%s%s:%s ", tags, tag.Tag, tag.Value)
			}
		}
		content += fmt.Sprintf(`[%s]["%d"]%d %-6s[""][gray] %s`, getStatusColor(run.Status), idx, run.Number, run.Status, tags)
		content += "\n"
		widget.Items = append(widget.Items, run.Number)
	}

	return content
}

func getStatusColor(status string) string {
	switch status {
	case sdk.StatusSuccess:
		return "green"
	case sdk.StatusBuilding, sdk.StatusWaiting:
		return "blue"
	case sdk.StatusFail:
		return "red"
	case sdk.StatusStopped:
		return "red"
	case sdk.StatusSkipped:
		return "grey"
	case sdk.StatusDisabled:
		return "grey"
	}
	return "red"
}
