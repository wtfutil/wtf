package githuballrepos

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GitHubFetcher interface {
	FetchData(orgs []string, username string) *WidgetData
}

// GitHubClient handles communication with the GitHub API
type GitHubClient struct {
	client *githubv4.Client
}

// NewGitHubClient creates a new GitHubClient
func NewGitHubClient(token string) *GitHubClient {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	return &GitHubClient{
		client: githubv4.NewClient(httpClient),
	}
}

// FetchData retrieves all required data from GitHub
func (c *GitHubClient) FetchData(orgs []string, username string) *WidgetData {
	data := &WidgetData{
		MyPRs:            make([]PR, 0),
		PRReviewRequests: make([]PR, 0),
		WatchedPRs:       make([]PR, 0),
	}

	for _, org := range orgs {
		data.PRsOpenedByMe += c.fetchPRCount(org, username, "author")
		data.PRReviewRequestsCount += c.fetchPRCount(org, username, "review-requested")
		data.OpenIssuesCount += c.fetchIssueCount(org)

		data.MyPRs = append(data.MyPRs, c.fetchPRs(org, username, "author")...)
		data.PRReviewRequests = append(data.PRReviewRequests, c.fetchPRs(org, username, "review-requested")...)
		data.WatchedPRs = append(data.WatchedPRs, c.fetchPRs(org, username, "involves")...)
	}

	return data
}

func (c *GitHubClient) fetchPRCount(org, username, filter string) int {
	var query struct {
		Search struct {
			IssueCount int
		} `graphql:"search(query: $query, type: ISSUE, first: 0)"`
	}

	variables := map[string]interface{}{
		"query": githubv4.String(fmt.Sprintf("org:%s is:pr is:open %s:%s", org, filter, username)),
	}

	err := c.client.Query(context.Background(), &query, variables)
	if err != nil {
		// Handle error (log it, etc.)
		return 0
	}

	return query.Search.IssueCount
}

func (c *GitHubClient) fetchIssueCount(org string) int {
	var query struct {
		Search struct {
			IssueCount int
		} `graphql:"search(query: $query, type: ISSUE, first: 0)"`
	}

	variables := map[string]interface{}{
		"query": githubv4.String(fmt.Sprintf("org:%s is:issue is:open", org)),
	}

	err := c.client.Query(context.Background(), &query, variables)
	if err != nil {
		// Handle error (log it, etc.)
		return 0
	}

	return query.Search.IssueCount
}

func (c *GitHubClient) fetchPRs(org, username, filter string) []PR {
	var query struct {
		Search struct {
			Nodes []struct {
				PullRequest struct {
					Title  string
					URL    string
					Author struct {
						Login string
					}
					Repository struct {
						Name string
					}
				} `graphql:"... on PullRequest"`
			}
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"search(query: $query, type: ISSUE, first: 100, after: $cursor)"`
	}

	variables := map[string]interface{}{
		"query":  githubv4.String(fmt.Sprintf("org:%s is:pr is:open %s:%s", org, filter, username)),
		"cursor": (*githubv4.String)(nil), // Null for first request
	}

	var allPRs []PR

	for {
		err := c.client.Query(context.Background(), &query, variables)
		if err != nil {
			// Handle error (log it, etc.)
			break
		}

		for _, node := range query.Search.Nodes {
			pr := node.PullRequest
			allPRs = append(allPRs, PR{
				Title:      pr.Title,
				URL:        pr.URL,
				Author:     pr.Author.Login,
				Repository: pr.Repository.Name,
			})
		}

		if !query.Search.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Search.PageInfo.EndCursor)
	}

	return allPRs
}
