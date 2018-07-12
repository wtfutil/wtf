package todoist

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

const checkWidth = 4

func (w *Widget) display() {
	if len(w.list) == 0 {
		return
	}
	list := w.list[w.idx]

	w.View.SetTitle(fmt.Sprintf("%s- [green]%s[white] ", w.Name, list.Project.Name))
	str := wtf.SigilStr(len(w.list), w.idx, w.View) + "\n"

	maxLen := w.list[w.idx].LongestLine()

	for index, item := range list.items {
		foreColor, backColor := "white", wtf.Config.UString("wtf.colors.background", "black")

		if index == list.index {
			foreColor = wtf.Config.UString("wtf.colors.highlight.fore", "black")
			backColor = wtf.Config.UString("wtf.colors.highlight.back", "orange")
		}

		row := fmt.Sprintf(
			"[%s:%s]| | %s[white]",
			foreColor,
			backColor,
			tview.Escape(item.Content),
		)

		row = row + wtf.PadRow((checkWidth+len(item.Content)), (checkWidth+maxLen+1)) + "\n"
		str = str + row
	}

	w.View.Clear()
	w.View.SetText(str)
}

func (w *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	if len(w.list) == 0 {
		return event
	}

	switch string(event.Rune()) {
	case "r":
		w.Refresh()
		return nil
	case "d":
		w.Delete()
		return nil
	case "c":
		w.Close()
		return nil
	}

	switch fromVim(event) {
	case tcell.KeyLeft:
		w.Prev()
		return nil
	case tcell.KeyRight:
		w.Next()
		return nil
	case tcell.KeyUp:
		w.UP()
		return nil
	case tcell.KeyDown:
		w.Down()
		return nil
	}

	return event
}
