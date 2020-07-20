package support

import (
	"context"
	"errors"
	"net/http"

	ghb "github.com/google/go-github/v32/github"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"
)

var sponsorQuery struct {
	User struct {
		SponsorshipsAsSponsor struct {
			Nodes []struct {
				Sponsorable struct {
					SponsorsListing struct {
						Slug string
					}
				}
			}
		} `graphql:"sponsorshipsAsSponsor(first: 10)"`
	} `graphql:"user(login: $loginName)"`
}

// GitHubUser represents a GitHub user account as defined by a GitHub API access key
// This is used to determine whether or not the WTF user is a sponsor (via GitHub sponsors)
// and/or a contributor to WTF
type GitHubUser struct {
	apiKey string

	loginName string

	clientV3 *ghb.Client
	clientV4 *githubv4.Client

	IsContributor bool
	IsSponsor     bool
}

// NewGitHubUser creates and returns an instance of GitHub user with the boolean fields
// populated
func NewGitHubUser(githubAPIKey string) *GitHubUser {
	ghUser := GitHubUser{
		apiKey: githubAPIKey,

		clientV3: nil,
		clientV4: nil,

		loginName: "",

		IsContributor: false,
		IsSponsor:     false,
	}

	if ghUser.hasAPIKey() {
		// Use the v3 API to get the contributors because this doesn't seem to be supported by the v4 API yet
		clientV3, _ := ghUser.authenticateV3()
		ghUser.clientV3 = clientV3

		// Use the v4 API to get sponsors because this doesn't seem to be supported in v3
		clientV4, _ := ghUser.authenticateV4()
		ghUser.clientV4 = clientV4
	}

	return &ghUser
}

/* -------------------- Exported Functions -------------------- */

// Load loads the user's data from GitHub
func (ghUser *GitHubUser) Load() error {
	err := ghUser.verifyGitHubClients()
	if err != nil {
		return err
	}

	err = ghUser.loadGitHubData()
	if err != nil {
		return err
	}

	return nil
}

/* -------------------- Unexported Functions -------------------- */

func (ghUser *GitHubUser) authenticateV3() (*ghb.Client, error) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghUser.apiKey},
	)

	oauthClient := oauth2.NewClient(context.Background(), src)
	client := ghb.NewClient(oauthClient)

	return client, nil
}

func (ghUser *GitHubUser) authenticateV4() (*githubv4.Client, error) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghUser.apiKey},
	)

	oauthClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(oauthClient)

	return client, nil
}

// hasAPIKey returns TRUE if the user has put a GitHub API key into their
// configuration and we've managed to find and read it
func (ghUser *GitHubUser) hasAPIKey() bool {
	return ghUser.apiKey != ""
}

func (ghUser *GitHubUser) loadGitHubData() error {
	var err error

	login, err := ghUser.loadLoginName()
	if err != nil {
		return err
	}
	ghUser.loginName = login

	var isContrib, isSponsor bool

	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		isContrib, err = ghUser.loadContributorStatus(ctx)
		return err
	})

	g.Go(func() error {
		isSponsor, err = ghUser.loadSponsorStatus(ctx)
		return err
	})

	err = g.Wait()
	if err != nil {
		return err
	}

	ghUser.IsContributor = isContrib
	ghUser.IsSponsor = isSponsor

	return nil
}

// loadLoginName figures out the GitHub user's login name from their API key
func (ghUser *GitHubUser) loadLoginName() (string, error) {
	user, _, err := ghUser.clientV3.Users.Get(context.Background(), "")
	if err != nil {
		return "", err
	}

	login := user.GetLogin()

	return login, nil
}

// loadContributorStatus figures out if this GitHub account has contributed to WTF
func (ghUser *GitHubUser) loadContributorStatus(ctx context.Context) (bool, error) {
	page := 1
	isContributor := false

	for {
		opts := &ghb.ListContributorsOptions{
			ListOptions: ghb.ListOptions{
				Page:    page,
				PerPage: 100,
			},
		}

		contributors, resp, err := ghUser.clientV3.Repositories.ListContributors(ctx, "wtfutil", "wtf", opts)
		if err != nil {
			return false, err
		}

		if resp.StatusCode != http.StatusOK || len(contributors) < 1 {
			break
		}

		for _, contrib := range contributors {
			if contrib.GetLogin() == ghUser.loginName {
				isContributor = true
				break
			}
		}

		page++
	}

	return isContributor, nil
}

// loadSponsorStatus figures out if this GitHub account has sponsored WTF
func (ghUser *GitHubUser) loadSponsorStatus(ctx context.Context) (bool, error) {
	vars := map[string]interface{}{
		"loginName": githubv4.String(ghUser.loginName),
	}

	err := ghUser.clientV4.Query(ctx, &sponsorQuery, vars)
	if err != nil {
		return false, err
	}

	isSponsor := false

	for _, spon := range sponsorQuery.User.SponsorshipsAsSponsor.Nodes {
		if spon.Sponsorable.SponsorsListing.Slug == "sponsors-senorprogrammer" {
			isSponsor = true
			break
		}
	}

	return isSponsor, nil
}

func (ghUser *GitHubUser) verifyGitHubClients() error {
	if ghUser.clientV3 == nil {
		return errors.New("github client v3 failed to load")
	}

	if ghUser.clientV4 == nil {
		return errors.New("github client v4 failed to load")
	}

	return nil
}
