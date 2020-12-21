package gitlabtodo

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	gitlab "github.com/xanzy/go-gitlab"
)

type Widget struct {
	view.ScrollableWidget

	todos        []*gitlab.Todo
	gitlabClient *gitlab.Client
	settings     *Settings
	err          error
}

func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, pages, settings.Common),

		settings: settings,
	}

	widget.gitlabClient, _ = gitlab.NewClient(settings.apiKey, gitlab.WithBaseURL(settings.domain))

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	todos, err := widget.getTodos(widget.settings.apiKey)
	widget.todos = todos
	widget.err = err
	widget.SetItemCount(len(todos))

	widget.Render()
}

// Render sets up the widget data for redrawing to the screen
func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("GitLab ToDos (%d)", len(widget.todos))

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if widget.todos == nil {
		return title, "No ToDos to display", false
	}

	str := widget.contentFrom(widget.todos)

	return title, str, false
}

func (widget *Widget) getTodos(apiKey string) ([]*gitlab.Todo, error) {
	opts := gitlab.ListTodosOptions{}

	todos, _, err := widget.gitlabClient.Todos.ListTodos(&opts)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// trim the todo body so it fits on a single line
func (widget *Widget) trimTodoBody(body string) string {
	r := []rune(body)

	// Cut at first occurence of a newline
	for i, a := range r {
		if a == '\n' {
			return string(r[:i])
		}
	}

	return body
}

func (widget *Widget) contentFrom(todos []*gitlab.Todo) string {
	var str string

	for idx, todo := range todos {
		row := fmt.Sprintf(`[%s]%2d. `, widget.RowColor(idx), idx+1)
		if widget.settings.showProject {
			row = fmt.Sprintf(`%s%s `, row, todo.Project.Path)
		}
		row = fmt.Sprintf(`%s[mediumpurple](%s)[%s] %s`,
			row,
			todo.Author.Username,
			widget.RowColor(idx),
			widget.trimTodoBody(todo.Body),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(todo.Body))
	}

	return str
}

func (widget *Widget) markAsDone() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.todos != nil && sel < len(widget.todos) {
		todo := widget.todos[sel]
		_, err := widget.gitlabClient.Todos.MarkTodoAsDone(todo.ID)
		if err == nil {
			widget.Refresh()
		}
	}
}

func (widget *Widget) openTodo() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.todos != nil && sel < len(widget.todos) {
		todo := widget.todos[sel]
		utils.OpenFile(todo.TargetURL)
	}
}
