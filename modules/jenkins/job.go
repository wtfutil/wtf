package jenkins

type Job struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Color string `json:"color"`
}
