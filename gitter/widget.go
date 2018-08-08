package gitter

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	"strconv"
)

const HelpText = `
 Keyboard commands for Gitter:

   /: Show/hide this help window
   j: Select the next message in the list
   k: Select the previous message in the list
   r: Refresh the data

   arrow down: Select the next message in the list
   arrow up:   Select the previous message in the list
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	messages []Message
	selected int
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget("Gitter", "gitter", true),
	}

	widget.HelpfulWidget.SetView(widget.View)
	widget.unselect()

	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	room, err := GetRoom(wtf.Config.UString("wtf.mods.gitter.roomUri", "wtfutil/Lobby"))
	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.Name)
		widget.View.SetText(err.Error())
		return
	}

	if room == nil {
		return
	}

	messages, err := GetMessages(room.ID, wtf.Config.UInt("wtf.mods.gitter.numberOfMessages", 10))
	widget.UpdateRefreshedAt()

	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.Name)
		widget.View.SetText(err.Error())
	} else {
		widget.messages = messages
	}

	widget.display()
	widget.View.ScrollToEnd()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	if widget.messages == nil {
		return
	}

	widget.View.SetWrap(true)
	widget.View.Clear()
	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s - %s", widget.Name, wtf.Config.UString("wtf.mods.gitter.roomUri", "wtfutil/Lobby"))))
	widget.View.SetText(widget.contentFrom(widget.messages))
	widget.View.Highlight(strconv.Itoa(widget.selected)).ScrollToHighlight()
}

func (widget *Widget) contentFrom(messages []Message) string {
	var str string
	for idx, message := range messages {
		str = str + fmt.Sprintf(
			`["%d"][""][%s] [blue]%s [lightslategray]%s: [%s]%s [aqua]%s`,
			idx,
			widget.rowColor(idx),
			message.From.DisplayName,
			message.From.Username,
			widget.rowColor(idx),
			message.Text,
			message.Sent.Format("Jan 02, 15:04 MST"),
		)

		str = str + "\n"
	}

	return str
}

func (widget *Widget) rowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		return wtf.DefaultFocussedRowColor()
	}

	return wtf.RowColor("gitter", idx)
}

func (widget *Widget) next() {
	widget.selected++
	if widget.messages != nil && widget.selected >= len(widget.messages) {
		widget.selected = 0
	}

	widget.display()
}

func (widget *Widget) prev() {
	widget.selected--
	if widget.selected < 0 && widget.messages != nil {
		widget.selected = len(widget.messages) - 1
	}

	widget.display()
}

func (widget *Widget) openMessage() {
	sel := widget.selected
	if sel >= 0 && widget.messages != nil && sel < len(widget.messages) {
		message := &widget.messages[widget.selected]
		wtf.OpenFile(message.Text)
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
	widget.display()
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
	case "j":
		widget.next()
		return nil
	case "k":
		widget.prev()
		return nil
	case "r":
		widget.Refresh()
		return nil
	}

	switch event.Key() {
	case tcell.KeyDown:
		widget.next()
		return nil
	case tcell.KeyEsc:
		widget.unselect()
		return event
	case tcell.KeyUp:
		widget.prev()
		return nil
	default:
		return event
	}
}
