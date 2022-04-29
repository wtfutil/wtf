package asana

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type TaskType int

const (
	TASK_TYPE TaskType = iota
	TASK_SECTION
	TASK_BREAK
)

type TaskItem struct {
	name        string
	numSubtasks int32
	dueOn       string
	id          string
	url         string
	taskType    TaskType
	completed   bool
	assignee    string
}

type Widget struct {
	view.ScrollableWidget

	tasks []*TaskItem

	mu       sync.Mutex
	err      error
	settings *Settings
	tviewApp *tview.Application
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),

		tviewApp: tviewApp,
		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.tasks = nil
	widget.err = nil
	widget.SetItemCount(0)

	widget.mu.Lock()
	defer widget.mu.Unlock()
	tasks, err := widget.Fetch(
		widget.settings.workspaceId,
		widget.settings.projectId,
		widget.settings.mode,
		widget.settings.sections,
		widget.settings.allUsers,
	)
	if err != nil {
		widget.err = err
	} else {
		widget.tasks = tasks
		widget.SetItemCount(len(tasks))
	}

	widget.Render()
}

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

func (widget *Widget) Fetch(workspaceId, projectId, mode string, sections []string, allUsers bool) ([]*TaskItem, error) {

	availableModes := map[string]interface{}{
		"project":          nil,
		"project_sections": nil,
		"workspace":        nil,
	}

	if _, ok := availableModes[mode]; !ok {
		return nil, fmt.Errorf("missing mode, or mode is invalid - please set to project, project_sections or workspace")
	}

	if widget.settings.apiKey != "" {
		widget.settings.token = widget.settings.apiKey
	} else {
		widget.settings.token = os.Getenv("WTF_ASANA_TOKEN")
	}

	if widget.settings.token == "" {
		return nil, fmt.Errorf("missing environment variable token or apikey config")
	}

	subMode := mode
	if allUsers && mode != "workspace" {
		subMode += "_all"
	}

	if projectId == "" && strings.HasPrefix(subMode, "project") {
		return nil, fmt.Errorf("missing project id")
	}

	if workspaceId == "" && subMode == "workspace" {
		return nil, fmt.Errorf("missing workspace id")
	}

	var tasks []*TaskItem
	var err error

	switch {
	case strings.HasPrefix(subMode, "project_sections"):
		tasks, err = fetchTasksFromProjectSections(widget.settings.token, projectId, sections, subMode)
	case strings.HasPrefix(subMode, "project"):
		tasks, err = fetchTasksFromProject(widget.settings.token, projectId, subMode)
	case subMode == "workspace":
		tasks, err = fetchTasksFromWorkspace(widget.settings.token, workspaceId, subMode)
	default:
		err = fmt.Errorf("no mode found")
	}

	if err != nil {
		return nil, err
	}

	return tasks, nil

}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {

	title := widget.CommonSettings().Title
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	data := widget.tasks
	if len(data) == 0 {
		return title, "No data", false
	}

	var str string

	for idx, taskItem := range data {
		switch {
		case taskItem.taskType == TASK_TYPE:
			if widget.settings.hideComplete && taskItem.completed {
				continue
			}

			rowColor := widget.RowColor(idx)

			completed := "[ []"
			if taskItem.completed {
				completed = "[x[]"
			}

			row := ""

			if widget.settings.allUsers && taskItem.assignee != "" {
				row = fmt.Sprintf(
					"[%s]  %s %s: %s",
					rowColor,
					completed,
					taskItem.assignee,
					taskItem.name,
				)
			} else {
				row = fmt.Sprintf(
					"[%s]  %s %s",
					rowColor,
					completed,
					taskItem.name,
				)
			}

			if taskItem.numSubtasks > 0 {
				row += fmt.Sprintf(" (%d)", taskItem.numSubtasks)
			}

			if taskItem.dueOn != "" {
				row += fmt.Sprintf(" due: %s", taskItem.dueOn)
			}

			row += " [white]"

			str += utils.HighlightableHelper(widget.View, row, idx, len(taskItem.name))

		case taskItem.taskType == TASK_SECTION:
			if idx > 1 {
				row := "[white] "

				str += utils.HighlightableHelper(widget.View, row, idx, len(taskItem.name))
			}
			row := fmt.Sprintf(
				"[white] %s [white]",
				taskItem.name,
			)

			str += utils.HighlightableHelper(widget.View, row, idx, len(taskItem.name))

			row = "[white] "

			str += utils.HighlightableHelper(widget.View, row, idx, len(taskItem.name))

		}

	}

	return title, str, false

}

func (widget *Widget) openTask() {
	sel := widget.GetSelected()

	if sel >= 0 && widget.tasks != nil && sel < len(widget.tasks) {
		task := widget.tasks[sel]
		if task.taskType == TASK_TYPE && task.url != "" {
			utils.OpenFile(task.url)
		}
	}
}

func (widget *Widget) toggleTaskCompletion() {
	sel := widget.GetSelected()

	if sel >= 0 && widget.tasks != nil && sel < len(widget.tasks) {
		task := widget.tasks[sel]
		if task.taskType == TASK_TYPE {
			widget.mu.Lock()

			err := toggleTaskCompletionById(widget.settings.token, task.id)
			if err != nil {
				widget.err = err
			}

			widget.mu.Unlock()
			widget.Refresh()
		}
	}
}
