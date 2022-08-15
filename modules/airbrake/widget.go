package airbrake

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

const module = "Airbrake"

var emojis = map[string]string{
	"bug":             "ð",
	"bell with slash": "ð",
}

type ShowType int

const (
	SHOW_TITLE ShowType = iota
	SHOW_COMPARE
)

type Widget struct {
	view.ScrollableWidget

	settings *Settings
	app      *tview.Application
	pages    *tview.Pages

	groups  []Group
	project *Project

	showType ShowType
	err      error
}

type GroupJSON struct {
	Groups []Group `json:"groups"`
}

type Group struct {
	ID           string `json:"id"`
	ProjectID    int64  `json:"projectId"`
	Errors       []Error
	NoticeCount  int64  `json:"noticeCount"`
	CreatedAt    string `json:"createdAt"`
	LastNoticeAt string `json:"lastNoticeAt"`
	Context      GroupContext
	Muted        bool  `json:"muted"`
	CommentCount int64 `json:"commentCount"`
}

type GroupContext struct {
	Environment string `json:"environment"`
	Severity    string `json:"severity"`
}

func (g *Group) Link() string {
	return fmt.Sprintf("https://airbrake.io/projects/%d/groups/%s", g.ProjectID, g.ID)
}

func (g *Group) Title() string {
	return fmt.Sprintf("%s: %s", g.Type(), g.Message())
}

func (g *Group) Type() string {
	return g.Errors[0].Type
}

func (g *Group) Message() string {
	err := g.Errors[0]
	return strings.ReplaceAll(err.Message, "\n", ". ")
}

func (g *Group) File() string {
	s := fmt.Sprintf("%s:%d", g.Errors[0].Backtrace[0].File,
		g.Errors[0].Backtrace[0].Line)
	return reverseString(utils.Truncate(reverseString(s), 51, true))
}

type Error struct {
	Type      string       `json:"type"`
	Message   string       `json:"message"`
	Backtrace []StackFrame `json:"backtrace"`
}

type StackFrame struct {
	File     string `json:"file"`
	Function string `json:"function"`
	Line     int64  `json:"line"`
}

type ProjectJSON struct {
	Project Project `json:"project"`
}

type Project struct {
	Name string `json:"name"`
}

func rotateShowType(showtype ShowType) ShowType {
	returnValue := SHOW_TITLE
	switch showtype {
	case SHOW_TITLE:
		returnValue = SHOW_COMPARE
	case SHOW_COMPARE:
		returnValue = SHOW_TITLE
	}
	return returnValue
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),

		app:      tviewApp,
		settings: settings,
		pages:    pages,
		showType: SHOW_TITLE,
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

	groups, err := groups(
		widget.settings.projectID,
		widget.settings.authToken)
	if err != nil {
		widget.err = err
		widget.groups = nil
		widget.SetItemCount(0)
	} else {
		widget.err = nil
		widget.groups = groups
		widget.SetItemCount(len(groups))
	}

	project, err := project(
		widget.settings.projectID,
		widget.settings.authToken)
	if err != nil {
		widget.err = err
		widget.project = nil
	} else {
		widget.err = nil
		widget.project = project
	}

	widget.Render()
}

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	if widget.err != nil {
		return module, widget.err.Error(), true
	}

	project := widget.project
	if project != nil && project.Name == "" {
		return module, "No project found", true
	}

	title := fmt.Sprintf("%s %s - %s's recent errors", emojis["bug"],
		module, project.Name)

	result := widget.groups
	if result == nil || len(widget.groups) == 0 {
		return title, "All your errors are resolved!", false
	}

	var str string
	for idx, g := range widget.groups {
		rowColor := widget.RowColor(idx)
		var row string

		if widget.showType == SHOW_TITLE {
			var buf bytes.Buffer
			if g.Muted {
				buf.WriteString(emojis["bell with slash"])
			} else {
				buf.WriteString("  ")
			}
			buf.WriteString(" " + g.Title())
			row = fmt.Sprintf("[%s]%2d. %s[white]", rowColor, idx+1, buf.String())
		} else {
			row = fmt.Sprintf(
				"[%s]%2d. %-31s %-11s  %-10s  count: %-9d  comments: %-2d[white]",
				rowColor, idx+1, utils.Truncate(g.Type(), 30, true),
				g.Context.Environment, g.Context.Severity,
				g.NoticeCount, g.CommentCount)
		}

		str += utils.HighlightableHelper(widget.View, row, idx, len(g.Type()))
	}

	return title, str, false
}

func (widget *Widget) openGroup() {
	sel := widget.GetSelected()

	if sel >= 0 && widget.groups != nil && sel < len(widget.groups) {
		group := widget.groups[sel]
		utils.OpenFile(group.Link())
	}
}

func (widget *Widget) viewGroup() {
	sel := widget.GetSelected()

	if sel >= 0 && widget.groups != nil && sel < len(widget.groups) {
		group := widget.groups[sel]

		closeFunc := func() {
			widget.pages.RemovePage("group info")
			widget.app.SetFocus(widget.View)
		}

		table := newGroupInfoTable(&group).render()
		table += utils.CenterText("Esc to close", 80)

		modal := view.NewBillboardModal(table, closeFunc)
		modal.SetTitle(fmt.Sprintf(" %s ", group.Title()))

		widget.pages.AddPage("group info", modal, false, true)
		widget.app.SetFocus(modal)

		widget.app.QueueUpdateDraw(func() {
			widget.app.Draw()
		})
	}
}

func (widget *Widget) resolveGroup() {
	sel := widget.GetSelected()

	if sel >= 0 && widget.groups != nil && sel < len(widget.groups) {
		group := widget.groups[sel]

		closeFunc := func() {
			widget.pages.RemovePage("resolve")
			widget.app.SetFocus(widget.View)
		}

		var tbl *resultTable
		err := resolveGroup(group.ProjectID, group.ID, widget.settings.authToken)
		if err == nil {
			tbl = newResultTable("Success", "Error Resolved")
			widget.Refresh()
		} else {
			tbl = newResultTable("Error", err.Error())
		}

		modal := view.NewBillboardModal(tbl.render(), closeFunc)
		modal.SetTitle(fmt.Sprintf(" %s ", group.Title()))

		widget.pages.AddPage("resolve", modal, false, true)
		widget.app.SetFocus(modal)

		widget.app.QueueUpdateDraw(func() {
			widget.app.Draw()
		})
	}
}

func (widget *Widget) muteGroup() {
	if widget.showType != SHOW_TITLE {
		return
	}

	sel := widget.GetSelected()

	if sel >= 0 && widget.groups != nil && sel < len(widget.groups) {
		group := widget.groups[sel]
		if !group.Muted {
			widget.err = muteGroup(group.ProjectID, group.ID, widget.settings.authToken)
			widget.Refresh()
		}
	}
}

func (widget *Widget) unmuteGroup() {
	if widget.showType != SHOW_TITLE {
		return
	}

	sel := widget.GetSelected()

	if sel >= 0 && widget.groups != nil && sel < len(widget.groups) {
		group := widget.groups[sel]
		if group.Muted {
			widget.err = unmuteGroup(group.ProjectID, group.ID, widget.settings.authToken)
			widget.Refresh()
		}
	}
}

func (widget *Widget) toggleDisplayText() {
	widget.showType = rotateShowType(widget.showType)
	widget.Render()
}
