package trello

type SearchResult struct {
	Total       int
	TrelloCards map[string][]TrelloCard
}
