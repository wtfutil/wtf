package toggl

type togglPerson struct {
	Since int64
	Data  struct {
		Fullname     string
		Time_entries []struct {
			Description string `json:"description"`
			Start       string `json:"start"`
			Stop        string `json:"stop"`
			Duration    int    `json:"duration"`
			Id          int    `json:"id"`
		}
	}
}
