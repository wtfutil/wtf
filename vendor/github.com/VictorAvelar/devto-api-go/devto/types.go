package devto

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// User contains information about a devto account
type User struct {
	Name            string  `json:"name,omitempty"`
	Username        string  `json:"username,omitempty"`
	TwitterUsername string  `json:"twitter_username,omitempty"`
	GithubUsername  string  `json:"github_username,omitempty"`
	WebsiteURL      *WebURL `json:"website_url,omitempty"`
	ProfileImage    *WebURL `json:"profile_image,omitempty"`
	ProfileImage90  *WebURL `json:"profile_image_90,omitempty"`
}

// Organization describes a company or group that
// publishes content to devto.
type Organization struct {
	Name           string  `json:"name,omitempty"`
	Username       string  `json:"username,omitempty"`
	Slug           string  `json:"slug,omitempty"`
	ProfileImage   *WebURL `json:"profile_image,omitempty"`
	ProfileImage90 *WebURL `json:"profile_image_90,omitempty"`
}

// FlareTag represents an article's flare tag, if the article
// has one.
type FlareTag struct {
	Name         string `json:"name"`
	BGColorHex   string `json:"bg_color_hex"`
	TextColorHex string `json:"text_color_hex"`
}

// Tags are a group of topics related to an article
type Tags []string

// This deserialization logic is so that if a listed article
// originates from the /articles endpoint instead of
// /articles/me/*, its Published field is returned as true,
// since /articles exclusively returns articles that have been
// published.
type listedArticleJSON struct {
	TypeOf                 string        `json:"type_of,omitempty"`
	ID                     uint32        `json:"id,omitempty"`
	Title                  string        `json:"title,omitempty"`
	Description            string        `json:"description,omitempty"`
	CoverImage             *WebURL       `json:"cover_image,omitempty"`
	PublishedAt            *time.Time    `json:"published_at,omitempty"`
	PublishedTimestamp     string        `json:"published_timestamp,omitempty"`
	TagList                Tags          `json:"tag_list,omitempty"`
	Slug                   string        `json:"slug,omitempty"`
	Path                   string        `json:"path,omitempty"`
	URL                    *WebURL       `json:"url,omitempty"`
	CanonicalURL           *WebURL       `json:"canonical_url,omitempty"`
	CommentsCount          uint          `json:"comments_count,omitempty"`
	PositiveReactionsCount uint          `json:"positive_reactions_count,omitempty"`
	User                   User          `json:"user,omitempty"`
	Organization           *Organization `json:"organization,omitempty"`
	FlareTag               *FlareTag     `json:"flare_tag,omitempty"`
	BodyMarkdown           string        `json:"body_markdown,omitempty"`
	Published              *bool         `json:"published,omitempty"`
}

func (j *listedArticleJSON) listedArticle() ListedArticle {
	a := ListedArticle{
		TypeOf:                 j.TypeOf,
		ID:                     j.ID,
		Title:                  j.Title,
		Description:            j.Description,
		CoverImage:             j.CoverImage,
		PublishedAt:            j.PublishedAt,
		PublishedTimestamp:     j.PublishedTimestamp,
		TagList:                j.TagList,
		Slug:                   j.Slug,
		Path:                   j.Path,
		URL:                    j.URL,
		CanonicalURL:           j.CanonicalURL,
		CommentsCount:          j.CommentsCount,
		PositiveReactionsCount: j.PositiveReactionsCount,
		User:                   j.User,
		Organization:           j.Organization,
		FlareTag:               j.FlareTag,
		BodyMarkdown:           j.BodyMarkdown,
	}

	if j.Published != nil {
		a.Published = *j.Published
	} else {
		// "published" currently is included in the API
		// response for dev.to's /articles/me/* endpoints,
		// but not in /articles, so we are setting this
		// to true since /articles only returns articles
		// that are published.
		a.Published = true
	}
	return a
}

