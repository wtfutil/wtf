package devto

import (
	"net/url"
	"strings"
	"time"
)

//User contains information about a devto account
type User struct {
	Name            string  `json:"name,omitempty"`
	Username        string  `json:"username,omitempty"`
	TwitterUsername string  `json:"twitter_username,omitempty"`
	GithubUsername  string  `json:"github_username,omitempty"`
	WebsiteURL      *WebURL `json:"website_url,omitempty"`
	ProfileImage    *WebURL `json:"profile_image,omitempty"`
	ProfileImage90  *WebURL `json:"profile_image_90,omitempty"`
}

//Organization describes a company or group that
//publishes content to devto.
type Organization struct {
	Name           string  `json:"name,omitempty"`
	Username       string  `json:"username,omitempty"`
	Slug           string  `json:"slug,omitempty"`
	ProfileImage   *WebURL `json:"profile_image,omitempty"`
	ProfileImage90 *WebURL `json:"profile_image_90,omitempty"`
}

//Tags are a group of topics related to an article
type Tags []string

//Article contains all the information related to a single
//information resource from devto.
type Article struct {
	TypeOf                 string       `json:"type_of,omitempty"`
	ID                     uint32       `json:"id,omitempty"`
	Title                  string       `json:"title,omitempty"`
	Description            string       `json:"description,omitempty"`
	CoverImage             *WebURL      `json:"cover_image,omitempty"`
	SocialImage            *WebURL      `json:"social_image,omitempty"`
	PublishedAt            *time.Time   `json:"published_at,omitempty"`
	EditedAt               *time.Time   `json:"edited_at,omitempty"`
	CrossPostedAt          *time.Time   `json:"crossposted_at,omitempty"`
	LastCommentAt          *time.Time   `json:"last_comment_at,omitempty"`
	TagList                Tags         `json:"tag_list,omitempty"`
	Tags                   string       `json:"tags,omitempty"`
	Slug                   string       `json:"slug,omitempty"`
	Path                   *WebURL      `json:"path,omitempty"`
	URL                    *WebURL      `json:"url,omitempty"`
	CanonicalURL           *WebURL      `json:"canonical_url,omitempty"`
	CommentsCount          uint         `json:"comments_count,omitempty"`
	PositiveReactionsCount uint         `json:"positive_reactions_count,omitempty"`
	PublishedTimestamp     *time.Time   `json:"published_timestamp,omitempty"`
	User                   User         `json:"user,omitempty"`
	Organization           Organization `json:"organization,omitempty"`
	BodyHTML               string       `json:"body_html,omitempty"`
	BodyMarkdown           string       `json:"body_markdown,omitempty"`
	Published              bool         `json:"published,omitempty"`
}

//ArticleListOptions holds the query values to pass as
//query string parameter to the Articles List action.
type ArticleListOptions struct {
	Tags     string `url:"tag,omitempty"`
	Username string `url:"username,omitempty"`
	State    string `url:"state,omitempty"`
	Top      string `url:"top,omitempty"`
	Page     int    `url:"page,omitempty"`
}

type WebURL struct {
	*url.URL
}

//UnmarshalJSON overrides the default unmarshal behaviour
//for URL
func (s *WebURL) UnmarshalJSON(b []byte) error {
	c := string(b)
	c = strings.Trim(c, "\"")
	uri, err := url.Parse(c)
	if err != nil {
		return err
	}
	s.URL = uri
	return nil
}
