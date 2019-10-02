package todo

import (
	"fmt"
	"io/ioutil"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/checklist"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"github.com/wtfutil/wtf/wtf"
	"gopkg.in/yaml.v2"
)

const (
	modalHeight = 7
	modalWidth  = 80
	offscreen   = -1000
)

// A Widget represents a Todo widget
type Widget struct {
	view.KeyboardWidget
	view.TextWidget

	app      *tview.Application
	settings *Settings
	filePath string
	list     checklist.Checklist
	pages    *tview.Pages
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget: view.NewKeyboardWidget(app, pages, settings.common),
		TextWidget:     view.NewTextWidget(app, settings.common),

		app:      app,
		settings: settings,
		filePath: settings.filePath,
		list:     checklist.NewChecklist(settings.common.Sigils.Checkbox.Checked, settings.common.Sigils.Checkbox.Unchecked),
		pages:    pages,
	}

	widget.init()

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.View.SetRegions(true)
	widget.View.SetScrollable(true)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.load()
	widget.display()
}

func (widget *Widget) SetList(list checklist.Checklist) {
	widget.list = list
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

/* -------------------- Unexported Functions -------------------- */

// edit opens a modal dialog that permits editing the text of the currently-selected item
func (widget *Widget) editSelected() {
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

// Loads the todo list from Yaml file
func (widget *Widget) load() {
	confDir, _ := cfg.WtfConfigDir()
	filePath := fmt.Sprintf("%s/%s", confDir, widget.filePath)

	fileData, _ := utils.ReadFileBytes(filePath)

	yaml.Unmarshal(fileData, &widget.list)

	widget.setItemChecks()
}

func (widget *Widget) newItem() {
	form := widget.modalForm("New Todo:", "")

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

	widget.app.QueueUpdate(func() {
		widget.app.Draw()
	})
}

// persist writes the todo list to Yaml file
func (widget *Widget) persist() {
	confDir, _ := cfg.WtfConfigDir()
	filePath := fmt.Sprintf("%s/%s", confDir, widget.filePath)

	fileData, _ := yaml.Marshal(&widget.list)

	err := ioutil.WriteFile(filePath, fileData, 0644)

	if err != nil {
		panic(err)
	}
}

// setItemChecks rolls through the checklist and ensures that all checklist
// items have the correct checked/unchecked icon per the user's preferences
func (widget *Widget) setItemChecks() {
	for _, item := range widget.list.Items {
		item.CheckedIcon = widget.settings.checked
		item.UncheckedIcon = widget.settings.unchecked
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
	widget.app.QueueUpdateDraw(func() {
		frame := widget.modalFrame(form)
		widget.pages.AddPage("modal", frame, false, true)
		widget.app.SetFocus(frame)
	})
}

func (widget *Widget) modalForm(lbl, text string) *tview.Form {
	form := tview.NewForm().SetFieldBackgroundColor(wtf.ColorFor(widget.settings.common.Colors.Background))
	form.SetButtonsAlign(tview.AlignCenter).SetButtonTextColor(wtf.ColorFor(widget.settings.common.Colors.Text))

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
