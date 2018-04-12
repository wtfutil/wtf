package textfile

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/homedir"
	"github.com/senorprogrammer/wtf/wtf"
)

var Config *config.Config

type Widget struct {
	wtf.TextWidget

	FilePath string
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("TextFile", "textfile"),
		FilePath:   Config.UString("wtf.mods.textfile.filepath"),
	}

	widget.View.SetWrap(true)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.View.SetTitle(fmt.Sprintf(" %s ", widget.FilePath))
	widget.RefreshedAt = time.Now()

	widget.View.Clear()

	fileData, err := widget.readFile()

	if err != nil {
		fmt.Fprintf(widget.View, "%s", err)
	} else {
		fmt.Fprintf(widget.View, "%s", fileData)
	}
}

/* -------------------- Uneported Functions -------------------- */

func (widget *Widget) readFile() (string, error) {
	absPath, _ := homedir.Expand(widget.FilePath)

	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
