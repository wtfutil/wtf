package todo

import (
	"fmt"
	"io/ioutil"
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

	widget.load()
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
		widget.persist()
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
	case tcell.KeyCtrlA:
		// Move the selected item up
		widget.list.Promote()
		widget.persist()
		widget.display()
		return nil
	case tcell.KeyCtrlD:
		// Delete the selected item
		widget.list.Delete()
		widget.persist()
		widget.display()
		return nil
	case tcell.KeyCtrlZ:
		// Move the selected item down
		widget.list.Demote()
		widget.persist()
		widget.display()
		return nil
	case tcell.KeyDown:
		// Select the next item down
		widget.list.Next()
		widget.display()
		return nil
	case tcell.KeyEsc:
		// Unselect the current row
		widget.list.Unselect()
		widget.display()
		return event
	case tcell.KeyUp:
		// Select the next item up
		widget.list.Prev()
		widget.display()
		return nil
	default:
		// Pass it along
		return event
	}
}

// Loads the todo list from Yaml file
func (widget *Widget) load() {
	confDir, _ := wtf.ConfigDir()
	filePath := fmt.Sprintf("%s/%s", confDir, widget.FilePath)

	fileData, _ := wtf.ReadFileBytes(filePath)
	yaml.Unmarshal(fileData, &widget.list)
}

// persist writes the todo list to Yaml file
func (widget *Widget) persist() {
	confDir, _ := wtf.ConfigDir()
	filePath := fmt.Sprintf("%s/%s", confDir, widget.FilePath)

	fileData, _ := yaml.Marshal(&widget.list)

	err := ioutil.WriteFile(filePath, fileData, 0644)

	if err != nil {
		panic(err)
	}
}
