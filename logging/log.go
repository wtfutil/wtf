package logging

import (
	"fmt"
	//"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

const maxBufferSize int64 = 1024

type Widget struct {
	wtf.TextWidget

	filePath string
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Logs ", "logging", true),

		filePath: logFilePath(),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func Log(msg string) {
	if logFileMissing() {
		return
	}

	f, err := os.OpenFile(logFilePath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	widget.UpdateRefreshedAt()
	widget.View.SetTitle(fmt.Sprintf("%s", widget.Name))

	widget.View.SetText(fmt.Sprintf("%s", widget.tailFile()))
}

/* -------------------- Unexported Functions -------------------- */

func logFileMissing() bool {
	return logFilePath() == ""
}

func logFilePath() string {
	dir, err := wtf.Home()
	if err != nil {
		return ""
	}

	return filepath.Join(dir, ".wtf", "log.txt")
}

func (widget *Widget) tailFile() string {
	file, err := os.Open(widget.filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return ""
	}

	bufferSize := maxBufferSize
	if maxBufferSize > stat.Size() {
		bufferSize = stat.Size()
	}

	startPos := stat.Size() - bufferSize

	buff := make([]byte, bufferSize)
	_, err = file.ReadAt(buff, startPos)
	if err != nil {
		return ""
	}

	dataArr := strings.Split(string(buff), "\n")

	// Reverse the array of lines
	// Offset by two to account for the blank line at the end
	last := len(dataArr) - 2
	for i := 0; i < len(dataArr)/2; i++ {
		dataArr[i], dataArr[last-i] = dataArr[last-i], dataArr[i]
	}

	return fmt.Sprintf("%s\n", strings.Join(dataArr, "\n"))
}
