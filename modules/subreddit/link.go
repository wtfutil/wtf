package subreddit

type Link struct {
	Score     int    `json:"ups"`
	Title     string `json:"title"`
	ItemURL   string `json:"url"`
	Permalink string `json:"permalink"`
}

type RedditDocument struct {
	Data Subreddit `json:"data"`
}

type RedditLinkDocument struct {
	Data Link `json:"data"`
}

type Subreddit struct {
	Children []RedditLinkDocument `json:"Children"`
}
