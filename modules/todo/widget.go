package todo

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
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
	filePath      string
	list          checklist.Checklist
	pages         *tview.Pages
	settings      *Settings
	showTagPrefix string
	showFilter    string
	tviewApp      *tview.Application
	Error         string

	view.ScrollableWidget

	// redrawChan chan bool
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),

		tviewApp:      tviewApp,
		settings:      settings,
		filePath:      settings.filePath,
		showTagPrefix: "",
		list:          checklist.NewChecklist(settings.Sigils.Checkbox.Checked, settings.Sigils.Checkbox.Unchecked),
		pages:         pages,

		// redrawChan: redrawChan,
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
	widget.Error = ""
	err := widget.load()
	if err != nil {
		widget.Error = err.Error()
	}
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
func (widget *Widget) load() error {
	confDir, _ := cfg.WtfConfigDir()
	filePath := fmt.Sprintf("%s/%s", confDir, widget.filePath)

	fileData, err := utils.ReadFileBytes(filePath)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileData, &widget.list)
	if err != nil {
		return err
	}

	// do initial sort based on dates to make sure everything is correct
	if widget.settings.parseDates {
		i := 0
		for i < widget.list.Len() {
			for {
				newIndex := widget.placeItemBasedOnDate(i)
				if newIndex == i {
					break
				}
			}
			i += 1
		}
	}

	widget.ScrollableWidget.SetItemCount(len(widget.list.Items))
	widget.setItemChecks()
	return nil
}

func (widget *Widget) newItem() {
	widget.processFormInput("New Todo:", "", func(t string) {
		text, date, tags := widget.getTextComponents(t)

		widget.list.Add(false, date, tags, text, widget.settings.newPos)
		widget.SetItemCount(len(widget.list.Items))
		if widget.settings.parseDates {
			if widget.settings.newPos == "first" {
				widget.placeItemBasedOnDate(0)
			} else {
				widget.placeItemBasedOnDate(widget.list.Len() - 1)
			}
		}
		widget.persist()
	})
}

func (widget *Widget) getTextComponents(text string) (string, *time.Time, []string) {
	var date *time.Time = nil
	if widget.settings.parseDates {
		text, date = widget.getTextAndDate(text)
	}

	tags := make([]string, 0)
	if widget.settings.parseTags {
		text, tags = getTodoTags(text)
	}

	text = strings.TrimSpace(text)
	return text, date, tags
}

func getTodoTags(text string) (string, []string) {
	tags := make([]string, 0)
	r, _ := regexp.Compile(`(?i)(^|\s)#[a-z0-9]+`)
	matches := r.FindAllString(text, -1)

	for _, tag := range matches {
		tag = strings.TrimSpace(tag)
		suffix := " "
		if strings.HasSuffix(text, tag) {
			suffix = ""
		}
		text = strings.Replace(text, tag+suffix, "", 1)
		tags = append(tags, tag[1:])
	}

	return text, tags
}

type PatternDuration struct {
	pattern string
	d       int
	m       int
	y       int
}

func (widget *Widget) getTextAndDate(text string) (string, *time.Time) {
	now := time.Now()
	textLower := strings.ToLower(text)
	// check for "in X days/weeks/months/years" pattern
	r, _ := regexp.Compile("(?i)^in [0-9]+ (day|week|month|year)(s|)")
	match := r.FindString(text)
	if len(match) > 0 && len(text) > len(match) {
		parts := strings.Split(text, " ")
		n, _ := strconv.Atoi(parts[1])
		unit := parts[2][:1]
		var target time.Time
		if unit == "d" {
			target = now.AddDate(0, 0, n)
		} else if unit == "w" {
			target = now.AddDate(0, 0, 7*n)
		} else if unit == "m" {
			target = now.AddDate(0, n, 0)
		} else {
			target = now.AddDate(n, 0, 0)
		}
		return text[len(match):], &target
	}

	// check for "today / tomorrow / next X"
	patterns := [...]PatternDuration{
		{pattern: "today", d: 0, m: 0, y: 0},
		{pattern: "tomorrow", d: 1, m: 0, y: 0},
		{pattern: "next week", d: 7, m: 0, y: 0},
		{pattern: "next month", d: 0, m: 1, y: 0},
		{pattern: "next year", d: 0, m: 0, y: 1},
	}
	for _, pd := range patterns {
		if strings.HasPrefix(textLower, pd.pattern) && len(text) > len(pd.pattern) {
			date := now.AddDate(pd.y, pd.m, pd.d)
			return text[len(pd.pattern):], &date
		}
	}

	// check for "next X" where X is name of a day (monday, etc)
	if strings.HasPrefix(textLower, "next") {
		parts := strings.Split(textLower, " ")
		if parts[0] == "next" && len(parts) > 2 {
			for i, d := range []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"} {
				if strings.ToLower(parts[1]) == d {
					date := now.AddDate(0, 0, int(now.Weekday())+7-i)
					return text[len(d)+5:], &date
				}
			}
		}
	}

	// check for YYYY-MM-DD prefix
	if len(text) > 10 {
		date, err := time.Parse("2006-01-02", text[:10])
		if err == nil {
			return text[10:], &date
		}
	}

	// check for MM-DD prefix
	if len(text) > 5 {
		date, err := time.Parse("2006-01-02", strconv.FormatInt(int64(now.Year()), 10)+"-"+text[:5])
		if err == nil {
			return text[5:], &date
		}
	}

	return text, nil
}

