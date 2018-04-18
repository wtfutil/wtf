package git

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/senorprogrammer/wtf/wtf"
)

func (widget *Widget) display() {
	widget.View.Clear()

	repoData := widget.currentData()
	if repoData == nil {
		fmt.Fprintf(widget.View, "%s", " Git repo data is unavailable (1)")
		return
	}

	title := fmt.Sprintf("[green]%s[white]\n", repoData.Repository)
	widget.View.SetTitle(fmt.Sprintf(" Git: %s ", title))

	str := widget.tickMarks(widget.Data) + "\n"
	str = str + " [red]Branch[white]\n"
	str = str + fmt.Sprintf(" %s", repoData.Branch)
	str = str + "\n"
	str = str + widget.formatChanges(repoData.ChangedFiles)
	str = str + "\n"
	str = str + widget.formatCommits(repoData.Commits)

	fmt.Fprintf(widget.View, "%s", str)
}

func (widget *Widget) formatChanges(data []string) string {
	str := ""
	str = str + " [red]Changed Files[white]\n"

	if len(data) == 1 {
		str = str + " [grey]none[white]\n"
	} else {
		for _, line := range data {
			str = str + widget.formatChange(line)
		}
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

func (widget *Widget) tickMarks(data []*GitRepo) string {
	str := ""

	if len(data) > 1 {
		marks := strings.Repeat("*", len(data))
		marks = marks[:widget.Idx] + "_" + marks[widget.Idx+1:]

		str = "[lightblue]" + fmt.Sprintf(wtf.RightAlignFormat(widget.View), marks) + "[white]"
	}

	return str
}
