package view

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

type TextWidget struct {
	Base
	View *tview.TextView
}

func NewTextWidget(app *tview.Application, commonSettings *cfg.Common, focusable bool) TextWidget {
	widget := TextWidget{
		Base: NewBase(app, commonSettings, focusable),
	}

	widget.View = widget.addView(widget.bordered)

	return widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *TextWidget) TextView() *tview.TextView {
	return widget.View
}

func (widget *TextWidget) Redraw(title, text string, wrap bool) {
	widget.app.QueueUpdateDraw(func() {
		widget.View.Clear()
		widget.View.SetWrap(wrap)
		widget.View.SetTitle(widget.ContextualTitle(title))
		widget.View.SetText(text)
	})
}

/* -------------------- Unexported Functions -------------------- */

func (widget *TextWidget) addView(bordered bool) *tview.TextView {
	view := tview.NewTextView()

	view.SetBackgroundColor(wtf.ColorFor(widget.commonSettings.Colors.Background))
	view.SetBorder(bordered)
	view.SetBorderColor(wtf.ColorFor(widget.BorderColor()))
	view.SetDynamicColors(true)
	view.SetTextColor(wtf.ColorFor(widget.commonSettings.Colors.Text))
	view.SetTitleColor(wtf.ColorFor(widget.commonSettings.Colors.Title))
	view.SetWrap(false)

	return view
}
