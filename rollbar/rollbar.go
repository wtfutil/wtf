package rollbar

type ActiveItems struct {
	Results Result `json:"result"`
}
type Item struct {
	Environment      string `json:"environment"`
	Title            string `json:"title"`
	Platform         string `json:"platform"`
	Status           string `json:"status"`
	TotalOccurrences int    `json:"total_occurrences"`
	Level            string `json:"level"`
	ID               int    `json:"counter"`
}
type Result struct {
	Items []Item `json:"items"`
}
