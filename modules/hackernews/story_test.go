package hackernews

import (
	"testing"

	"gotest.tools/assert"
)

func Test_CommentLink(t *testing.T) {
	story := Story{
		ID: 3,
	}

	assert.Equal(t, "https://news.ycombinator.com/item?id=3", story.CommentLink())
}

func Test_Link(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		url      string
		expected string
	}{
		{
			name:     "no external link",
			id:       1,
			url:      "",
			expected: "https://news.ycombinator.com/item?id=1",
		},
		{
			name:     "with external link",
			id:       1,
			url:      "https://www.link.ca",
			expected: "https://www.link.ca",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			story := Story{
				ID:  tt.id,
				URL: tt.url,
			}

			actual := story.Link()

			assert.Equal(t, tt.expected, actual)
		})
	}
}
