package todo

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
	"gopkg.in/yaml.v2"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget

	FilePath string
	list     *List
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" 📝 Todo ", "todo"),
		FilePath:   Config.UString("wtf.mods.todo.filename"),

		list: &List{selected: -1},
	}

	widget.init()
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	confDir, _ := wtf.ConfigDir()

	fileData, _ := wtf.ReadYamlFile(fmt.Sprintf("%s/%s", confDir, widget.FilePath))
	yaml.Unmarshal(fileData, &widget.list)

	widget.display()
	widget.RefreshedAt = time.Now()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) init() {
	_, err := wtf.CreateFile(widget.FilePath)
	if err != nil {
		panic(err)
	}
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case " ":
		// Check/uncheck selected item
		widget.list.Toggle()
		widget.display()
		return nil
	case "e":
		// Edit selected item
		return nil
	case "n":
		// Add a new item
		return nil
	}

	switch event.Key() {
	case tcell.KeyCtrlD:
		// Delete selected item
		return nil
	case tcell.KeyDown:
		widget.list.Next()
		widget.display()
		return nil
	//case tcell.KeySpac:
	//// Check/uncheck an item
	//return nil
	case tcell.KeyEsc:
		// Unselect the current row and pass the key on through to unselect the widget
		widget.list.Unselect()
		widget.display()
		return event
	case tcell.KeyUp:
		// Select next item up
		widget.list.Prev()
		widget.display()
		return nil
	default:
		return event
	}
}
