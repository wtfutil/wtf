package github

import (
	"context"
	"fmt"
	"net/http"
	"os"

	ghb "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubActivity struct {
	Notifications []*GithubNotification
	User          string
	apiKey        string
}

type GithubNotification struct {
	Title string
	Type  string
	URL   string
}

type GithubActivityThread struct {
	URL *string `json:"html_url,omitempty"`
}

func NewGithubActivities(user string) *GithubActivity {
	repo := GithubActivity{
		apiKey: os.Getenv("WTF_GITHUB_TOKEN"),
		User:   user,
	}

	return &repo
}

func (repo *GithubActivity) Refresh() {
	repo.Notifications, _ = repo.notifications()
}

func (repo *GithubActivity) notifications() ([]*GithubNotification, error) {
	oauthClient := repo.oauthClient()
	github := ghb.NewClient(oauthClient)

	opts := &ghb.NotificationListOptions{}

	notifications, _, err := github.Activity.ListNotifications(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	trueNotifications := []*GithubNotification{}
	for _, notification := range notifications {
		url := notification.GetSubject().GetURL()
		u := fmt.Sprintf(url)
		req, err := github.NewRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		notificationThread := new(GithubActivityThread)
		_, err = github.Do(context.Background(), req, &notificationThread)

		htmlURL := ""
		if err != nil {
			htmlURL = err.Error()
		} else {
			htmlURL = *notificationThread.URL
		}

		trueNotification := GithubNotification{
			Title: notification.GetSubject().GetTitle(),
			Type:  notification.GetSubject().GetType(),
			URL:   htmlURL,
		}

		trueNotifications = append(trueNotifications, &trueNotification)
	}
	return trueNotifications, nil
}

/* -------------------- Unexported Functions -------------------- */

func (repo *GithubActivity) oauthClient() *http.Client {
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: repo.apiKey},
	)

	return oauth2.NewClient(context.Background(), tokenService)
}
