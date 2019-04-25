package trello

import (
	"fmt"

	"github.com/adlio/trello"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	app      *tview.Application
	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, false),

		app:      app,
		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	client := trello.NewClient(
		widget.settings.apiKey,
		widget.settings.accessToken,
	)

	// Get the cards
	searchResult, err := GetCards(
		client,
		widget.settings.username,
		widget.settings.board,
		widget.settings.list,
	)

	var title string
	var content string

	if err != nil {
		widget.View.SetWrap(true)
		title = widget.Name()
		content = err.Error()
	} else {
		widget.View.SetWrap(false)
		title = fmt.Sprintf(
			"[white]%s: [green]%s ",
			widget.Name(),
			widget.settings.board,
		)
		content = widget.contentFrom(searchResult)
	}

	widget.app.QueueUpdateDraw(func() {
		widget.View.SetTitle(title)
		widget.View.SetText(content)
	})
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(searchResult *SearchResult) string {
	str := ""

	for list, cardArray := range searchResult.TrelloCards {
		str = fmt.Sprintf("%s [red]Cards in %s[white]\n", str, list)
		for _, card := range cardArray {
			str = fmt.Sprintf("%s [green]%s[white]\n", str, card.Name)
		}
		str = fmt.Sprintf("%s\n", str)
	}

	return str
}
