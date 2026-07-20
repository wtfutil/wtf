package github

import (
	"context"
	"fmt"
	"time"

	ghb "github.com/google/go-github/v89/github"
	"github.com/wtfutil/wtf/utils"
)

const apiTimeout = 30 * time.Second

const (
	pullRequestsPath = "/pulls"
	issuesPath       = "/issues"
)

// Repo defines a new GitHub Repo structure
type Repo struct {
	apiKey    string
	baseURL   string
	uploadURL string

	Name         string
	Owner        string
	PullRequests []*ghb.PullRequest
	RemoteRepo   *ghb.Repository
	Err          error

	// Cached display data populated during Refresh
	CachedMyPullRequests   []*ghb.PullRequest
	CachedMyReviewRequests []*ghb.PullRequest
	CachedCustomQueries    map[string]*ghb.IssuesSearchResult
}

// NewGithubRepo returns a new Github Repo with a name, owner, apiKey, baseURL and uploadURL
func NewGithubRepo(name, owner, apiKey, baseURL, uploadURL string) *Repo {
	repo := Repo{
		Name:  name,
		Owner: owner,

		apiKey:    apiKey,
		baseURL:   baseURL,
		uploadURL: uploadURL,
	}

	return &repo
}

// Open will open the GitHub Repo URL using the utils helper
func (repo *Repo) Open() {
	utils.OpenFile(*repo.RemoteRepo.HTMLURL)
}

// OpenPulls will open the GitHub Pull Requests URL using the utils helper
func (repo *Repo) OpenPulls() {
	utils.OpenFile(*repo.RemoteRepo.HTMLURL + pullRequestsPath)
}

// OpenIssues will open the GitHub Issues URL using the utils helper
func (repo *Repo) OpenIssues() {
	utils.OpenFile(*repo.RemoteRepo.HTMLURL + issuesPath)
}

// Refresh reloads the github data via the Github API and caches display data
func (repo *Repo) Refresh(username string, enableStatus bool, customQueries []customQuery) {
	ctx, cancel := context.WithTimeout(context.Background(), apiTimeout)
	defer cancel()

	prs, err := repo.loadPullRequests(ctx)
	repo.Err = err
	repo.PullRequests = prs
	if err != nil {
		return
	}
	remote, err := repo.loadRemoteRepository(ctx)
	repo.Err = err
	repo.RemoteRepo = remote
	if err != nil {
		return
	}

	// Pre-compute display data so content() never does HTTP
	repo.CachedMyReviewRequests = repo.computeMyReviewRequests(username)
	repo.CachedMyPullRequests = repo.computeMyPullRequests(ctx, username, enableStatus)

	repo.CachedCustomQueries = make(map[string]*ghb.IssuesSearchResult)
	for _, q := range customQueries {
		repo.CachedCustomQueries[q.filter] = repo.fetchCustomIssueQuery(ctx, q.filter, q.perPage)
	}
}

/* -------------------- Counts -------------------- */

// IssueCount return the total amount of issues as an int
func (repo *Repo) IssueCount() int {
	if repo.RemoteRepo == nil {
		return 0
	}

	issuesLessPulls := *repo.RemoteRepo.OpenIssuesCount - len(repo.PullRequests)

	return issuesLessPulls
}

// PullRequestCount returns the total amount of pull requests as an int
func (repo *Repo) PullRequestCount() int {
	return len(repo.PullRequests)
}

// StarCount returns the total amount of stars this repo has gained as an int
func (repo *Repo) StarCount() int {
	if repo.RemoteRepo == nil {
		return 0
	}

	return *repo.RemoteRepo.StargazersCount
}

/* -------------------- Unexported Functions -------------------- */

func (repo *Repo) isGitHubEnterprise() bool {
	if len(repo.baseURL) > 0 {
		if repo.uploadURL == "" {
			repo.uploadURL = repo.baseURL
		}
		return true
	}
	return false
}

func (repo *Repo) githubClient() (*ghb.Client, error) {
	opts := []ghb.ClientOptionsFunc{ghb.WithAuthToken(repo.apiKey)}

	if repo.isGitHubEnterprise() {
		opts = append(opts, ghb.WithEnterpriseURLs(repo.baseURL, repo.uploadURL))
	}

	return ghb.NewClient(opts...)
}

// myPullRequests returns a list of pull requests created by username on this repo
func (repo *Repo) computeMyPullRequests(ctx context.Context, username string, showStatus bool) []*ghb.PullRequest {
	prs := []*ghb.PullRequest{}

	for _, pr := range repo.PullRequests {
		user := *pr.User

		if *user.Login == username {
			prs = append(prs, pr)
		}
	}

	if showStatus {
		prs = repo.individualPRs(ctx, prs)
	}

	return prs
}

// individualPRs takes a list of pull requests (presumably returned from
// github.PullRequests.List) and fetches them individually to get more detailed
// status info on each. see: https://developer.github.com/v3/git/#checking-mergeability-of-pull-requests
func (repo *Repo) individualPRs(ctx context.Context, prs []*ghb.PullRequest) []*ghb.PullRequest {
	github, err := repo.githubClient()
	if err != nil {
		return prs
	}

	var ret []*ghb.PullRequest
	for i := range prs {
		pr, _, err := github.PullRequests.Get(ctx, repo.Owner, repo.Name, prs[i].GetNumber())
		if err != nil {
			// worst case, just keep the original one
			ret = append(ret, prs[i])
		} else {
			ret = append(ret, pr)
		}
	}
	return ret
}

// myReviewRequests returns a list of pull requests for which username has been
// requested to do a code review
func (repo *Repo) computeMyReviewRequests(username string) []*ghb.PullRequest {
	prs := []*ghb.PullRequest{}

	for _, pr := range repo.PullRequests {
		for _, reviewer := range pr.RequestedReviewers {
			if *reviewer.Login == username {
				prs = append(prs, pr)
			}
		}
	}

	return prs
}

func (repo *Repo) fetchCustomIssueQuery(ctx context.Context, filter string, perPage int) *ghb.IssuesSearchResult {
	github, err := repo.githubClient()
	if err != nil {
		return nil
	}

	opts := &ghb.SearchOptions{}
	if perPage != 0 {
		opts.PerPage = perPage
	}

	prs, _, _ := github.Search.Issues(ctx, fmt.Sprintf("%s repo:%s/%s", filter, repo.Owner, repo.Name), opts)
	return prs
}

func (repo *Repo) loadPullRequests(ctx context.Context) ([]*ghb.PullRequest, error) {
	github, err := repo.githubClient()
	if err != nil {
		return nil, err
	}

	opts := &ghb.PullRequestListOptions{}
	opts.PerPage = 100

	prs, _, err := github.PullRequests.List(ctx, repo.Owner, repo.Name, opts)

	if err != nil {
		return nil, err
	}

	return prs, nil
}

func (repo *Repo) loadRemoteRepository(ctx context.Context) (*ghb.Repository, error) {
	github, err := repo.githubClient()

	if err != nil {
		return nil, err
	}

	repository, _, err := github.Repositories.Get(ctx, repo.Owner, repo.Name)

	if err != nil {
		return nil, err
	}

	return repository, nil
}
