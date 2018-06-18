package circleci

type Build struct {
	AuthorEmail string `json:"author_email"`
	AuthorName  string `json:"author_name"`
	Branch      string `json:"branch"`
	BuildNum    int    `json:"build_num"`
	Reponame    string `json:"reponame"`
	Status      string `json:"status"`
}
