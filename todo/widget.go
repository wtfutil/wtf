package todo

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	"gopkg.in/yaml.v2"
)

// Config is a pointer to the global config object
var Config *config.Config

const helpText = `
	Todo

	h - Displays the help text
	o - Opens the todo file in the operating system

	space - checks an item on or off
`

type Widget struct {
	wtf.TextWidget

	app      *tview.Application
	pages    *tview.Pages
	filePath string
	list     *List
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" 📝 Todo ", "todo", true),

		app:      app,
		pages:    pages,
		filePath: Config.UString("wtf.mods.todo.filename"),
		list:     &List{selected: -1},
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

// edit opens a modal dialog that permits editing the text of the currently-selected item
func (widget *Widget) editItem() {
	if widget.list.Selected() == nil {
		return
	}

	form := widget.modalForm("Edit:", widget.list.Selected().Text)

	saveFctn := func() {
		text := form.GetFormItem(0).(*tview.InputField).GetText()

		widget.list.Update(text)
		widget.persist()
		widget.pages.RemovePage("modal")
		widget.app.SetFocus(widget.View)
		widget.display()
	}

	widget.addButtons(form, saveFctn)
	widget.modalFocus(form)
}

func (widget *Widget) help() {
	cancelFn := func() {
		widget.pages.RemovePage("billboard")
		widget.app.SetFocus(widget.View)
		widget.display()
	}

	billboard := wtf.NewBillboardModal(helpText, cancelFn)

	widget.pages.AddPage("billboard", billboard, false, true)
	widget.app.SetFocus(billboard)
}

func (widget *Widget) init() {
	_, err := wtf.CreateFile(widget.filePath)
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
	case "h":
		// Show help menu
		widget.help()
		return nil
	case "j":
		// Select the next item down
		widget.list.Next()
		widget.display()
		return nil
	case "k":
		// Select the next item up
		widget.list.Prev()
		widget.display()
		return nil
	case "n":
		// Add a new item
		widget.newItem()
		return nil
	case "o":
		// Open the file
		wtf.OpenFile(widget.filePath)
		return nil
	}

	switch event.Key() {
	case tcell.KeyCtrlD:
		// Delete the selected item
		widget.list.Delete()
		widget.persist()
		widget.display()
		return nil
	case tcell.KeyCtrlJ:
		// Move selected item down in the list
		widget.list.Demote()
		widget.persist()
		widget.display()
		return nil
	case tcell.KeyCtrlK:
		// Move selected item up in the list
		widget.list.Promote()
		widget.persist()
		widget.display()
		return nil
	case tcell.KeyDown:
		// Select the next item down
		widget.list.Next()
		widget.display()
		return nil
	case tcell.KeyEnter:
		widget.editItem()
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
	filePath := fmt.Sprintf("%s/%s", confDir, widget.filePath)

	fileData, _ := wtf.ReadFileBytes(filePath)
	yaml.Unmarshal(fileData, &widget.list)
}

// persist writes the todo list to Yaml file
func (widget *Widget) persist() {
	confDir, _ := wtf.ConfigDir()
	filePath := fmt.Sprintf("%s/%s", confDir, widget.filePath)

	fileData, _ := yaml.Marshal(&widget.list)

	err := ioutil.WriteFile(filePath, fileData, 0644)

	if err != nil {
		panic(err)
	}
}

func (widget *Widget) newItem() {
	form := widget.modalForm("New:", "")

	saveFctn := func() {
		text := form.GetFormItem(0).(*tview.InputField).GetText()

		widget.list.Add(text)
		widget.persist()
		widget.pages.RemovePage("modal")
		widget.app.SetFocus(widget.View)
		widget.display()
	}

	widget.addButtons(form, saveFctn)
	widget.modalFocus(form)
}

/* -------------------- Modal Form -------------------- */

func (widget *Widget) addButtons(form *tview.Form, saveFctn func()) {
	widget.addSaveButton(form, saveFctn)
	widget.addCancelButton(form)
}

func (widget *Widget) addCancelButton(form *tview.Form) {
	form.AddButton("Cancel", func() {
		widget.pages.RemovePage("modal")
		widget.app.SetFocus(widget.View)
		widget.display()
	})
}

func (widget *Widget) addSaveButton(form *tview.Form, fctn func()) {
	form.AddButton("Save", fctn)
}

func (widget *Widget) modalFocus(form *tview.Form) {
	frame := widget.modalFrame(form)
	widget.pages.AddPage("modal", frame, false, true)
	widget.app.SetFocus(frame)
}

func (widget *Widget) modalForm(lbl, text string) *tview.Form {
	form := tview.NewForm().
		SetButtonsAlign(tview.AlignCenter).
		SetButtonTextColor(tview.Styles.PrimaryTextColor)

	form.AddInputField(lbl, text, 60, nil, nil)

	return form
}

func (widget *Widget) modalFrame(form *tview.Form) *tview.Frame {
	_, _, w, h := widget.View.GetInnerRect()

	frame := tview.NewFrame(form).SetBorders(0, 0, 0, 0, 0, 0)
	frame.SetBorder(true)
	frame.SetRect(w+20, h+2, 80, 7)

	return frame
}
