package hackernews

import "fmt"

const (
	hnStoryPath = "https://news.ycombinator.com/item?id="
)

// Story represents a story submission on HackerNews
type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

// CommentLink return the link to the HackerNews story comments page
func (story *Story) CommentLink() string {
	return fmt.Sprintf("%s%d", hnStoryPath, story.ID)
}

// Link returns the link to a story. If the story has an external link, that is returned
// If the story has no external link, the HackerNews comments link is returned instead
func (story *Story) Link() string {
	if story.URL != "" {
		return story.URL
	}

	// Fall back to the HackerNews comment link
	return story.CommentLink()
}
