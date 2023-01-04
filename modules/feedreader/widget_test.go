package feedreader

import (
	"testing"

	"github.com/mmcdole/gofeed"
	"gotest.tools/assert"
)

func Test_getShowText(t *testing.T) {
	tests := []struct {
		name     string
		feedItem *FeedItem
		showType ShowType
		expected string
	}{
		{
			name:     "with nil FeedItem",
			feedItem: nil,
			showType: SHOW_TITLE,
			expected: "",
		},
		{
			name: "with plain title",
			feedItem: &FeedItem{
				item: &gofeed.Item{Title: "Cats and Dogs"},
			},
			showType: SHOW_TITLE,
			expected: "[white]Cats and Dogs",
		},
		{
			name: "with escaped title",
			feedItem: &FeedItem{
				item: &gofeed.Item{Title: "&lt;Cats and Dogs&gt;"},
			},
			showType: SHOW_TITLE,
			expected: "[white]<Cats and Dogs>",
		},
		{
			name: "with unescaped title",
			feedItem: &FeedItem{
				item: &gofeed.Item{Title: "<Cats and Dogs>"},
			},
			showType: SHOW_TITLE,
			expected: "[white]<Cats and Dogs>",
		},
		{
			name: "with source-title",
			feedItem: &FeedItem{
				sourceTitle: "WTF",
				item:        &gofeed.Item{Title: "<Cats and Dogs>"},
			},
			showType: SHOW_TITLE,
			expected: "[green]WTF [white]<Cats and Dogs>",
		},
		{
			name: "with link",
			feedItem: &FeedItem{
				item: &gofeed.Item{Title: "Cats and Dogs", Link: "https://cats.com/dog.xml"},
			},
			showType: SHOW_LINK,
			expected: "https://cats.com/dog.xml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{
				settings: &Settings{
					colors: colors{
						source:      "green",
						publishDate: "orange",
					},
					showSource: true,
				},
				showType: tt.showType,
			}

			actual := widget.getShowText(tt.feedItem, "white")

			assert.Equal(t, tt.expected, actual)
		})
	}
}
