package trello

type TrelloCard struct {
	ID          string
	Name        string
	List        string
	Description string
}

type TrelloList struct {
	ID   string
	Name string
}
