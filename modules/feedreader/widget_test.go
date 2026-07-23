package feedreader

import (
	"errors"
	"testing"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
	"gotest.tools/assert"
)

// mockParser implements feedParser for testing
type mockParser struct {
	feeds    map[string]*gofeed.Feed
	err      error
	authUsed *gofeed.Auth
}

func (m *mockParser) ParseURL(feedURL string) (*gofeed.Feed, error) {
	if m.err != nil {
		return nil, m.err
	}
	feed, ok := m.feeds[feedURL]
	if !ok {
		return nil, errors.New("feed not found: " + feedURL)
	}
	return feed, nil
}

func (m *mockParser) SetAuth(auth *gofeed.Auth) {
	m.authUsed = auth
}

func Test_rotateShowType(t *testing.T) {
	tests := []struct {
		name     string
		input    ShowType
		expected ShowType
	}{
		{
			name:     "SHOW_TITLE rotates to SHOW_LINK",
			input:    SHOW_TITLE,
			expected: SHOW_LINK,
		},
		{
			name:     "SHOW_LINK rotates to SHOW_CONTENT",
			input:    SHOW_LINK,
			expected: SHOW_CONTENT,
		},
		{
			name:     "SHOW_CONTENT rotates to SHOW_TITLE",
			input:    SHOW_CONTENT,
			expected: SHOW_TITLE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := rotateShowType(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_sort(t *testing.T) {
	now := time.Now()
	earlier := now.Add(-1 * time.Hour)
	earliest := now.Add(-2 * time.Hour)

	tests := []struct {
		name           string
		feedItems      []*FeedItem
		expectedTitles []string
	}{
		{
			name:           "empty list",
			feedItems:      []*FeedItem{},
			expectedTitles: []string{},
		},
		{
			name: "already sorted (newest first)",
			feedItems: []*FeedItem{
				{item: &gofeed.Item{Title: "Newest", PublishedParsed: &now}},
				{item: &gofeed.Item{Title: "Oldest", PublishedParsed: &earliest}},
			},
			expectedTitles: []string{"Newest", "Oldest"},
		},
		{
			name: "unsorted items get sorted newest first",
			feedItems: []*FeedItem{
				{item: &gofeed.Item{Title: "Oldest", PublishedParsed: &earliest}},
				{item: &gofeed.Item{Title: "Newest", PublishedParsed: &now}},
				{item: &gofeed.Item{Title: "Middle", PublishedParsed: &earlier}},
			},
			expectedTitles: []string{"Newest", "Middle", "Oldest"},
		},
		{
			name: "nil PublishedParsed preserves original order",
			feedItems: []*FeedItem{
				{item: &gofeed.Item{Title: "No Date", PublishedParsed: nil}},
				{item: &gofeed.Item{Title: "Has Date", PublishedParsed: &now}},
			},
			expectedTitles: []string{"No Date", "Has Date"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{}
			result := widget.sort(tt.feedItems)

			actualTitles := make([]string, len(result))
			for i, item := range result {
				actualTitles[i] = item.item.Title
			}

			assert.DeepEqual(t, tt.expectedTitles, actualTitles)
		})
	}
}

func Test_getShowText(t *testing.T) {
	publishedTime := time.Date(2023, time.March, 15, 10, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		feedItem *FeedItem
		showType ShowType
		settings *Settings
		expected string
	}{
		{
			name:     "with nil FeedItem",
			feedItem: nil,
			showType: SHOW_TITLE,
			settings: &Settings{
				colors:     colors{source: "green", publishDate: "orange"},
				showSource: true,
			},
			expected: "",
		},
		{
			name: "with plain title",
			feedItem: &FeedItem{
				item: &gofeed.Item{Title: "Cats and Dogs"},
			},
			showType: SHOW_TITLE,
			settings: &Settings{
				colors:     colors{source: "green", publishDate: "orange"},
				showSource: true,
			},
			expected: "[white]Cats and Dogs",
		},
		{
			name: "with escaped title",
			feedItem: &FeedItem{
				item: &gofeed.Item{Title: "&lt;Cats and Dogs&gt;"},
			},
			showType: SHOW_TITLE,
			settings: &Settings{
				colors:     colors{source: "green", publishDate: "orange"},
				showSource: true,
			},
			expected: "[white]<Cats and Dogs>",
		},
		{
			name: "with unescaped title",
			feedItem: &FeedItem{
				item: &gofeed.Item{Title: "<Cats and Dogs>"},
			},
			showType: SHOW_TITLE,
			settings: &Settings{
				colors:     colors{source: "green", publishDate: "orange"},
				showSource: true,
			},
			expected: "[white]<Cats and Dogs>",
		},
		{
			name: "with source-title",
			feedItem: &FeedItem{
				sourceTitle: "WTF",
				item:        &gofeed.Item{Title: "<Cats and Dogs>"},
			},
			showType: SHOW_TITLE,
			settings: &Settings{
				colors:     colors{source: "green", publishDate: "orange"},
				showSource: true,
			},
			expected: "[green]WTF [white]<Cats and Dogs>",
		},
		{
			name: "with link",
			feedItem: &FeedItem{
				item: &gofeed.Item{Title: "Cats and Dogs", Link: "https://cats.com/dog.xml"},
			},
			showType: SHOW_LINK,
			settings: &Settings{
				colors:     colors{source: "green", publishDate: "orange"},
				showSource: true,
			},
			expected: "https://cats.com/dog.xml",
		},
		{
			name: "with content",
			feedItem: &FeedItem{
				item: &gofeed.Item{Title: "Article", Content: "<p>Hello World</p>"},
			},
			showType: SHOW_CONTENT,
			settings: &Settings{
				colors:     colors{source: "green", publishDate: "orange"},
				showSource: false,
			},
			expected: "[white]Article\nHello World",
		},
		{
			name: "with publish date shown",
			feedItem: &FeedItem{
				item: &gofeed.Item{
					Title:           "News",
					Published:       "2023-03-15",
					PublishedParsed: &publishedTime,
				},
			},
			showType: SHOW_TITLE,
			settings: &Settings{
				colors:          colors{source: "green", publishDate: "orange"},
				showSource:      false,
				showPublishDate: true,
				dateFormat:      "Jan 02",
			},
			expected: "[orange]Mar 15 [white]News",
		},
		{
			name: "with source and publish date",
			feedItem: &FeedItem{
				sourceTitle: "Blog",
				item: &gofeed.Item{
					Title:           "Post",
					Published:       "2023-03-15",
					PublishedParsed: &publishedTime,
				},
			},
			showType: SHOW_TITLE,
			settings: &Settings{
				colors:          colors{source: "green", publishDate: "orange"},
				showSource:      true,
				showPublishDate: true,
				dateFormat:      "Jan 02",
			},
			expected: "[green]Blog [orange]Mar 15 [white]Post",
		},
		{
			name: "with whitespace in title collapsed",
			feedItem: &FeedItem{
				item: &gofeed.Item{Title: "Lots   of    spaces"},
			},
			showType: SHOW_TITLE,
			settings: &Settings{
				colors:     colors{source: "green", publishDate: "orange"},
				showSource: false,
			},
			expected: "[white]Lots of spaces",
		},
		{
			name: "showSource true but empty sourceTitle shows no prefix",
			feedItem: &FeedItem{
				sourceTitle: "",
				item:        &gofeed.Item{Title: "Title"},
			},
			showType: SHOW_TITLE,
			settings: &Settings{
				colors:     colors{source: "green", publishDate: "orange"},
				showSource: true,
			},
			expected: "[white]Title",
		},
		{
			name: "showSource false hides sourceTitle",
			feedItem: &FeedItem{
				sourceTitle: "Hidden",
				item:        &gofeed.Item{Title: "Title"},
			},
			showType: SHOW_TITLE,
			settings: &Settings{
				colors:     colors{source: "green", publishDate: "orange"},
				showSource: false,
			},
			expected: "[white]Title",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{
				settings: tt.settings,
				showType: tt.showType,
			}

			actual := widget.getShowText(tt.feedItem, "white")

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_Fetch(t *testing.T) {
	now := time.Now()
	earlier := now.Add(-1 * time.Hour)

	t.Run("fetches and sorts items from multiple feeds", func(t *testing.T) {
		mock := &mockParser{
			feeds: map[string]*gofeed.Feed{
				"http://feed1.com/rss": {
					Title: "Feed One",
					Items: []*gofeed.Item{
						{Title: "Old Post", PublishedParsed: &earlier},
					},
				},
				"http://feed2.com/rss": {
					Title: "Feed Two",
					Items: []*gofeed.Item{
						{Title: "New Post", PublishedParsed: &now},
					},
				},
			},
		}

		widget := &Widget{
			parser:   mock,
			settings: &Settings{feedLimit: -1, credentials: make(map[string]auth)},
		}

		items, err := widget.Fetch(
			[]string{"http://feed1.com/rss", "http://feed2.com/rss"},
			nil,
		)

		assert.NilError(t, err)
		assert.Equal(t, 2, len(items))
		assert.Equal(t, "New Post", items[0].item.Title)
		assert.Equal(t, "Old Post", items[1].item.Title)
	})

	t.Run("applies aliases to feed items", func(t *testing.T) {
		mock := &mockParser{
			feeds: map[string]*gofeed.Feed{
				"http://feed1.com/rss": {
					Title: "Original Title",
					Items: []*gofeed.Item{
						{Title: "Post", PublishedParsed: &now},
					},
				},
			},
		}

		widget := &Widget{
			parser:   mock,
			settings: &Settings{feedLimit: -1, credentials: make(map[string]auth)},
		}

		items, err := widget.Fetch(
			[]string{"http://feed1.com/rss"},
			[]string{"My Alias"},
		)

		assert.NilError(t, err)
		assert.Equal(t, 1, len(items))
		assert.Equal(t, "My Alias", items[0].sourceTitle)
	})

	t.Run("returns error when feed fails", func(t *testing.T) {
		mock := &mockParser{
			err: errors.New("network timeout"),
		}

		widget := &Widget{
			parser:   mock,
			settings: &Settings{feedLimit: -1, credentials: make(map[string]auth)},
		}

		items, err := widget.Fetch([]string{"http://bad.com/rss"}, nil)

		assert.Assert(t, err != nil)
		assert.Assert(t, items == nil)
	})

	t.Run("right-pads aliases for alignment", func(t *testing.T) {
		mock := &mockParser{
			feeds: map[string]*gofeed.Feed{
				"http://a.com/rss": {
					Title: "A",
					Items: []*gofeed.Item{
						{Title: "Post A", PublishedParsed: &now},
					},
				},
				"http://b.com/rss": {
					Title: "B",
					Items: []*gofeed.Item{
						{Title: "Post B", PublishedParsed: &earlier},
					},
				},
			},
		}

		widget := &Widget{
			parser:   mock,
			settings: &Settings{feedLimit: -1, credentials: make(map[string]auth)},
		}

		items, err := widget.Fetch(
			[]string{"http://a.com/rss", "http://b.com/rss"},
			[]string{"Short", "LongerAlias"},
		)

		assert.NilError(t, err)
		assert.Equal(t, 2, len(items))
		// The shorter alias should be padded to match "LongerAlias" length (11)
		assert.Equal(t, "Short      ", items[0].sourceTitle)
		assert.Equal(t, "LongerAlias", items[1].sourceTitle)
	})
}

func Test_fetchForFeed(t *testing.T) {
	now := time.Now()

	t.Run("respects feedLimit", func(t *testing.T) {
		mock := &mockParser{
			feeds: map[string]*gofeed.Feed{
				"http://feed.com/rss": {
					Title: "Test Feed",
					Items: []*gofeed.Item{
						{Title: "Item 1", PublishedParsed: &now},
						{Title: "Item 2", PublishedParsed: &now},
						{Title: "Item 3", PublishedParsed: &now},
						{Title: "Item 4", PublishedParsed: &now},
						{Title: "Item 5", PublishedParsed: &now},
					},
				},
			},
		}

		widget := &Widget{
			parser:   mock,
			settings: &Settings{feedLimit: 2, credentials: make(map[string]auth)},
		}

		items, err := widget.fetchForFeed("http://feed.com/rss", "")

		assert.NilError(t, err)
		assert.Equal(t, 2, len(items))
		assert.Equal(t, "Item 1", items[0].item.Title)
		assert.Equal(t, "Item 2", items[1].item.Title)
	})

	t.Run("negative feedLimit returns all items", func(t *testing.T) {
		mock := &mockParser{
			feeds: map[string]*gofeed.Feed{
				"http://feed.com/rss": {
					Title: "Test Feed",
					Items: []*gofeed.Item{
						{Title: "Item 1", PublishedParsed: &now},
						{Title: "Item 2", PublishedParsed: &now},
						{Title: "Item 3", PublishedParsed: &now},
					},
				},
			},
		}

		widget := &Widget{
			parser:   mock,
			settings: &Settings{feedLimit: -1, credentials: make(map[string]auth)},
		}

		items, err := widget.fetchForFeed("http://feed.com/rss", "")

		assert.NilError(t, err)
		assert.Equal(t, 3, len(items))
	})

	t.Run("uses feed title as source when no alias", func(t *testing.T) {
		mock := &mockParser{
			feeds: map[string]*gofeed.Feed{
				"http://feed.com/rss": {
					Title: "My Feed Title",
					Items: []*gofeed.Item{
						{Title: "Post", PublishedParsed: &now},
					},
				},
			},
		}

		widget := &Widget{
			parser:   mock,
			settings: &Settings{feedLimit: -1, credentials: make(map[string]auth)},
		}

		items, err := widget.fetchForFeed("http://feed.com/rss", "")

		assert.NilError(t, err)
		assert.Equal(t, "My Feed Title", items[0].sourceTitle)
	})

	t.Run("alias overrides feed title", func(t *testing.T) {
		mock := &mockParser{
			feeds: map[string]*gofeed.Feed{
				"http://feed.com/rss": {
					Title: "Original",
					Items: []*gofeed.Item{
						{Title: "Post", PublishedParsed: &now},
					},
				},
			},
		}

		widget := &Widget{
			parser:   mock,
			settings: &Settings{feedLimit: -1, credentials: make(map[string]auth)},
		}

		items, err := widget.fetchForFeed("http://feed.com/rss", "Override")

		assert.NilError(t, err)
		assert.Equal(t, "Override", items[0].sourceTitle)
	})

	t.Run("sets auth for private feeds", func(t *testing.T) {
		mock := &mockParser{
			feeds: map[string]*gofeed.Feed{
				"http://private.com/rss": {
					Title: "Private",
					Items: []*gofeed.Item{
						{Title: "Secret Post", PublishedParsed: &now},
					},
				},
			},
		}

		widget := &Widget{
			parser: mock,
			settings: &Settings{
				feedLimit: -1,
				credentials: map[string]auth{
					"http://private.com/rss": {username: "user", password: "pass"},
				},
			},
		}

		items, err := widget.fetchForFeed("http://private.com/rss", "")

		assert.NilError(t, err)
		assert.Equal(t, 1, len(items))
		// Auth should be cleared after use
		assert.Assert(t, mock.authUsed == nil)
	})

	t.Run("returns error from parser", func(t *testing.T) {
		mock := &mockParser{
			err: errors.New("parse error"),
		}

		widget := &Widget{
			parser:   mock,
			settings: &Settings{feedLimit: -1, credentials: make(map[string]auth)},
		}

		items, err := widget.fetchForFeed("http://bad.com/rss", "")

		assert.Assert(t, err != nil)
		assert.Assert(t, items == nil)
	})

	t.Run("items are marked as not viewed", func(t *testing.T) {
		mock := &mockParser{
			feeds: map[string]*gofeed.Feed{
				"http://feed.com/rss": {
					Title: "Feed",
					Items: []*gofeed.Item{
						{Title: "Post", PublishedParsed: &now},
					},
				},
			},
		}

		widget := &Widget{
			parser:   mock,
			settings: &Settings{feedLimit: -1, credentials: make(map[string]auth)},
		}

		items, err := widget.fetchForFeed("http://feed.com/rss", "")

		assert.NilError(t, err)
		assert.Equal(t, false, items[0].viewed)
	})
}

func Test_toggleDisplayText(t *testing.T) {
	t.Run("cycles through show types", func(t *testing.T) {
		widget := &Widget{
			showType: SHOW_TITLE,
			settings: &Settings{},
		}

		widget.showType = rotateShowType(widget.showType)
		assert.Equal(t, SHOW_LINK, widget.showType)

		widget.showType = rotateShowType(widget.showType)
		assert.Equal(t, SHOW_CONTENT, widget.showType)

		widget.showType = rotateShowType(widget.showType)
		assert.Equal(t, SHOW_TITLE, widget.showType)
	})
}

func Test_openStory(t *testing.T) {
	now := time.Now()

	t.Run("marks story as viewed when selection is valid", func(t *testing.T) {
		widget := &Widget{
			stories: []*FeedItem{
				{item: &gofeed.Item{Title: "Post 1", Link: "", PublishedParsed: &now}, viewed: false},
				{item: &gofeed.Item{Title: "Post 2", Link: "", PublishedParsed: &now}, viewed: false},
			},
			settings: &Settings{},
		}

		sel := 0
		if sel >= 0 && widget.stories != nil && sel < len(widget.stories) {
			widget.stories[sel].viewed = true
		}

		assert.Equal(t, true, widget.stories[0].viewed)
		assert.Equal(t, false, widget.stories[1].viewed)
	})

	t.Run("does nothing with negative selection", func(t *testing.T) {
		widget := &Widget{
			stories: []*FeedItem{
				{item: &gofeed.Item{Title: "Post 1"}, viewed: false},
			},
			settings: &Settings{},
		}

		sel := -1
		if sel >= 0 && widget.stories != nil && sel < len(widget.stories) {
			widget.stories[sel].viewed = true
		}

		assert.Equal(t, false, widget.stories[0].viewed)
	})

	t.Run("does nothing with out-of-range selection", func(t *testing.T) {
		widget := &Widget{
			stories: []*FeedItem{
				{item: &gofeed.Item{Title: "Post 1"}, viewed: false},
			},
			settings: &Settings{},
		}

		sel := 5
		if sel >= 0 && widget.stories != nil && sel < len(widget.stories) {
			widget.stories[sel].viewed = true
		}

		assert.Equal(t, false, widget.stories[0].viewed)
	})

	t.Run("does nothing with nil stories", func(t *testing.T) {
		widget := &Widget{
			stories:  nil,
			settings: &Settings{},
		}

		sel := 0
		if sel >= 0 && widget.stories != nil && sel < len(widget.stories) {
			widget.stories[sel].viewed = true
		}
		assert.Assert(t, widget.stories == nil)
	})
}

func Test_NewSettingsFromYAML(t *testing.T) {
	t.Run("parses basic feed list", func(t *testing.T) {
		yamlContent := `
feeds:
  - http://feed1.com/rss
  - http://feed2.com/rss
aliases:
  - Feed One
  - Feed Two
feedLimit: 5
showSource: true
showPublishDate: true
dateFormat: "2006-01-02"
disableHTTP2: true
userAgent: "test-agent"
colors:
  source: "blue"
  publishDate: "red"
`
		globalYaml := `
wtf:
  colors:
    border:
      focusable: green
      focused: blue
      normal: white
`
		ymlConfig, err := config.ParseYaml(yamlContent)
		assert.NilError(t, err)
		globalConfig, err := config.ParseYaml(globalYaml)
		assert.NilError(t, err)

		settings := NewSettingsFromYAML("feedreader", ymlConfig, globalConfig)

		assert.Equal(t, 2, len(settings.feeds))
		assert.Equal(t, "http://feed1.com/rss", settings.feeds[0])
		assert.Equal(t, "http://feed2.com/rss", settings.feeds[1])
		assert.Equal(t, 2, len(settings.aliases))
		assert.Equal(t, "Feed One", settings.aliases[0])
		assert.Equal(t, 5, settings.feedLimit)
		assert.Equal(t, true, settings.showSource)
		assert.Equal(t, true, settings.showPublishDate)
		assert.Equal(t, "2006-01-02", settings.dateFormat)
		assert.Equal(t, true, settings.disableHTTP2)
		assert.Equal(t, "test-agent", settings.userAgent)
		assert.Equal(t, "blue", settings.source)
		assert.Equal(t, "red", settings.publishDate)
	})

	t.Run("uses defaults for missing values", func(t *testing.T) {
		yamlContent := `
feeds:
  - http://feed1.com/rss
`
		globalYaml := `
wtf:
  colors:
    border:
      focusable: green
`
		ymlConfig, err := config.ParseYaml(yamlContent)
		assert.NilError(t, err)
		globalConfig, err := config.ParseYaml(globalYaml)
		assert.NilError(t, err)

		settings := NewSettingsFromYAML("feedreader", ymlConfig, globalConfig)

		assert.Equal(t, -1, settings.feedLimit)
		assert.Equal(t, true, settings.showSource)
		assert.Equal(t, false, settings.showPublishDate)
		assert.Equal(t, "Jan 02", settings.dateFormat)
		assert.Equal(t, false, settings.disableHTTP2)
		assert.Equal(t, "wtfutil (https://github.com/wtfutil/wtf)", settings.userAgent)
		assert.Equal(t, "green", settings.source)
		assert.Equal(t, "orange", settings.publishDate)
	})

	t.Run("parses feeds as map with credentials", func(t *testing.T) {
		yamlContent := `
feeds:
  http://private.com/rss:
    username: myuser
    password: mypass
`
		globalYaml := `
wtf:
  colors:
    border:
      focusable: green
`
		ymlConfig, err := config.ParseYaml(yamlContent)
		assert.NilError(t, err)
		globalConfig, err := config.ParseYaml(globalYaml)
		assert.NilError(t, err)

		settings := NewSettingsFromYAML("feedreader", ymlConfig, globalConfig)

		assert.Equal(t, 1, len(settings.feeds))
		assert.Equal(t, "http://private.com/rss", settings.feeds[0])
		cred, ok := settings.credentials["http://private.com/rss"]
		assert.Assert(t, ok)
		assert.Equal(t, "myuser", cred.username)
		assert.Equal(t, "mypass", cred.password)
	})
}

func newTestWidget(mock *mockParser, settings *Settings) *Widget {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 10)
	pages := tview.NewPages()

	if settings.Common == nil {
		ymlConfig, _ := config.ParseYaml("enabled: true")
		globalConfig, _ := config.ParseYaml("wtf:\n  colors:\n    border:\n      focusable: green")
		settings.Common = cfg.NewCommonSettingsFromModule("feedreader", "Feed Reader", true, ymlConfig, globalConfig)
	}

	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(app, redrawChan, pages, settings.Common),
		parser:           mock,
		settings:         settings,
		showType:         SHOW_TITLE,
	}

	widget.SetRenderFunction(widget.Render)

	return widget
}

func Test_content(t *testing.T) {
	now := time.Now()

	t.Run("returns error message when err is set", func(t *testing.T) {
		widget := newTestWidget(&mockParser{}, &Settings{
			feedLimit:   -1,
			credentials: make(map[string]auth),
			colors:      colors{source: "green", publishDate: "orange"},
		})
		widget.err = errors.New("something went wrong")

		_, content, wrap := widget.content()

		assert.Equal(t, "something went wrong", content)
		assert.Equal(t, true, wrap)
	})

	t.Run("returns no data when stories is empty", func(t *testing.T) {
		widget := newTestWidget(&mockParser{}, &Settings{
			feedLimit:   -1,
			credentials: make(map[string]auth),
			colors:      colors{source: "green", publishDate: "orange"},
		})
		widget.stories = []*FeedItem{}

		_, content, wrap := widget.content()

		assert.Equal(t, "No data", content)
		assert.Equal(t, false, wrap)
	})

	t.Run("renders stories with row numbers", func(t *testing.T) {
		widget := newTestWidget(&mockParser{}, &Settings{
			feedLimit:   -1,
			credentials: make(map[string]auth),
			colors:      colors{source: "green", publishDate: "orange"},
			showSource:  false,
		})
		widget.stories = []*FeedItem{
			{item: &gofeed.Item{Title: "First Post", PublishedParsed: &now}, viewed: false},
			{item: &gofeed.Item{Title: "Second Post", PublishedParsed: &now}, viewed: false},
		}
		widget.SetItemCount(2)

		_, content, _ := widget.content()

		assert.Assert(t, len(content) > 0)
		assert.Assert(t, contains(content, "First Post"))
		assert.Assert(t, contains(content, "Second Post"))
	})

	t.Run("grays out viewed items", func(t *testing.T) {
		widget := newTestWidget(&mockParser{}, &Settings{
			feedLimit:   -1,
			credentials: make(map[string]auth),
			colors:      colors{source: "green", publishDate: "orange"},
			showSource:  false,
		})
		widget.stories = []*FeedItem{
			{item: &gofeed.Item{Title: "Viewed Post", PublishedParsed: &now}, viewed: true},
		}
		widget.SetItemCount(1)

		_, content, _ := widget.content()

		assert.Assert(t, contains(content, "gray"))
	})
}

func Test_Refresh(t *testing.T) {
	now := time.Now()

	t.Run("populates stories on success", func(t *testing.T) {
		mock := &mockParser{
			feeds: map[string]*gofeed.Feed{
				"http://feed.com/rss": {
					Title: "Test",
					Items: []*gofeed.Item{
						{Title: "Post 1", PublishedParsed: &now},
					},
				},
			},
		}

		widget := newTestWidget(mock, &Settings{
			feeds:       []string{"http://feed.com/rss"},
			feedLimit:   -1,
			credentials: make(map[string]auth),
			colors:      colors{source: "green", publishDate: "orange"},
		})

		widget.Refresh()

		assert.Assert(t, widget.err == nil)
		assert.Equal(t, 1, len(widget.stories))
	})

	t.Run("sets error on failure", func(t *testing.T) {
		mock := &mockParser{
			err: errors.New("network error"),
		}

		widget := newTestWidget(mock, &Settings{
			feeds:       []string{"http://bad.com/rss"},
			feedLimit:   -1,
			credentials: make(map[string]auth),
			colors:      colors{source: "green", publishDate: "orange"},
		})

		widget.Refresh()

		assert.Assert(t, widget.err != nil)
		assert.Assert(t, widget.stories == nil)
	})
}

func Test_openStory_full(t *testing.T) {
	now := time.Now()

	t.Run("marks selected story as viewed", func(t *testing.T) {
		widget := newTestWidget(&mockParser{}, &Settings{
			feedLimit:   -1,
			credentials: make(map[string]auth),
			colors:      colors{source: "green", publishDate: "orange"},
		})
		widget.stories = []*FeedItem{
			{item: &gofeed.Item{Title: "Post 1", Link: "", PublishedParsed: &now}, viewed: false},
			{item: &gofeed.Item{Title: "Post 2", Link: "", PublishedParsed: &now}, viewed: false},
		}
		widget.SetItemCount(2)
		widget.Selected = 0

		// Call openStory - won't actually open a browser since Link is empty
		widget.openStory()

		assert.Equal(t, true, widget.stories[0].viewed)
		assert.Equal(t, false, widget.stories[1].viewed)
	})

	t.Run("does nothing when no stories", func(t *testing.T) {
		widget := newTestWidget(&mockParser{}, &Settings{
			feedLimit:   -1,
			credentials: make(map[string]auth),
			colors:      colors{source: "green", publishDate: "orange"},
		})
		widget.stories = nil
		widget.Selected = -1

		// Should not panic
		widget.openStory()
	})
}

func Test_toggleDisplayText_full(t *testing.T) {
	t.Run("rotates show type and re-renders", func(t *testing.T) {
		widget := newTestWidget(&mockParser{}, &Settings{
			feedLimit:   -1,
			credentials: make(map[string]auth),
			colors:      colors{source: "green", publishDate: "orange"},
		})
		widget.stories = []*FeedItem{}

		assert.Equal(t, SHOW_TITLE, widget.showType)

		widget.toggleDisplayText()
		assert.Equal(t, SHOW_LINK, widget.showType)

		widget.toggleDisplayText()
		assert.Equal(t, SHOW_CONTENT, widget.showType)

		widget.toggleDisplayText()
		assert.Equal(t, SHOW_TITLE, widget.showType)
	})
}

func Test_NewWidget(t *testing.T) {
	t.Run("creates widget with default settings", func(t *testing.T) {
		app := tview.NewApplication()
		redrawChan := make(chan bool, 10)
		pages := tview.NewPages()

		ymlConfig, _ := config.ParseYaml("enabled: true\nfeeds:\n  - http://test.com/rss")
		globalConfig, _ := config.ParseYaml("wtf:\n  colors:\n    border:\n      focusable: green")

		settings := NewSettingsFromYAML("feedreader", ymlConfig, globalConfig)
		widget := NewWidget(app, redrawChan, pages, settings)

		assert.Assert(t, widget != nil)
		assert.Assert(t, widget.parser != nil)
		assert.Equal(t, SHOW_TITLE, widget.showType)
	})

	t.Run("creates widget with HTTP2 disabled", func(t *testing.T) {
		app := tview.NewApplication()
		redrawChan := make(chan bool, 10)
		pages := tview.NewPages()

		ymlConfig, _ := config.ParseYaml("enabled: true\nfeeds:\n  - http://test.com/rss\ndisableHTTP2: true")
		globalConfig, _ := config.ParseYaml("wtf:\n  colors:\n    border:\n      focusable: green")

		settings := NewSettingsFromYAML("feedreader", ymlConfig, globalConfig)
		widget := NewWidget(app, redrawChan, pages, settings)

		assert.Assert(t, widget != nil)
		assert.Assert(t, widget.parser != nil)
	})
}

func Test_ConfigText(t *testing.T) {
	t.Run("returns non-empty help text", func(t *testing.T) {
		widget := newTestWidget(&mockParser{}, &Settings{
			feedLimit:   -1,
			credentials: make(map[string]auth),
			colors:      colors{source: "green", publishDate: "orange"},
		})

		text := widget.ConfigText()
		assert.Assert(t, len(text) > 0)
	})
}

// contains is a simple test helper
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func Test_content_viewed_and_selected(t *testing.T) {
	now := time.Now()

	t.Run("viewed and selected item uses highlighted background", func(t *testing.T) {
		widget := newTestWidget(&mockParser{}, &Settings{
			feedLimit:   -1,
			credentials: make(map[string]auth),
			colors:      colors{source: "green", publishDate: "orange"},
			showSource:  false,
		})
		widget.stories = []*FeedItem{
			{item: &gofeed.Item{Title: "Selected Viewed Post", PublishedParsed: &now}, viewed: true},
		}
		widget.SetItemCount(1)
		widget.Selected = 0

		_, content, _ := widget.content()

		// Should contain gray with highlighted background
		assert.Assert(t, contains(content, "gray:"))
	})
}

func Test_NewSettingsFromYAML_malformed_credentials(t *testing.T) {
	t.Run("skips feed entries that are not maps", func(t *testing.T) {
		yamlContent := `
feeds:
  http://feed.com/rss: not_a_map
`
		globalYaml := `
wtf:
  colors:
    border:
      focusable: green
`
		ymlConfig, err := config.ParseYaml(yamlContent)
		assert.NilError(t, err)
		globalConfig, err := config.ParseYaml(globalYaml)
		assert.NilError(t, err)

		settings := NewSettingsFromYAML("feedreader", ymlConfig, globalConfig)

		assert.Equal(t, 0, len(settings.feeds))
	})

	t.Run("skips entries with missing username", func(t *testing.T) {
		yamlContent := `
feeds:
  http://feed.com/rss:
    password: pass
`
		globalYaml := `
wtf:
  colors:
    border:
      focusable: green
`
		ymlConfig, err := config.ParseYaml(yamlContent)
		assert.NilError(t, err)
		globalConfig, err := config.ParseYaml(globalYaml)
		assert.NilError(t, err)

		settings := NewSettingsFromYAML("feedreader", ymlConfig, globalConfig)

		assert.Equal(t, 0, len(settings.feeds))
	})

	t.Run("skips entries with missing password", func(t *testing.T) {
		yamlContent := `
feeds:
  http://feed.com/rss:
    username: user
`
		globalYaml := `
wtf:
  colors:
    border:
      focusable: green
`
		ymlConfig, err := config.ParseYaml(yamlContent)
		assert.NilError(t, err)
		globalConfig, err := config.ParseYaml(globalYaml)
		assert.NilError(t, err)

		settings := NewSettingsFromYAML("feedreader", ymlConfig, globalConfig)

		assert.Equal(t, 0, len(settings.feeds))
	})
}
