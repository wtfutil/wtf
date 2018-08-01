package textfile

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	"github.com/zyedidia/highlight"
)

const HelpText = `
  Keyboard commands for Textfile:

    /: Show/hide this help window
    o: Open the text file in the operating system
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	filePath string
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget("TextFile", "textfile", true),

		filePath: wtf.Config.UString("wtf.mods.textfile.filePath"),
	}

	widget.HelpfulWidget.SetView(widget.View)

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()
	widget.View.SetTitle(widget.ContextualTitle(widget.fileName()))

	filePath, _ := wtf.ExpandHomeDir(widget.filePath)

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fileData = []byte{}
	}

	if err != nil {
		widget.View.SetText(err.Error())
	} else {
		var (
			defs   []*highlight.Def
			buffer bytes.Buffer
		)

		gopath := os.Getenv("GOPATH")
		files, err := ioutil.ReadDir(gopath + "/src/github.com/senorprogrammer/wtf/vendor/github.com/zyedidia/highlight/syntax_files")
		if err != nil {
			widget.View.SetText(err.Error())
			return
		}

		// Iterate over available syntax files if any def matches.
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".yaml") {
				input, _ := ioutil.ReadFile(gopath + "/src/github.com/senorprogrammer/wtf/vendor/github.com/zyedidia/highlight/syntax_files/" + f.Name())
				d, err := highlight.ParseDef(input)
				if err != nil {
					fmt.Println(err)
					continue
				}
				defs = append(defs, d)
			}
		}

		highlight.ResolveIncludes(defs)
		// Attempt to detect the filetype.
		def := highlight.DetectFiletype(defs, filePath, bytes.Split(fileData, []byte("\n"))[0])
		// If no highlighter found, dont modify.
		if def == nil {
			widget.View.SetText(string(fileData))
			return
		}

		// Add highlight information to the text based on matches.
		h := highlight.NewHighlighter(def)
		matches := h.HighlightString(string(fileData))

		lines := strings.Split(string(fileData), "\n")
		for lineN, l := range lines {
			colN := 0
			for _, c := range l {
				if group, ok := matches[lineN][colN]; ok {
					// There are more possible groups available than just these ones
					if group == highlight.Groups["statement"] {
						buffer.WriteString("[green]")
					} else if group == highlight.Groups["identifier"] {
						buffer.WriteString("[blue]")
					} else if group == highlight.Groups["preproc"] {
						buffer.WriteString("[darkred]")
					} else if group == highlight.Groups["special"] {
						buffer.WriteString("[red]")
					} else if group == highlight.Groups["constant.string"] {
						buffer.WriteString("[cyan]")
					} else if group == highlight.Groups["constant"] {
						buffer.WriteString("[lightcyan]")
					} else if group == highlight.Groups["constant.specialChar"] {
						buffer.WriteString("[magenta]")
					} else if group == highlight.Groups["type"] {
						buffer.WriteString("[yellow]")
					} else if group == highlight.Groups["constant.number"] {
						buffer.WriteString("[darkcyan]")
					} else if group == highlight.Groups["comment"] {
						buffer.WriteString("[darkgreen]")
					} else {
						buffer.WriteString("[default]")
					}
				}
				buffer.WriteString(string(c))
				colN++
			}
			if group, ok := matches[lineN][colN]; ok {
				if group == highlight.Groups["default"] || group == highlight.Groups[""] {
					buffer.WriteString("[default]") // Handle default fallback after setting.
				}
			}

			buffer.WriteString("\n")
		}

		widget.View.SetText(buffer.String())
	}
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) fileName() string {
	return filepath.Base(widget.filePath)
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
		return nil
	case "o":
		wtf.OpenFile(widget.filePath)
		return nil
	}

	return event
}
