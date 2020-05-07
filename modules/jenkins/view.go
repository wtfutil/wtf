package jenkins

type View struct {
	Description          string `json:"description"`
	Jobs                 []Job  `json:"jobs"`
	ActiveConfigurations []Job  `json:"activeConfigurations"`
	Name                 string `json:"name"`
	Url                  string `json:"url"`
}