// ListedArticle represents an article returned from one of
// the list articles endpoints (/articles, /articles/me/*).
type ListedArticle struct {
	TypeOf                 string        `json:"type_of,omitempty"`
	ID                     uint32        `json:"id,omitempty"`
	Title                  string        `json:"title,omitempty"`
	Description            string        `json:"description,omitempty"`
	CoverImage             *WebURL       `json:"cover_image,omitempty"`
	PublishedAt            *time.Time    `json:"published_at,omitempty"`
	PublishedTimestamp     string        `json:"published_timestamp,omitempty"`
	TagList                Tags          `json:"tag_list,omitempty"`
	Slug                   string        `json:"slug,omitempty"`
	Path                   string        `json:"path,omitempty"`
	URL                    *WebURL       `json:"url,omitempty"`
	CanonicalURL           *WebURL       `json:"canonical_url,omitempty"`
	CommentsCount          uint          `json:"comments_count,omitempty"`
	PositiveReactionsCount uint          `json:"positive_reactions_count,omitempty"`
	User                   User          `json:"user,omitempty"`
	Organization           *Organization `json:"organization,omitempty"`
	FlareTag               *FlareTag     `json:"flare_tag,omitempty"`
	// Only present in "/articles/me/*" endpoints
	BodyMarkdown string `json:"body_markdown,omitempty"`
	Published    bool   `json:"published,omitempty"`
}

// UnmarshalJSON implements the JSON Unmarshaler interface.
func (a *ListedArticle) UnmarshalJSON(b []byte) error {
	var j listedArticleJSON
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}

	*a = j.listedArticle()
	return nil
}

// Article contains all the information related to a single
// information resource from devto.
type Article struct {
	TypeOf                 string     `json:"type_of,omitempty"`
	ID                     uint32     `json:"id,omitempty"`
	Title                  string     `json:"title,omitempty"`
	Description            string     `json:"description,omitempty"`
	CoverImage             *WebURL    `json:"cover_image,omitempty"`
	SocialImage            *WebURL    `json:"social_image,omitempty"`
	ReadablePublishDate    string     `json:"readable_publish_date"`
	Published              bool       `json:"published,omitempty"`
	PublishedAt            *time.Time `json:"published_at,omitempty"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	EditedAt               *time.Time `json:"edited_at,omitempty"`
	CrossPostedAt          *time.Time `json:"crossposted_at,omitempty"`
	LastCommentAt          *time.Time `json:"last_comment_at,omitempty"`
	TagList                string     `json:"tag_list,omitempty"`
	Tags                   Tags       `json:"tags,omitempty"`
	Slug                   string     `json:"slug,omitempty"`
	Path                   *WebURL    `json:"path,omitempty"`
	URL                    *WebURL    `json:"url,omitempty"`
	CanonicalURL           *WebURL    `json:"canonical_url,omitempty"`
	CommentsCount          uint       `json:"comments_count,omitempty"`
	PositiveReactionsCount uint       `json:"positive_reactions_count,omitempty"`
	User                   User       `json:"user,omitempty"`
	BodyHTML               string     `json:"body_html,omitempty"`
	BodyMarkdown           string     `json:"body_markdown,omitempty"`
}

// ArticleUpdate represents an update to an article; it is
// used as the payload in POST and PUT requests for writing
// articles.
type ArticleUpdate struct {
	Title          string   `json:"title"`
	BodyMarkdown   string   `json:"body_markdown"`
	Published      bool     `json:"published"`
	Series         *string  `json:"series"`
	MainImage      string   `json:"main_image,omitempty"`
	CanonicalURL   string   `json:"canonical_url,omitempty"`
	Description    string   `json:"description,omitempty"`
	Tags           []string `json:"tags,omitempty"`
	OrganizationID int32    `json:"organization_id,omitempty"`
}

// ArticleListOptions holds the query values to pass as
// query string parameter to the Articles List action.
type ArticleListOptions struct {
	Tags     string `url:"tag,omitempty"`
	Username string `url:"username,omitempty"`
	State    string `url:"state,omitempty"`
	Top      string `url:"top,omitempty"`
	Page     int    `url:"page,omitempty"`
}

// MyArticlesOptions defines pagination options used as query
// params in the dev.to "list my articles" endpoints.
type MyArticlesOptions struct {
	Page    int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
}

// WebURL is a class embed to override default unmarshal
// behavior.
type WebURL struct {
	*url.URL
}

// UnmarshalJSON overrides the default unmarshal behaviour
// for URL
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

// ErrorResponse is an error returned from a dev.to API
// endpoint.
type ErrorResponse struct {
	ErrorMessage string `json:"error"`
	Status       int    `json:"status"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf(`%d error: "%s"`, e.Status, e.ErrorMessage)
}
