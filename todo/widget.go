package todo

import (
	"fmt"
	"io/ioutil"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/cfg"
	"github.com/senorprogrammer/wtf/checklist"
	"github.com/senorprogrammer/wtf/wtf"
	"gopkg.in/yaml.v2"
)

const HelpText = `
 Keyboard commands for Todo:

   /: Show/hide this help window
   j: Select the next item in the list
   k: Select the previous item in the list
   n: Create a new list item
   o: Open the todo file in the operating system

   arrow down: Select the next item in the list
   arrow up:   Select the previous item in the list

   ctrl-d: Delete the selected item

   esc:    Unselect the todo list
   return: Edit selected item
   space:  Check the selected item on or off
`

const offscreen = -1000
const modalWidth = 80
const modalHeight = 7

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	app      *tview.Application
	filePath string
	list     checklist.Checklist
	pages    *tview.Pages
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget("Todo", "todo", true),

		app:      app,
		filePath: wtf.Config.UString("wtf.mods.todo.filename"),
		list:     checklist.NewChecklist(),
		pages:    pages,
	}

	widget.init()
	widget.HelpfulWidget.SetView(widget.View)

	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()
	widget.load()
	widget.display()

	widget.View.SetTitle(widget.ContextualTitle(widget.Name))
}

func (widget *Widget) SetList(newList checklist.Checklist) {
	widget.list = newList
}

/* -------------------- Unexported Functions -------------------- */

// edit opens a modal dialog that permits editing the text of the currently-selected item
func (widget *Widget) editItem() {
	if widget.list.SelectedItem() == nil {
		return
	}

	form := widget.modalForm("Edit:", widget.list.SelectedItem().Text)

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

func (widget *Widget) init() {
	_, err := cfg.CreateFile(widget.filePath)
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
	case "/":
		widget.ShowHelp()
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
		confDir, _ := cfg.ConfigDir()
		wtf.OpenFile(fmt.Sprintf("%s/%s", confDir, widget.filePath))
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
	confDir, _ := cfg.ConfigDir()
	filePath := fmt.Sprintf("%s/%s", confDir, widget.filePath)

	fileData, _ := wtf.ReadFileBytes(filePath)
	yaml.Unmarshal(fileData, &widget.list)
}

func (widget *Widget) newItem() {
	form := widget.modalForm("New:", "")

	saveFctn := func() {
		text := form.GetFormItem(0).(*tview.InputField).GetText()

		widget.list.Add(false, text)
		widget.persist()
		widget.pages.RemovePage("modal")
		widget.app.SetFocus(widget.View)
		widget.display()
	}

	widget.addButtons(form, saveFctn)
	widget.modalFocus(form)
}

// persist writes the todo list to Yaml file
func (widget *Widget) persist() {
	confDir, _ := cfg.ConfigDir()
	filePath := fmt.Sprintf("%s/%s", confDir, widget.filePath)

	fileData, _ := yaml.Marshal(&widget.list)

	err := ioutil.WriteFile(filePath, fileData, 0644)

	if err != nil {
		panic(err)
	}
}

/* -------------------- Modal Form -------------------- */

func (widget *Widget) addButtons(form *tview.Form, saveFctn func()) {
	widget.addSaveButton(form, saveFctn)
	widget.addCancelButton(form)
}

func (widget *Widget) addCancelButton(form *tview.Form) {
	cancelFn := func() {
		widget.pages.RemovePage("modal")
		widget.app.SetFocus(widget.View)
		widget.display()
	}

	form.AddButton("Cancel", cancelFn)
	form.SetCancelFunc(cancelFn)
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
	frame := tview.NewFrame(form).SetBorders(0, 0, 0, 0, 0, 0)
	frame.SetRect(offscreen, offscreen, modalWidth, modalHeight)
	frame.SetBorder(true)
	frame.SetBorders(1, 1, 0, 0, 1, 1)

	drawFunc := func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		w, h := screen.Size()
		frame.SetRect((w/2)-(width/2), (h/2)-(height/2), width, height)
		return x, y, width, height
	}

	frame.SetDrawFunc(drawFunc)

	return frame
}
