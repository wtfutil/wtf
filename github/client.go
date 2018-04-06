package github

import (
	"context"
	//"fmt"
	"net/http"
	"os"

	ghb "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Client struct {
	apiKey string
}

func NewClient() *Client {
	client := Client{
		apiKey: os.Getenv("WTF_GITHUB_TOKEN"),
	}

	return &client
}

func (client *Client) PullRequests(orgName string, repoName string) ([]*ghb.PullRequest, error) {
	oauthClient := client.oauthClient()
	github := ghb.NewClient(oauthClient)

	opts := &ghb.PullRequestListOptions{}

	prs, _, err := github.PullRequests.List(context.Background(), orgName, repoName, opts)

	if err != nil {
		return nil, err
		//fmt.Printf("Problem in getting pull request information %v\n", err)
		//os.Exit(1)
	}

	return prs, nil
}

func (client *Client) Repository(orgName string, repoName string) (*ghb.Repository, error) {
	oauthClient := client.oauthClient()
	github := ghb.NewClient(oauthClient)

	repo, _, err := github.Repositories.Get(context.Background(), orgName, repoName)

	if err != nil {
		return nil, err
		//fmt.Printf("Problem in getting repository information %v\n", err)
		//os.Exit(1)
	}

	return repo, nil
}

/* -------------------- Unexported Functions -------------------- */

func (client *Client) oauthClient() *http.Client {
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: client.apiKey},
	)

	return oauth2.NewClient(context.Background(), tokenService)
}
