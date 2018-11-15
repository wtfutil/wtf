package git

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

const HelpText = `
  Keyboard commands for Git:

    /: Show/hide this help window
    c: Checkout to branch
    h: Previous git repository
    l: Next git repository
    p: Pull current git repository

    arrow left:  Previous git repository
    arrow right: Next git repository
`

const offscreen = -1000
const modalWidth = 80
const modalHeight = 7

type Widget struct {
	wtf.HelpfulWidget
	wtf.MultiSourceWidget
	wtf.TextWidget

	app      *tview.Application
	GitRepos []*GitRepo
	pages    *tview.Pages
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget:     wtf.NewHelpfulWidget(app, pages, HelpText),
		MultiSourceWidget: wtf.NewMultiSourceWidget("git", "repository", "repositories"),
		TextWidget:        wtf.NewTextWidget(app, "Git", "git", true),

		app:   app,
		pages: pages,
	}

	widget.LoadSources()
	widget.SetDisplayFunction(widget.display)

	widget.HelpfulWidget.SetView(widget.View)
	widget.View.SetInputCapture(widget.keyboardIntercept)

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
		widget.app.SetFocus(widget.View)
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
	repoPaths := wtf.ToStrs(wtf.Config.UList("wtf.mods.git.repositories"))

	widget.GitRepos = widget.gitRepos(repoPaths)
	sort.Slice(widget.GitRepos, func(i, j int) bool {
		return widget.GitRepos[i].Path < widget.GitRepos[j].Path
	})
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
		widget.app.SetFocus(widget.View)
		widget.display()
	}

	form.AddButton("Cancel", cancelFn)
	form.SetCancelFunc(cancelFn)
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
		if strings.HasSuffix(repoPath, "/") {
			repos = append(repos, findGitRepositories(make([]*GitRepo, 0), repoPath)...)

		} else {
			repo := NewGitRepo(repoPath)
			repos = append(repos, repo)
		}
	}

	return repos
}

func findGitRepositories(repositories []*GitRepo, directory string) []*GitRepo {
	directory = strings.TrimSuffix(directory, "/")

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = directory + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				repo := NewGitRepo(path)
				repositories = append(repositories, repo)
				continue
			}
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}
			repositories = findGitRepositories(repositories, path)
		}
	}

	return repositories
}

func (widget *Widget) Next() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.GitRepos) {
		widget.Idx = 0
	}

	if widget.DisplayFunction != nil {
		widget.DisplayFunction()
	}
}

func (widget *Widget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.GitRepos) - 1
	}

	if widget.DisplayFunction != nil {
		widget.DisplayFunction()
	}
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
		return nil
	case "h":
		widget.Prev()
		return nil
	case "l":
		widget.Next()
		return nil
	case "p":
		widget.Pull()
		return nil
	case "c":
		widget.Checkout()
		return nil
	}

	switch event.Key() {
	case tcell.KeyLeft:
		widget.Prev()
		return nil
	case tcell.KeyRight:
		widget.Next()
		return nil
	default:
		return event
	}
}
