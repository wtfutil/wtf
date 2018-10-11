package logger

import (
	"fmt"
	//"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

const maxBufferSize int64 = 1024

type Widget struct {
	wtf.TextWidget

	filePath string
}

func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, "Logs", "logger", true),

		filePath: logFilePath(),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func Log(msg string) {
	if logFileMissing() {
		return
	}

	f, err := os.OpenFile(logFilePath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(msg)
}

func (widget *Widget) Refresh() {
	if logFileMissing() {
		return
	}

	widget.View.SetTitle(widget.Name)

	logLines := widget.tailFile()
	widget.View.SetText(widget.contentFrom(logLines))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(logLines []string) string {
	str := ""

	for _, line := range logLines {
		chunks := strings.Split(line, " ")

		if len(chunks) >= 4 {
			str = str + fmt.Sprintf(
				"[green]%s[white] [yellow]%s[white] %s\n",
				chunks[0],
				chunks[1],
				strings.Join(chunks[3:], " "),
			)
		}
	}

	return str
}

func logFileMissing() bool {
	return logFilePath() == ""
}

func logFilePath() string {
	dir, err := wtf.Home()
	if err != nil {
		return ""
	}

	return filepath.Join(dir, ".config", "wtf", "log.txt")
}

func (widget *Widget) tailFile() []string {
	file, err := os.Open(widget.filePath)
	if err != nil {
		return []string{}
	}
	defer file.Close()

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
