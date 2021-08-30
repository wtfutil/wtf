package todo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	filePath string
	list     checklist.Checklist
	pages    *tview.Pages
	settings *Settings
	tviewApp *tview.Application
	view.ScrollableWidget
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, pages, settings.Common),

		tviewApp: tviewApp,
		settings: settings,
		filePath: settings.filePath,
		list:     checklist.NewChecklist(settings.Sigils.Checkbox.Checked, settings.Sigils.Checkbox.Unchecked),
		pages:    pages,
	}

	widget.init()

	widget.initializeKeyboardControls()

	widget.View.SetRegions(true)
	widget.View.SetScrollable(true)

	widget.SetRenderFunction(widget.display)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// SelectedItem returns the currently-selected checklist item or nil if no item is selected
func (widget *Widget) SelectedItem() *checklist.ChecklistItem {
	var selectedItem *checklist.ChecklistItem
	if widget.isItemSelected() {
		selectedItem = widget.list.Items[widget.Selected]
	}

	return selectedItem
}

// Refresh updates the data for this widget and displays it onscreen
func (widget *Widget) Refresh() {
	widget.load()
	widget.display()
}

func (widget *Widget) SetList(list checklist.Checklist) {
	widget.list = list
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) init() {
	_, err := cfg.CreateFile(widget.filePath)
	if err != nil {
		return
	}
}

// isItemSelected returns whether any item of the todo is selected or not
func (widget *Widget) isItemSelected() bool {
	return widget.Selected >= 0 && widget.Selected < len(widget.list.Items)
}

// Loads the todo list from3 Yaml file
func (widget *Widget) load() {
	confDir, _ := cfg.WtfConfigDir()
	filePath := fmt.Sprintf("%s/%s", confDir, widget.filePath)

	fileData, _ := utils.ReadFileBytes(filePath)

	err := yaml.Unmarshal(fileData, &widget.list)
	if err != nil {
		return
	}

	widget.ScrollableWidget.SetItemCount(len(widget.list.Items))
	widget.setItemChecks()
}

func (widget *Widget) newItem() {
	form := widget.modalForm("New Todo:", "")

	saveFctn := func() {
		text := widget.parseText(form.GetFormItem(0).(*tview.InputField).GetText())
		date := getTodoDate(text)


		widget.list.Add(false, date, text, widget.settings.newPos)
		widget.SetItemCount(len(widget.list.Items))
		widget.persist()
		widget.pages.RemovePage("modal")
		widget.tviewApp.SetFocus(widget.View)
		widget.display()
	}

	widget.addButtons(form, saveFctn)
	widget.modalFocus(form)

	widget.tviewApp.QueueUpdate(func() {
		widget.tviewApp.Draw()
	})
}

type PatternDuration struct {
	pattern string
	d int
	m int
	y int
}

