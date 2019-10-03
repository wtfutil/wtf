package trello

import (
	"fmt"

	"github.com/adlio/trello"
)

func GetCards(client *trello.Client, username string, boardName string, listNames []string) (*SearchResult, error) {
	boardID, err := getBoardID(client, username, boardName)
	if err != nil {
		return nil, err
	}

	lists, err := getLists(client, boardID, listNames)
	if err != nil {
		return nil, err
	}

	searchResult := &SearchResult{Total: 0}
	searchResult.TrelloCards = make(map[string][]TrelloCard)

	for _, list := range lists {
		cards, err := getCardsOnList(client, list.ID)
		if err != nil {
			return nil, err
		}

		searchResult.Total += len(cards)
		cardArray := make([]TrelloCard, 0)

		for _, card := range cards {
			trelloCard := TrelloCard{
				ID:          card.ID,
				List:        list.Name,
				Name:        card.Name,
				Description: card.Desc,
			}
			cardArray = append(cardArray, trelloCard)
		}

		searchResult.TrelloCards[list.Name] = cardArray
	}

	return searchResult, nil
}

func getBoardID(client *trello.Client, username, boardName string) (string, error) {
	member, err := client.GetMember(username, trello.Defaults())
	if err != nil {
		return "", err
	}

	boards, err := member.GetBoards(trello.Defaults())
	if err != nil {
		return "", err
	}

	for _, board := range boards {
		if board.Name == boardName {
			return board.ID, nil
		}
	}

	return "", fmt.Errorf("could not find board with name %s", boardName)
}

func getLists(client *trello.Client, boardID string, listNames []string) ([]TrelloList, error) {
	comparison := make(map[string]string, len(listNames))
	results := []TrelloList{}
	//convert to a map for comparison
	for _, item := range listNames {
		comparison[item] = ""
	}
	board, err := client.GetBoard(boardID, trello.Defaults())
	if err != nil {
		return nil, err
	}

	boardLists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return nil, err
	}

	for _, list := range boardLists {
		if _, ok := comparison[list.Name]; ok {
			results = append(results, TrelloList{ID: list.ID, Name: list.Name})
		}
	}

	return results, nil
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
