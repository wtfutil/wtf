package github

import (
	"net/http"
	"os"

	ghb "github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type GithubRepo struct {
	apiKey string

	Name         string
	Owner        string
	PullRequests []*ghb.PullRequest
	RemoteRepo   *ghb.Repository
}

func NewGithubRepo(name, owner string) *GithubRepo {
	repo := GithubRepo{
		apiKey: os.Getenv("WTF_GITHUB_TOKEN"),
		Name:   name,
		Owner:  owner,
	}

	return &repo
}

// Refresh reloads the github data via the Github API
func (repo *GithubRepo) Refresh() {
	repo.PullRequests, _ = repo.loadPullRequests()
	repo.RemoteRepo, _ = repo.loadRemoteRepository()
}

/* -------------------- Counts -------------------- */

func (repo *GithubRepo) IssueCount() int {
	return *repo.RemoteRepo.OpenIssuesCount
}

func (repo *GithubRepo) PullRequestCount() int {
	return len(repo.PullRequests)
}

func (repo *GithubRepo) StarCount() int {
	return *repo.RemoteRepo.StargazersCount
}

/* -------------------- Unexported Functions -------------------- */

func (repo *GithubRepo) oauthClient() *http.Client {
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: repo.apiKey},
	)

	return oauth2.NewClient(context.Background(), tokenService)
}

// myPullRequests returns a list of pull requests created by username on this repo
func (repo *GithubRepo) myPullRequests(username string) []*ghb.PullRequest {
	prs := []*ghb.PullRequest{}

	for _, pr := range repo.PullRequests {
		user := *pr.User

		if *user.Login == username {
			prs = append(prs, pr)
		}
	}

	return prs
}

// myReviewRequests returns a list of pull requests for which username has been
// requested to do a code review
func (repo *GithubRepo) myReviewRequests(username string) []*ghb.PullRequest {
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

func (repo *GithubRepo) loadPullRequests() ([]*ghb.PullRequest, error) {
	oauthClient := repo.oauthClient()
	github := ghb.NewClient(oauthClient)

	opts := &ghb.PullRequestListOptions{}

	prs, _, err := github.PullRequests.List(context.Background(), repo.Owner, repo.Name, opts)

	if err != nil {
		return nil, err
	}

	return prs, nil
}

func (repo *GithubRepo) loadRemoteRepository() (*ghb.Repository, error) {
	oauthClient := repo.oauthClient()
	github := ghb.NewClient(oauthClient)

	repository, _, err := github.Repositories.Get(context.Background(), repo.Owner, repo.Name)

	if err != nil {
		return nil, err
	}

	return repository, nil
}
