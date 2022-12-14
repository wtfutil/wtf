package git

import (
	"log"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

const (
	modalHeight = 7
	modalWidth  = 80
	offscreen   = -1000
)

type Widget struct {
	view.MultiSourceWidget
	view.TextWidget

	GitRepos []*GitRepo

	pages    *tview.Pages
	settings *Settings
	tviewApp *tview.Application
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "repository", "repositories"),
		TextWidget:        view.NewTextWidget(tviewApp, redrawChan, pages, settings.Common),

		tviewApp: tviewApp,
		pages:    pages,
		settings: settings,
	}

	widget.initializeKeyboardControls()

	widget.SetDisplayFunction(widget.display)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Checkout() {
	form := widget.modalForm("Branch to checkout:", "")

	checkoutFctn := func() {
		text := form.GetFormItem(0).(*tview.InputField).GetText()
		repoToCheckout := widget.GitRepos[widget.Idx]
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
	repoToPull := widget.GitRepos[widget.Idx]
	repoToPull.pull()
	widget.Refresh()

}

func (widget *Widget) Refresh() {
	repoPaths := utils.ToStrs(widget.settings.repositories)

	widget.GitRepos = widget.gitRepos(repoPaths)

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
	frame := widget.modalFrame(form)
	widget.pages.AddPage("modal", frame, false, true)
	widget.tviewApp.SetFocus(frame)
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

func (widget *Widget) currentData() *GitRepo {
	if len(widget.GitRepos) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.GitRepos) {
		return nil
	}

	return widget.GitRepos[widget.Idx]
}

func (widget *Widget) gitRepos(repoPaths []string) []*GitRepo {
	repos := []*GitRepo{}

	for _, repoPath := range repoPaths {
		if strings.HasSuffix(repoPath, string(os.PathSeparator)) {
			repos = append(repos, widget.findGitRepositories(make([]*GitRepo, 0), repoPath)...)

		} else {
			repo := NewGitRepo(
				repoPath,
				widget.settings.commitCount,
				widget.settings.commitFormat,
				widget.settings.dateFormat,
			)

			repos = append(repos, repo)
		}
	}

	return repos
}

func (widget *Widget) findGitRepositories(repositories []*GitRepo, directory string) []*GitRepo {
	directory = strings.TrimSuffix(directory, string(os.PathSeparator))

	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = directory + string(os.PathSeparator) + file.Name()

			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, string(os.PathSeparator)+".git")

				repo := NewGitRepo(
					path,
					widget.settings.commitCount,
					widget.settings.commitFormat,
					widget.settings.dateFormat,
				)

				repositories = append(repositories, repo)
				continue
			}
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}
			repositories = widget.findGitRepositories(repositories, path)
		}
	}

	return repositories
}
