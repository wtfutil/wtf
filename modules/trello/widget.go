package trello

import (
	"fmt"

	"github.com/adlio/trello"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, false),

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

	wrap := false
	if err != nil {
		wrap = true
		title = widget.CommonSettings.Title
		content = err.Error()
	} else {
		widget.View.SetWrap(false)
		title = fmt.Sprintf(
			"[white]%s: [green]%s ",
			widget.CommonSettings.Title,
			widget.settings.board,
		)
		content = widget.contentFrom(searchResult)
	}

	widget.Redraw(title, content, wrap)
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
