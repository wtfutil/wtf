package mercurial

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

const (
	modalHeight = 7
	modalWidth  = 80
	offscreen   = -1000
)

// A Widget represents a Mercurial widget
type Widget struct {
	view.KeyboardWidget
	view.MultiSourceWidget
	view.TextWidget

	app      *tview.Application
	Data     []*MercurialRepo
	pages    *tview.Pages
	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:    view.NewKeyboardWidget(app, pages, settings.common),
		MultiSourceWidget: view.NewMultiSourceWidget(settings.common, "repository", "repositories"),
		TextWidget:        view.NewTextWidget(app, settings.common),

		app:      app,
		pages:    pages,
		settings: settings,
	}

	widget.SetDisplayFunction(widget.display)

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Checkout() {
	form := widget.modalForm("Branch to checkout:", "")

	checkoutFctn := func() {
		text := form.GetFormItem(0).(*tview.InputField).GetText()
		repoToCheckout := widget.Data[widget.Idx]
		repoToCheckout.checkout(text)
		widget.pages.RemovePage("modal")
		widget.app.SetFocus(widget.View)

		widget.display()

		widget.Refresh()
	}

	widget.addButtons(form, checkoutFctn)
	widget.modalFocus(form)
}

func (widget *Widget) Pull() {
	repoToPull := widget.Data[widget.Idx]
	repoToPull.pull()
	widget.Refresh()
}

func (widget *Widget) Refresh() {
	repoPaths := utils.ToStrs(widget.settings.repositories)

	widget.Data = widget.mercurialRepos(repoPaths)

	widget.display()
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addCheckoutButton(form *tview.Form, fctn func()) {
	form.AddButton("Checkout", fctn)
}

func (widget *Widget) addButtons(form *tview.Form, checkoutFctn func()) {
	widget.addCheckoutButton(form, checkoutFctn)
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

func (widget *Widget) modalFocus(form *tview.Form) {
	widget.app.QueueUpdateDraw(func() {
		frame := widget.modalFrame(form)
		widget.pages.AddPage("modal", frame, false, true)
		widget.app.SetFocus(frame)
	})
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

func (widget *Widget) currentData() *MercurialRepo {
	if len(widget.Data) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.Data) {
		return nil
	}

	return widget.Data[widget.Idx]
}

func (widget *Widget) mercurialRepos(repoPaths []string) []*MercurialRepo {
	repos := []*MercurialRepo{}

	for _, repoPath := range repoPaths {
		repo := NewMercurialRepo(repoPath, widget.settings.commitCount, widget.settings.commitFormat)
		repos = append(repos, repo)
	}

	return repos
}
