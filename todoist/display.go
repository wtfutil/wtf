package todoist

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

func (w *Widget) display() {
	if len(w.list) == 0 {
		return
	}
	list := w.list[w.idx]

	w.View.SetTitle(fmt.Sprintf("%s- [green]%s[white] ", w.Name, list.Project.Name))
	str := wtf.SigilStr(len(w.list), w.idx, w.View) + "\n"

	for index, item := range list.items {
		if index == list.index {
			str = str + fmt.Sprintf("[%s]", wtf.Config.UString("wtf.colors.border.focused", "grey"))
		}
		str = str + fmt.Sprintf("| | %s[white]\n", tview.Escape(item.Content))
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
