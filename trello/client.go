package trello

import (
	"fmt"

	"github.com/adlio/trello"
	"github.com/senorprogrammer/wtf/wtf"
)

func GetCards(client *trello.Client, lists map[string]string) (*SearchResult, error) {
	boardID, err := getBoardID(client)
	if err != nil {
		return nil, err
	}

	lists, err = getListIDs(client, boardID, lists)
	if err != nil {
		return nil, err
	}

	searchResult := &SearchResult{Total: 0}
	searchResult.TrelloCards = make(map[string][]TrelloCard)

	for listName, listID := range lists {
		cards, err := getCardsOnList(client, listID)
		if err != nil {
			return nil, err
		}
		searchResult.Total = searchResult.Total + len(cards)
		cardArray := make([]TrelloCard, 0)
		for _, card := range cards {
			trelloCard := TrelloCard{
				ID:          card.ID,
				List:        listName,
				Name:        card.Name,
				Description: card.Desc,
			}
			cardArray = append(cardArray, trelloCard)
		}
		searchResult.TrelloCards[listName] = cardArray
	}

	return searchResult, nil
}

func getBoardID(client *trello.Client) (string, error) {
	member, err := client.GetMember(wtf.Config.UString("wtf.mods.trello.username"), trello.Defaults())
	if err != nil {
		return "", err
	}

	boards, err := member.GetBoards(trello.Defaults())
	if err != nil {
		return "", err
	}

	for _, board := range boards {
		if board.Name == wtf.Config.UString("wtf.mods.trello.board") {
			return board.ID, nil
		}
	}

	return "", fmt.Errorf("could not find board with name %s", wtf.Config.UString("wtf.mods.trello.board"))
}

func getListIDs(client *trello.Client, boardID string, lists map[string]string) (map[string]string, error) {
	board, err := client.GetBoard(boardID, trello.Defaults())
	if err != nil {
		return nil, err
	}

	boardLists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return nil, err
	}

	for _, list := range boardLists {
		if _, ok := lists[list.Name]; ok {
			lists[list.Name] = list.ID
		}
	}

	return lists, nil
}

func getCardsOnList(client *trello.Client, listID string) ([]*trello.Card, error) {
	list, err := client.GetList(listID, trello.Defaults())
	if err != nil {
		return nil, err
	}

	cards, err := list.GetCards(trello.Defaults())
	if err != nil {
		return nil, err
	}

	return cards, nil
}
