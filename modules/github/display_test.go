package github

import (
	"strings"
	"testing"

	ghb "github.com/google/go-github/v89/github"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
)

func testSettings() *Settings {
	return &Settings{
		Common: &cfg.Common{
			Colors: cfg.ColorTheme{
				TextTheme: cfg.TextTheme{
					Subheading: "white",
					Title:      "green",
				},
			},
			Title: "GitHub",
		},
		enableStatus:           false,
		showMyPullRequests:     true,
		showOpenReviewRequests: true,
		showStats:              true,
	}
}

func newTestWidget(settings *Settings) *Widget {
	return &Widget{
		TextWidget: view.NewTextWidget(tview.NewApplication(), make(chan bool), tview.NewPages(), settings.Common),
		settings:   settings,
	}
}

func Test_displayMyPullRequests_escapesTitles(t *testing.T) {
	settings := testSettings()
	widget := newTestWidget(settings)
	widget.View.SetDynamicColors(true)
	widget.View.SetRegions(true)

	t.Run("empty pull requests returns none", func(t *testing.T) {
		repo := &Repo{
			CachedMyPullRequests: []*ghb.PullRequest{},
		}
		output := widget.displayMyPullRequests(repo)
		assert.Equal(t, " [grey]none[white]\n", output)
	})

	t.Run("escapes single-word and multi-word bracketed titles", func(t *testing.T) {
		repo := &Repo{
			CachedMyPullRequests: []*ghb.PullRequest{
				{
					Number: ghb.Ptr(101),
					Title:  ghb.Ptr("[Fix] Correct title display"),
				},
				{
					Number: ghb.Ptr(102),
					Title:  ghb.Ptr("[CI Perf] Optimize test pipeline"),
				},
				{
					Number: ghb.Ptr(103),
					Title:  ghb.Ptr("[Bug][Core] Fix data race"),
				},
				{
					Number: ghb.Ptr(104),
					Title:  ghb.Ptr("Standard PR without brackets"),
				},
			},
		}

		output := widget.displayMyPullRequests(repo)
		require.NotEmpty(t, output)

		// Directly check that the raw formatted string contains the tview-escaped titles
		assert.Contains(t, output, tview.Escape("[Fix] Correct title display"))
		assert.Contains(t, output, tview.Escape("[CI Perf] Optimize test pipeline"))
		assert.Contains(t, output, tview.Escape("[Bug][Core] Fix data race"))
		assert.Contains(t, output, tview.Escape("Standard PR without brackets"))

		// Render through tview TextView and verify bracketed titles are not stripped
		widget.View.SetText(output)
		rendered := widget.View.GetText(true)

		assert.Contains(t, rendered, "[Fix] Correct title display")
		assert.Contains(t, rendered, "[CI Perf] Optimize test pipeline")
		assert.Contains(t, rendered, "[Bug][Core] Fix data race")
		assert.Contains(t, rendered, "Standard PR without brackets")
	})
}

func Test_displayCustomQuery_escapesTitles(t *testing.T) {
	settings := testSettings()
	widget := newTestWidget(settings)
	widget.View.SetDynamicColors(true)
	widget.View.SetRegions(true)

	t.Run("nil query result returns invalid query", func(t *testing.T) {
		repo := &Repo{
			CachedCustomQueries: make(map[string]*ghb.IssuesSearchResult),
		}
		output := widget.displayCustomQuery(repo, "nonexistent")
		assert.Equal(t, " [grey]Invalid Query[white]\n", output)
	})

	t.Run("empty issues returns none", func(t *testing.T) {
		repo := &Repo{
			CachedCustomQueries: map[string]*ghb.IssuesSearchResult{
				"is:open": {
					Issues: []*ghb.Issue{},
				},
			},
		}
		output := widget.displayCustomQuery(repo, "is:open")
		assert.Equal(t, " [grey]none[white]\n", output)
	})

	t.Run("escapes single-word and bracketed issue titles", func(t *testing.T) {
		repo := &Repo{
			CachedCustomQueries: map[string]*ghb.IssuesSearchResult{
				"is:issue is:open": {
					Issues: []*ghb.Issue{
						{
							Number: ghb.Ptr(201),
							Title:  ghb.Ptr("[Bug] Crash on launch"),
						},
						{
							Number: ghb.Ptr(202),
							Title:  ghb.Ptr("[Feature Request] Add dark theme"),
						},
						{
							Number: ghb.Ptr(203),
							Title:  ghb.Ptr("Standard issue description"),
						},
					},
				},
			},
		}

		output := widget.displayCustomQuery(repo, "is:issue is:open")
		require.NotEmpty(t, output)

		assert.Contains(t, output, tview.Escape("[Bug] Crash on launch"))
		assert.Contains(t, output, tview.Escape("[Feature Request] Add dark theme"))
		assert.Contains(t, output, tview.Escape("Standard issue description"))

		widget.View.SetText(output)
		rendered := widget.View.GetText(true)

		assert.Contains(t, rendered, "[Bug] Crash on launch")
		assert.Contains(t, rendered, "[Feature Request] Add dark theme")
		assert.Contains(t, rendered, "Standard issue description")
	})
}

