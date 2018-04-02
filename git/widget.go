package git

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.BaseWidget
	View *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		BaseWidget: wtf.BaseWidget{
			Name:        "Git",
			RefreshedAt: time.Now(),
			RefreshInt:  10,
		},
	}

	widget.addView()
	go wtf.Refresh(&widget)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data := Fetch()

	str := fmt.Sprintf("[green]%s[white] [dodgerblue]%s[white]\n", data["repo"][0], data["branch"][0])

	widget.View.SetTitle(fmt.Sprintf(" 🤞 %s ", str))
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", widget.contentFrom(data))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

func (widget *Widget) contentFrom(data map[string][]string) string {
	str := "\n"
	str = str + widget.formatChanges(data["changes"])
	str = str + "\n"
	str = str + widget.formatCommits(data["commits"])

	return str
}

func (widget *Widget) formatChanges(data []string) string {
	str := ""
	str = str + " [red]Changed Files[white]\n"

	for _, line := range data {
		str = str + widget.formatChange(line)
	}

	return str
}

func (widget *Widget) formatChange(line string) string {
	if len(line) == 0 {
		return ""
	}

	line = strings.TrimSpace(line)
	firstChar, _ := utf8.DecodeRuneInString(line)

	// Revisit this and kill the ugly duplication
	switch firstChar {
	case 'A':
		line = strings.Replace(line, "A", "[green]A[white]", 1)
	case 'D':
		line = strings.Replace(line, "D", "[red]D[white]", 1)
	case 'M':
		line = strings.Replace(line, "M", "[yellow]M[white]", 1)
	case 'R':
		line = strings.Replace(line, "R", "[purple]R[white]", 1)
	default:
		line = line
	}

	return fmt.Sprintf(" %s\n", strings.Replace(line, "\"", "", -1))
}

func (widget *Widget) formatCommits(data []string) string {
	str := ""
	str = str + " [red]Recent Commits[white]\n"

	for _, line := range data {
		str = str + widget.formatCommit(line)
	}

	return str
}

func (widget *Widget) formatCommit(line string) string {
	return fmt.Sprintf(" %s\n", strings.Replace(line, "\"", "", -1))
}