// persist writes the todo list to Yaml file
func (widget *Widget) persist() {
	confDir, _ := cfg.WtfConfigDir()
	filePath := fmt.Sprintf("%s/%s", confDir, widget.filePath)

	fileData, _ := yaml.Marshal(&widget.list)

	err := os.WriteFile(filePath, fileData, 0644)

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

	widget.processFormInput("Edit:", widget.SelectedItem().EditText(), func(t string) {
		text, date, tags := widget.getTextComponents(t)

		widget.updateSelectedItem(text, date, tags)
		if widget.settings.parseDates {
			widget.Selected = widget.placeItemBasedOnDate(widget.Selected)
		}
		widget.persist()
	})
}

// processFormInput is a helper function that creates a form and calls onSave on the received input
func (widget *Widget) processFormInput(prompt string, initValue string, onSave func(string)) {
	form := widget.modalForm(prompt, initValue)

	saveFctn := func() {
		onSave(form.GetFormItem(0).(*tview.InputField).GetText())

		widget.pages.RemovePage("modal")
		widget.tviewApp.SetFocus(widget.View)
		widget.display()
	}

	widget.addButtons(form, saveFctn)
	widget.modalFocus(form)

	// Tell the app to force redraw the screen
	widget.Base.RedrawChan <- true
}

// updateSelectedItem update the text of the selected item.
func (widget *Widget) updateSelectedItem(text string, date *time.Time, tags []string) {
	selectedItem := widget.SelectedItem()
	if selectedItem == nil {
		return
	}

	selectedItem.Text = text
	selectedItem.Date = date
	selectedItem.Tags = tags
}

func (widget *Widget) placeItemBasedOnDate(index int) int {
	// potentially move todo up
	for index > 0 && widget.todoDateIsEarlier(index, index-1) {
		widget.list.Swap(index, index-1)
		index -= 1
	}
	// potentially move todo down
	for index < widget.list.Len()-1 && widget.todoDateIsEarlier(index+1, index) {
		widget.list.Swap(index, index+1)
		index += 1
	}
	return index
}

func (widget *Widget) todoDateIsEarlier(i, j int) bool {
	if widget.list.Items[i].Date == nil && widget.list.Items[j].Date == nil {
		return false
	}
	defaultVal := getNowDate().AddDate(0, 0, widget.settings.undatedAsDays)
	if widget.list.Items[i].Date == nil {
		return defaultVal.Before(*widget.list.Items[j].Date)
	} else if widget.list.Items[j].Date == nil {
		return widget.list.Items[i].Date.Before(defaultVal)
	} else {
		return widget.list.Items[i].Date.Before(*widget.list.Items[j].Date)
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
	frame := widget.modalFrame(form)
	widget.pages.AddPage("modal", frame, false, true)
	widget.tviewApp.SetFocus(frame)

	// Tell the app to force redraw the screen
	widget.Base.RedrawChan <- true
}

func (widget *Widget) modalForm(lbl, text string) *tview.Form {
	form := tview.NewForm()
	form.SetFieldBackgroundColor(wtf.ColorFor(widget.settings.Colors.Background))
	form.SetButtonsAlign(tview.AlignCenter)
	form.SetButtonTextColor(wtf.ColorFor(widget.settings.Colors.Text))

	form.AddInputField(lbl, text, 60, nil, nil)

	return form
}

func (widget *Widget) modalFrame(form *tview.Form) *tview.Frame {
	frame := tview.NewFrame(form)
	frame.SetBorders(0, 0, 0, 0, 0, 0)
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
