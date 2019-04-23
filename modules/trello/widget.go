package trello

import (
	"fmt"

	"github.com/adlio/trello"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	settings *Settings
}

func NewWidget(refreshChan chan<- string, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(refreshChan, settings.common, false),

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

	var content string
	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.Name())
		content = err.Error()
	} else {
		widget.View.SetWrap(false)
		widget.View.SetTitle(
			fmt.Sprintf(
				"[white]%s: [green]%s ",
				widget.Name(),
				widget.settings.board,
			),
		)
		content = widget.contentFrom(searchResult)
	}

	widget.View.SetText(content)
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
