package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rivo/tview"
	log "github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/view"
)

const (
	maxBufferSize int64 = 1024
)

type Widget struct {
	view.TextWidget

	filePath string
	settings *Settings
}

func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		filePath: log.LogFilePath(),
		settings: settings,
	}

	return &widget
}

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	if log.LogFileMissing() {
		return widget.CommonSettings().Title, "File missing", false
	}

	logLines := widget.tailFile()
	str := ""

	for _, line := range logLines {
		chunks := strings.Split(line, " ")

		if len(chunks) >= 4 {
			str += fmt.Sprintf(
				"[green]%s[white] [yellow]%s[white] %s\n",
				chunks[0],
				chunks[1],
				strings.Join(chunks[3:], " "),
			)
		}
	}

	return widget.CommonSettings().Title, str, false
}

func (widget *Widget) tailFile() []string {
	file, err := os.Open(widget.filePath)
	if err != nil {
		return []string{}
	}
	defer func() { _ = file.Close() }()

	stat, err := file.Stat()
	if err != nil {
		return []string{}
	}

	bufferSize := maxBufferSize
	if maxBufferSize > stat.Size() {
		bufferSize = stat.Size()
	}

	startPos := stat.Size() - bufferSize

	buff := make([]byte, bufferSize)
	_, err = file.ReadAt(buff, startPos)
	if err != nil {
		return []string{}
	}

	logLines := strings.Split(string(buff), "\n")

	// Reverse the array of lines
	// Offset by two to account for the blank line at the end
	last := len(logLines) - 2
	for i := 0; i < len(logLines)/2; i++ {
		logLines[i], logLines[last-i] = logLines[last-i], logLines[i]
	}

	return logLines
}
