package textfile

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget

	FilePath string
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" 📄 Text File ", "textfile"),
		FilePath:   Config.UString("wtf.mods.textfile.filename"),
	}

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.View.SetTitle(fmt.Sprintf(" 📄 %s ", widget.FilePath))
	widget.RefreshedAt = time.Now()

	widget.View.Clear()

	fileData, err := wtf.ReadFile(widget.FilePath)

	if err != nil {
		fmt.Fprintf(widget.View, "%s", err)
	} else {
		fmt.Fprintf(widget.View, "%s", fileData)
	}
}

/* -------------------- Unexported Functions -------------------- */
func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "o":
		wtf.OpenFile(widget.FilePath)
		return nil
	}

	return event
}
