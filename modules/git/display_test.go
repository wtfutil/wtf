package git

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
)

func newTestSettings() *Settings {
	return &Settings{
		Common: &cfg.Common{
			Colors: cfg.ColorTheme{TextTheme: cfg.TextTheme{Subheading: "red"}},
		},
	}
}

func Test_formatChange(t *testing.T) {
	widget := &Widget{settings: newTestSettings()}

	tests := []struct {
		name string
		line string
		want string
	}{
		{"empty line returns empty string", "", ""},
		{"added file is colorized green", ` A "foo.txt"`, " [green]A[white] foo.txt\n"},
		{"deleted file is colorized red", ` D "foo.txt"`, " [red]D[white] foo.txt\n"},
		{"modified file is colorized yellow", ` M "foo.txt"`, " [yellow]M[white] foo.txt\n"},
		{"renamed file is colorized purple", ` R "foo.txt" -> "bar.txt"`, " [purple]R[white] foo.txt -> bar.txt\n"},
		{"unrecognized status is left unmodified", ` ? foo.txt`, " ? foo.txt\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, widget.formatChange(tt.line))
		})
	}
}

func Test_formatChanges(t *testing.T) {
	widget := &Widget{settings: newTestSettings()}

	tests := []struct {
		name string
		data []string
		want string
	}{
		{
			name: "single empty entry renders 'none'",
			data: []string{""},
			want: " [red]Changed Files[white]\n [grey]none[white]\n",
		},
		{
			name: "multiple entries are each formatted",
			data: []string{` M "foo.txt"`, ` A "bar.txt"`, ""},
			want: " [red]Changed Files[white]\n [yellow]M[white] foo.txt\n [green]A[white] bar.txt\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, widget.formatChanges(tt.data))
		})
	}
}

func Test_formatCommit(t *testing.T) {
	widget := &Widget{settings: newTestSettings()}

	assert.Equal(t, " abc123 fixed the thing\n", widget.formatCommit(`abc123 "fixed the thing"`))
	assert.Equal(t, " \n", widget.formatCommit(""))
}

func Test_formatCommits(t *testing.T) {
	widget := &Widget{settings: newTestSettings()}

	data := []string{`abc123 "first"`, `def456 "second"`}
	want := " [red]Recent Commits[white]\n abc123 first\n def456 second\n"

	assert.Equal(t, want, widget.formatCommits(data))
}

func Test_content(t *testing.T) {
	tests := []struct {
		name          string
		settings      *Settings
		repos         []*GitRepo
		idx           int
		wantTitle     string
		wantUnavail   bool
		wantBodyEmpty bool
	}{
		{
			name:        "no repo data is unavailable",
			settings:    newTestSettings(),
			repos:       nil,
			wantTitle:   "",
			wantUnavail: true,
		},
		{
			name:     "index out of range is unavailable",
			settings: newTestSettings(),
			repos: []*GitRepo{
				{Repository: "/home/user/project", Branch: "main"},
			},
			idx:         5,
			wantUnavail: true,
		},
		{
			name: "full repository path shown when lastFolderTitle disabled",
			settings: &Settings{
				Common:   newTestSettings().Common,
				sections: []interface{}{},
			},
			repos: []*GitRepo{
				{Repository: "/home/user/project", Branch: "main", ChangedFiles: []string{""}, Commits: []string{}},
			},
			wantTitle: "/home/user/project[white]",
		},
		{
			name: "last folder shown when lastFolderTitle enabled",
			settings: &Settings{
				Common:          newTestSettings().Common,
				sections:        []interface{}{},
				lastFolderTitle: true,
			},
			repos: []*GitRepo{
				{Repository: "/home/user/project", Branch: "main", ChangedFiles: []string{""}, Commits: []string{}},
			},
			wantTitle: "project[white]",
		},
		{
			name: "branch is appended to title when branchInTitle enabled",
			settings: &Settings{
				Common:        newTestSettings().Common,
				sections:      []interface{}{},
				branchInTitle: true,
			},
			repos: []*GitRepo{
				{Repository: "/home/user/project", Branch: "main", ChangedFiles: []string{""}, Commits: []string{}},
			},
			wantTitle: "/home/user/project <main>[white]",
		},
		{
			name: "module name is prefixed when showModuleName enabled",
			settings: &Settings{
				Common: &cfg.Common{
					Colors: cfg.ColorTheme{TextTheme: cfg.TextTheme{Subheading: "red"}},
					Title:  "Git",
				},
				sections:       []interface{}{},
				showModuleName: true,
			},
			repos: []*GitRepo{
				{Repository: "/home/user/project", Branch: "main", ChangedFiles: []string{""}, Commits: []string{}},
			},
			wantTitle: "Git - /home/user/project[white]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{
				settings: tt.settings,
				GitRepos: tt.repos,
			}
			widget.Idx = tt.idx
			widget.TextWidget = view.TextWidget{
				Base: view.NewBase(nil, nil, nil, tt.settings.Common),
				View: tview.NewTextView(),
			}

			title, body, wrap := widget.content()

			assert.Equal(t, tt.wantTitle, title)
			assert.False(t, wrap)

			if tt.wantUnavail {
				assert.Equal(t, " Git repo data is unavailable ", body)
			} else {
				assert.NotEmpty(t, body)
			}
		})
	}
}
