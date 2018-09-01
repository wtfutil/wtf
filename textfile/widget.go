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
	wtf.MultiSourceWidget
	wtf.TextWidget
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget:     wtf.NewHelpfulWidget(app, pages, HelpText),
		MultiSourceWidget: wtf.NewMultiSourceWidget(),
		TextWidget:        wtf.NewTextWidget("TextFile", "textfile", true),
	}

	widget.LoadSources("textfile", "fileName", "fileNames")

	widget.HelpfulWidget.SetView(widget.View)

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Next() {
	widget.MultiSourceWidget.Idx = widget.MultiSourceWidget.Idx + 1
	if widget.Idx == len(widget.Sources) {
		widget.Idx = 0
	}

	widget.display()
}

func (widget *Widget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.Sources) - 1
	}

	widget.display()
}

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	widget.View.SetTitle(widget.ContextualTitle(widget.fileName()))

	text := wtf.SigilStr(len(widget.Sources), widget.Idx, widget.View) + "\n"

	if wtf.Config.UBool("wtf.mods.textfile.format", false) {
		text = text + widget.formattedText()
	} else {
		text = text + widget.plainText()
	}

	widget.View.SetText(text)
}

func (widget *Widget) fileName() string {
	return filepath.Base(widget.CurrentSource())
}

func (widget *Widget) formattedText() string {
	filePath, _ := wtf.ExpandHomeDir(widget.CurrentSource())

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

func (widget *Widget) plainText() string {
	filePath, _ := wtf.ExpandHomeDir(widget.CurrentSource())

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
		wtf.OpenFile(widget.CurrentSource())
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
