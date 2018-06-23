package trello

import (
	"fmt"
	"os"

	"github.com/adlio/trello"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Trello ", "trello", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	client := trello.NewClient(os.Getenv("WTF_TRELLO_APP_KEY"), os.Getenv("WTF_TRELLO_ACCESS_TOKEN"))

	// Get the cards
	searchResult, err := GetCards(client, getLists())
	widget.UpdateRefreshedAt()

	var content string
	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.Name)
		content = err.Error()
	} else {
		widget.View.SetWrap(false)
		widget.View.SetTitle(
			fmt.Sprintf(
				"[white]%s: [green]%s ",
				widget.Name,
				wtf.Config.UString("wtf.mods.trello.board"),
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

func getLists() map[string]string {
	list := make(map[string]string)
	// see if project is set to a single string
	configPath := "wtf.mods.trello.list"
	singleList, err := wtf.Config.String(configPath)
	if err == nil {
		list[singleList] = ""
		return list
	}
	// else, assume list
	multiList := wtf.Config.UList(configPath)
	for _, proj := range multiList {
		if str, ok := proj.(string); ok {
			list[str] = ""
		}
	}
	return list
}