func Test_displayMyReviewRequests_escapesTitles(t *testing.T) {
	settings := testSettings()
	widget := newTestWidget(settings)
	widget.View.SetDynamicColors(true)
	widget.View.SetRegions(true)

	t.Run("empty review requests returns none", func(t *testing.T) {
		repo := &Repo{
			CachedMyReviewRequests: []*ghb.PullRequest{},
		}
		output := widget.displayMyReviewRequests(repo)
		assert.Equal(t, " [grey]none[white]\n", output)
	})

	t.Run("escapes bracketed review request titles", func(t *testing.T) {
		repo := &Repo{
			CachedMyReviewRequests: []*ghb.PullRequest{
				{
					Number: ghb.Ptr(301),
					Title:  ghb.Ptr("[WIP] Experimental refactor"),
				},
				{
					Number: ghb.Ptr(302),
					Title:  ghb.Ptr("[Docs] Update contributing guide"),
				},
				{
					Number: ghb.Ptr(303),
					Title:  ghb.Ptr("Refactor logger"),
				},
			},
		}

		output := widget.displayMyReviewRequests(repo)
		require.NotEmpty(t, output)

		assert.Contains(t, output, tview.Escape("[WIP] Experimental refactor"))
		assert.Contains(t, output, tview.Escape("[Docs] Update contributing guide"))
		assert.Contains(t, output, tview.Escape("Refactor logger"))

		widget.View.SetText(output)
		rendered := widget.View.GetText(true)

		assert.Contains(t, rendered, "[WIP] Experimental refactor")
		assert.Contains(t, rendered, "[Docs] Update contributing guide")
		assert.Contains(t, rendered, "Refactor logger")
	})
}

func Test_mergeString(t *testing.T) {
	t.Run("returns empty string when enableStatus is false", func(t *testing.T) {
		settings := testSettings()
		settings.enableStatus = false
		widget := newTestWidget(settings)

		pr := &ghb.PullRequest{MergeableState: ghb.Ptr("clean")}
		assert.Equal(t, "", widget.mergeString(pr))
	})

	t.Run("returns merge state icon when enableStatus is true", func(t *testing.T) {
		settings := testSettings()
		settings.enableStatus = true
		widget := newTestWidget(settings)

		prClean := &ghb.PullRequest{MergeableState: ghb.Ptr("clean")}
		assert.Equal(t, "[green]\u2713[white] ", widget.mergeString(prClean))

		prDirty := &ghb.PullRequest{MergeableState: ghb.Ptr("dirty")}
		assert.Equal(t, "[red]\u0021[white] ", widget.mergeString(prDirty))

		prUnknown := &ghb.PullRequest{MergeableState: ghb.Ptr("unknown_state")}
		assert.Equal(t, "? ", widget.mergeString(prUnknown))
	})
}

func Test_content_rendersEscapedTitles(t *testing.T) {
	settings := testSettings()
	settings.customQueries = []customQuery{
		{
			title:  "Open Issues",
			filter: "is:open is:issue",
		},
	}

	widget := newTestWidget(settings)
	widget.View.SetDynamicColors(true)
	widget.View.SetRegions(true)

	repo := &Repo{
		Name:  "wtf",
		Owner: "wtfutil",
		RemoteRepo: &ghb.Repository{
			OpenIssuesCount: ghb.Ptr(5),
			StargazersCount: ghb.Ptr(17000),
		},
		PullRequests: []*ghb.PullRequest{
			{Number: ghb.Ptr(10)},
		},
		CachedMyPullRequests: []*ghb.PullRequest{
			{
				Number: ghb.Ptr(100),
				Title:  ghb.Ptr("[Fix] Title formatting bug"),
			},
		},
		CachedMyReviewRequests: []*ghb.PullRequest{
			{
				Number: ghb.Ptr(200),
				Title:  ghb.Ptr("[WIP] Dashboard enhancement"),
			},
		},
		CachedCustomQueries: map[string]*ghb.IssuesSearchResult{
			"is:open is:issue": {
				Issues: []*ghb.Issue{
					{
						Number: ghb.Ptr(300),
						Title:  ghb.Ptr("[Bug] Tag stripping on terminal"),
					},
				},
			},
		},
	}

	widget.GithubRepos = []*Repo{repo}
	widget.Idx = 0

	title, description, wrap := widget.content()

	assert.Contains(t, title, "wtfutil - wtf")
	assert.False(t, wrap)

	widget.View.SetText(description)
	rendered := widget.View.GetText(true)

	assert.True(t, strings.Contains(rendered, "[Fix] Title formatting bug"))
	assert.True(t, strings.Contains(rendered, "[WIP] Dashboard enhancement"))
	assert.True(t, strings.Contains(rendered, "[Bug] Tag stripping on terminal"))
}
