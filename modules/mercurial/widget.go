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
	view.MultiSourceWidget
	view.TextWidget

	Data     []*MercurialRepo
	pages    *tview.Pages
	settings *Settings
	tviewApp *tview.Application
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "repository", "repositories"),
		TextWidget:        view.NewTextWidget(tviewApp, pages, settings.Common),

		tviewApp: tviewApp,
		pages:    pages,
		settings: settings,
	}

	widget.SetDisplayFunction(widget.display)

	widget.initializeKeyboardControls()

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
		widget.tviewApp.SetFocus(widget.View)

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
		widget.tviewApp.SetFocus(widget.View)
		widget.display()
	}

	form.AddButton("Cancel", cancelFn)
	form.SetCancelFunc(cancelFn)
}

func (widget *Widget) modalFocus(form *tview.Form) {
	widget.tviewApp.QueueUpdateDraw(func() {
		frame := widget.modalFrame(form)
		widget.pages.AddPage("modal", frame, false, true)
		widget.tviewApp.SetFocus(frame)
	})
}

func (widget *Widget) modalForm(lbl, text string) *tview.Form {
	form := tview.NewForm()
	form.SetButtonsAlign(tview.AlignCenter)
	form.SetButtonTextColor(tview.Styles.PrimaryTextColor)

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
