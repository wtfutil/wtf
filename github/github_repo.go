package github

import ()

type GithubRepo struct {
	Name  string
	Owner string
}

func NewGithubRepo(owner string, name string) *GithubRepo {
	repo := GithubRepo{
		Name:  name,
		Owner: owner,
	}

	return &repo
}
