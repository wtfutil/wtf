package jenkins

type View struct {
	Class       string   `json:"_class"`
	Description string   `json:"description"`
	Jobs        []Job    `json:"jobs"`
	Name        string   `json:"name"`
	Property    []string `json:"property"`
	Url         string   `json:"url"`
}
