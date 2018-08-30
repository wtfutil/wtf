package textfile

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

const HelpText = `
  Keyboard commands for Textfile:

    /: Show/hide this help window
    h: Previous text file
    l: Next text file
    o: Open the text file in the operating system

    arrow left:  Previous text file
    arrow right: Next text file
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	filePaths []string
	idx       int
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget("TextFile", "textfile", true),

		idx: 0,
	}

	widget.loadFilePaths()

	widget.HelpfulWidget.SetView(widget.View)

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Next() {
	widget.idx = widget.idx + 1
	if widget.idx == len(widget.filePaths) {
		widget.idx = 0
	}

	widget.display()
}

func (widget *Widget) Prev() {
	widget.idx = widget.idx - 1
	if widget.idx < 0 {
		widget.idx = len(widget.filePaths) - 1
	}

	widget.display()
}

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) currentFilePath() string {
	return widget.filePaths[widget.idx]
}

func (widget *Widget) display() {
	widget.View.SetTitle(widget.ContextualTitle(widget.fileName()))

	text := wtf.SigilStr(len(widget.filePaths), widget.idx, widget.View) + "\n"

	if wtf.Config.UBool("wtf.mods.textfile.format", false) {
		text = text + widget.formattedText()
	} else {
		text = text + widget.plainText()
	}

	widget.View.SetText(text)
}

func (widget *Widget) fileName() string {
	return filepath.Base(widget.currentFilePath())
}

func (widget *Widget) formattedText() string {
	filePath, _ := wtf.ExpandHomeDir(widget.currentFilePath())

	file, err := os.Open(filePath)
	if err != nil {
		return err.Error()
	}

	lexer := lexers.Match(filePath)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	style := styles.Get(wtf.Config.UString("wtf.mods.textfile.formatStyle", "vim"))
	if style == nil {
		style = styles.Fallback
	}
	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	contents, _ := ioutil.ReadAll(file)
	iterator, _ := lexer.Tokenise(nil, string(contents))

	var buf bytes.Buffer
	formatter.Format(&buf, style, iterator)

	return tview.TranslateANSI(buf.String())
}

// loadFilePaths parses file paths from the config and stores them in an array
// It is backwards-compatible, supporting both the original, singular filePath and
// the current plural filePaths
func (widget *Widget) loadFilePaths() {
	var emptyArray []interface{}

	filePath := wtf.Config.UString("wtf.mods.textfile.filePath", "")
	filePaths := wtf.ToStrs(wtf.Config.UList("wtf.mods.textfile.filePaths", emptyArray))

	if filePath != "" {
		filePaths = append(filePaths, filePath)
	}

	widget.filePaths = filePaths
}

func (widget *Widget) plainText() string {
	filePath, _ := wtf.ExpandHomeDir(widget.currentFilePath())

	text, err := ioutil.ReadFile(filePath) // just pass the file name
	if err != nil {
		return err.Error()
	}
	return string(text)
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
	case "o":
		wtf.OpenFile(widget.currentFilePath())
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

	return event
}
