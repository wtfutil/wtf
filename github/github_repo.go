package github

import (
	"context"
	"net/http"
	"os"

	ghb "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubRepo struct {
	apiKey string

	Name         string
	Owner        string
	PullRequests []*ghb.PullRequest
}

func NewGithubRepo(name, owner string) *GithubRepo {
	repo := GithubRepo{
		apiKey: os.Getenv("WTF_GITHUB_TOKEN"),
		Name:   name,
		Owner:  owner,
	}

	repo.PullRequests, _ = repo.pullRequests()

	return &repo
}

func (repo *GithubRepo) pullRequests() ([]*ghb.PullRequest, error) {
	oauthClient := repo.oauthClient()
	github := ghb.NewClient(oauthClient)

	opts := &ghb.PullRequestListOptions{}

	prs, _, err := github.PullRequests.List(context.Background(), repo.Owner, repo.Name, opts)

	if err != nil {
		return nil, err
	}

	return prs, nil
}

func (repo *GithubRepo) Repository() (*ghb.Repository, error) {
	oauthClient := repo.oauthClient()
	github := ghb.NewClient(oauthClient)

	repository, _, err := github.Repositories.Get(context.Background(), repo.Owner, repo.Name)

	if err != nil {
		return nil, err
	}

	return repository, nil
}

/* -------------------- Unexported Functions -------------------- */

func (repo *GithubRepo) oauthClient() *http.Client {
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: repo.apiKey},
	)

	return oauth2.NewClient(context.Background(), tokenService)
}
