package github

import (
	"context"
	"fmt"
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

	repo.PullRequests, _ = repo.allPullRequests()

	return &repo
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

// TODO: This should return a slice of pull requests and let Display handle the output
func (repo *GithubRepo) myPullRequests(username string) string {
	if len(repo.PullRequests) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""

	for _, pr := range repo.PullRequests {
		user := *pr.User

		if *user.Login == username {
			str = str + fmt.Sprintf(" [green]%4d[white] %s\n", *pr.Number, *pr.Title)
		}
	}

	if str == "" {
		return " [grey]none[white]\n"
	}

	return str
}

// TODO: This should return a slice of pull requests and let Display handle the output
func (repo *GithubRepo) pullRequetsForMeToReview(username string) string {
	if len(repo.PullRequests) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""

	for _, pr := range repo.PullRequests {
		for _, reviewer := range pr.RequestedReviewers {
			if *reviewer.Login == username {
				str = str + fmt.Sprintf(" [green]%d[white] %s\n", *pr.Number, *pr.Title)
			}
		}
	}

	if str == "" {
		return " [grey]none[white]\n"
	}

	return str
}

func (repo *GithubRepo) allPullRequests() ([]*ghb.PullRequest, error) {
	oauthClient := repo.oauthClient()
	github := ghb.NewClient(oauthClient)

	opts := &ghb.PullRequestListOptions{}

	prs, _, err := github.PullRequests.List(context.Background(), repo.Owner, repo.Name, opts)

	if err != nil {
		return nil, err
	}

	return prs, nil
}