func (widget *Widget) parseText(text string) string {
	if !widget.settings.parseDates {
		return text
	}

	now := time.Now()
	textLower := strings.ToLower(text)
	// check for "in X days/weeks/months/years" pattern
	r, _ := regexp.Compile("(?i)^in [0-9]+ (day|week|month|year)(s|)")
	match := r.FindString(text)
	if len(match) > 0 && len(text) > len(match) {
		parts := strings.Split(text, " ")
		n, _ := strconv.Atoi(parts[1])
		unit := parts[2][:1]
		target := time.Now()
		if unit == "d" {
			target = now.AddDate(0, 0, n)
		} else if unit == "w" {
			target = now.AddDate(0, 0, 7 * n)
		} else if unit == "m" {
			target = now.AddDate(0, n, 0)
		} else {
			target = now.AddDate(n, 0, 0)
		}
		return widget._textWithDate(target, text[len(match):])
	}

	// check for "today / tomorrow / next X"
	patterns := [...]PatternDuration{
		{pattern: "today",d:0,m:0,y:0},
		{pattern: "tomorrow",d:1,m:0,y:0},
		{pattern: "next week",d:7,m:0,y:0},
		{pattern: "next month",d:0,m:1,y:0},
		{pattern: "next year",d:0,m:0,y:1},
	}
	for _, pd := range patterns {
		if strings.HasPrefix(textLower, pd.pattern) && len(text) > len(pd.pattern) {
			return widget._textWithDate(now.AddDate(pd.y,pd.m,pd.d), text[len(pd.pattern):])
		}
	}

	// check for "next X" where X is name of a day (monday, etc)
	if strings.HasPrefix(textLower, "next") {
		parts := strings.Split(textLower, " ")
		if parts[0] == "next" && len(parts) > 2 {
			for i, d := range []string{"sunday", "monday","tuesday","wednesday","thursday","friday","saturday"} {
				if strings.ToLower(parts[1]) == d {
					return widget._textWithDate(now.AddDate(0,0,int(now.Weekday())+7-i), text[len(d) + 5:])
				}
			}
		}
	}

	// check for YYYY-MM-DD prefix
	if len(text) > 10 {
		date, err := time.Parse("2006-01-02", text[:10])
		if err == nil {
			return widget._textWithDate(date, text[10:])
		}
	}

	// check for MM-DD prefix
	if len(text) > 5 {
		date, err := time.Parse("2006-01-02", strconv.FormatInt(int64(now.Year()),10) + "-" + text[:5])
		if err == nil {
			return widget._textWithDate(date, text[5:])
		}
	}

	return text
}

// helper for setting dated todo text in a stable format
func (widget *Widget) _textWithDate(date time.Time, text string) string {
	return fmt.Sprintf("[%d-%02d-%02d]", date.Year(), date.Month(), date.Day()) + text
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

// updateSelected sets the text of the currently-selected item to the provided text
func (widget *Widget) updateSelected() {
	if !widget.isItemSelected() {
		return
	}

	form := widget.modalForm("Edit:", widget.SelectedItem().Text)

	saveFctn := func() {
		text := widget.parseText(form.GetFormItem(0).(*tview.InputField).GetText())

		widget.updateSelectedItem(text)
		widget.persist()
		widget.pages.RemovePage("modal")
		widget.tviewApp.SetFocus(widget.View)
		widget.display()
	}

	widget.addButtons(form, saveFctn)
	widget.modalFocus(form)

	widget.tviewApp.QueueUpdate(func() {
		widget.tviewApp.Draw()
	})
}

// updateSelectedItem update the text of the selected item.
func (widget *Widget) updateSelectedItem(text string) {
	selectedItem := widget.SelectedItem()
	if selectedItem == nil {
		return
	}

	selectedItem.Text = text
	selectedItem.Date = getTodoDate(text)
}

/* -------------------- Modal Form -------------------- */

func (widget *Widget) addButtons(form *tview.Form, saveFctn func()) {
	widget.addSaveButton(form, saveFctn)
	widget.addCancelButton(form)
}

func (widget *Widget) addCancelButton(form *tview.Form) {
	cancelFn := func() {
		widget.pages.RemovePage("modal")
		widget.tviewApp.SetFocus(widget.View)
		widget.display()
	}

	form.AddButton("Cancel", cancelFn)
	form.SetCancelFunc(cancelFn)
}

func (widget *Widget) addSaveButton(form *tview.Form, fctn func()) {
	form.AddButton("Save", fctn)
}

func (widget *Widget) modalFocus(form *tview.Form) {
	widget.tviewApp.QueueUpdateDraw(func() {
		frame := widget.modalFrame(form)
		widget.pages.AddPage("modal", frame, false, true)
		widget.tviewApp.SetFocus(frame)
	})
}

func (widget *Widget) modalForm(lbl, text string) *tview.Form {
	form := tview.NewForm().SetFieldBackgroundColor(wtf.ColorFor(widget.settings.Colors.Background))
	form.SetButtonsAlign(tview.AlignCenter).SetButtonTextColor(wtf.ColorFor(widget.settings.Colors.Text))

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
