package mercurial

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
)

func testSettings() *Settings {
	return &Settings{
		Common: &cfg.Common{
			Colors: cfg.ColorTheme{
				TextTheme: cfg.TextTheme{
					Subheading: "red",
					Title:      "green",
				},
			},
			Title: "Mercurial",
		},
	}
}

func Test_formatChange(t *testing.T) {
	tests := []struct {
		name string
		line string
		want string
	}{
		{name: "empty line", line: "", want: ""},
		{name: "added file", line: "A foo.txt", want: " [green]A[white] foo.txt\n"},
		{name: "deleted file", line: "D bar.txt", want: " [red]D[white] bar.txt\n"},
		{name: "modified file", line: "M baz.txt", want: " [yellow]M[white] baz.txt\n"},
		{name: "renamed file", line: "R qux.txt", want: " [purple]R[white] qux.txt\n"},
		{name: "untouched marker", line: "? unknown.txt", want: " ? unknown.txt\n"},
		{name: "leading/trailing whitespace trimmed", line: "  M spaced.txt  ", want: " [yellow]M[white] spaced.txt\n"},
		{name: "quotes stripped", line: `M "quoted file.txt"`, want: " [yellow]M[white] quoted file.txt\n"},
	}

	widget := &Widget{settings: testSettings()}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, widget.formatChange(tt.line))
		})
	}
}

func Test_formatChanges(t *testing.T) {
	widget := &Widget{settings: testSettings()}

	tests := []struct {
		name string
		data []string
		want string
	}{
		{
			name: "no changes",
			data: []string{""},
			want: " [red]Changed Files[white]\n [grey]none[white]\n",
		},
		{
			name: "single change",
			data: []string{"M foo.txt", ""},
			want: " [red]Changed Files[white]\n [yellow]M[white] foo.txt\n",
		},
		{
			name: "multiple changes",
			data: []string{"A foo.txt", "D bar.txt", ""},
			want: " [red]Changed Files[white]\n [green]A[white] foo.txt\n [red]D[white] bar.txt\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, widget.formatChanges(tt.data))
		})
	}
}

func Test_formatCommit(t *testing.T) {
	widget := &Widget{settings: testSettings()}

	tests := []struct {
		name string
		line string
		want string
	}{
		{name: "plain commit", line: "1:abcdef added feature", want: " 1:abcdef added feature\n"},
		{name: "commit with quotes stripped", line: `2:123456 fixed "bug"`, want: " 2:123456 fixed bug\n"},
		{name: "empty commit line", line: "", want: " \n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, widget.formatCommit(tt.line))
		})
	}
}

func Test_formatCommits(t *testing.T) {
	widget := &Widget{settings: testSettings()}

	tests := []struct {
		name string
		data []string
		want string
	}{
		{
			name: "no commits",
			data: []string{},
			want: " [red]Recent Commits[white]\n",
		},
		{
			name: "single commit",
			data: []string{"1:abc first commit"},
			want: " [red]Recent Commits[white]\n 1:abc first commit\n",
		},
		{
			name: "multiple commits",
			data: []string{"1:abc first", "2:def second"},
			want: " [red]Recent Commits[white]\n 1:abc first\n 2:def second\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, widget.formatCommits(tt.data))
		})
	}
}

func Test_content_noData(t *testing.T) {
	settings := testSettings()
	widget := &Widget{
		TextWidget: newTestTextWidget(settings),
		settings:   settings,
	}

	title, description, wrap := widget.content()

	assert.Equal(t, "Mercurial", title)
	assert.Equal(t, " Mercurial repo data is unavailable ", description)
	assert.False(t, wrap)
}

func Test_content_withData(t *testing.T) {
	settings := testSettings()
	widget := &Widget{
		TextWidget: newTestTextWidget(settings),
		settings:   settings,
		Data: []*MercurialRepo{
			{
				Branch:       "default",
				Bookmark:     "my-bookmark",
				Repository:   "/some/repo",
				ChangedFiles: []string{"M foo.txt", ""},
				Commits:      []string{"1:abc first commit"},
			},
		},
	}

	title, description, wrap := widget.content()

	assert.Contains(t, title, "/some/repo")
	assert.Contains(t, description, "default:my-bookmark")
	assert.Contains(t, description, "foo.txt")
	assert.Contains(t, description, "first commit")
	assert.False(t, wrap)
}

// newTestTextWidget creates a minimal view.TextWidget suitable for exercising
// content() without needing a running tview application.
func newTestTextWidget(settings *Settings) view.TextWidget {
	return view.NewTextWidget(tview.NewApplication(), make(chan bool), tview.NewPages(), settings.Common)
}
