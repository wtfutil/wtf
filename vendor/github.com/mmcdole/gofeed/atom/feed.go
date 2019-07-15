package atom

import (
	"encoding/json"
	"time"

	"github.com/mmcdole/gofeed/extensions"
)

// Feed is an Atom Feed
type Feed struct {
	Title         string         `json:"title,omitempty"`
	ID            string         `json:"id,omitempty"`
	Updated       string         `json:"updated,omitempty"`
	UpdatedParsed *time.Time     `json:"updatedParsed,omitempty"`
	Subtitle      string         `json:"subtitle,omitempty"`
	Links         []*Link        `json:"links,omitempty"`
	Language      string         `json:"language,omitempty"`
	Generator     *Generator     `json:"generator,omitempty"`
	Icon          string         `json:"icon,omitempty"`
	Logo          string         `json:"logo,omitempty"`
	Rights        string         `json:"rights,omitempty"`
	Contributors  []*Person      `json:"contributors,omitempty"`
	Authors       []*Person      `json:"authors,omitempty"`
	Categories    []*Category    `json:"categories,omitempty"`
	Entries       []*Entry       `json:"entries"`
	Extensions    ext.Extensions `json:"extensions,omitempty"`
	Version       string         `json:"version"`
}

func (f Feed) String() string {
	json, _ := json.MarshalIndent(f, "", "    ")
	return string(json)
}

// Entry is an Atom Entry
type Entry struct {
	Title           string         `json:"title,omitempty"`
	ID              string         `json:"id,omitempty"`
	Updated         string         `json:"updated,omitempty"`
	UpdatedParsed   *time.Time     `json:"updatedParsed,omitempty"`
	Summary         string         `json:"summary,omitempty"`
	Authors         []*Person      `json:"authors,omitempty"`
	Contributors    []*Person      `json:"contributors,omitempty"`
	Categories      []*Category    `json:"categories,omitempty"`
	Links           []*Link        `json:"links,omitempty"`
	Rights          string         `json:"rights,omitempty"`
	Published       string         `json:"published,omitempty"`
	PublishedParsed *time.Time     `json:"publishedParsed,omitempty"`
	Source          *Source        `json:"source,omitempty"`
	Content         *Content       `json:"content,omitempty"`
	Extensions      ext.Extensions `json:"extensions,omitempty"`
}

// Category is category metadata for Feeds and Entries
type Category struct {
	Term   string `json:"term,omitempty"`
	Scheme string `json:"scheme,omitempty"`
	Label  string `json:"label,omitempty"`
}

// Person represents a person in an Atom feed
// for things like Authors, Contributors, etc
type Person struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	URI   string `json:"uri,omitempty"`
}

// Link is an Atom link that defines a reference
// from an entry or feed to a Web resource
type Link struct {
	Href     string `json:"href,omitempty"`
	Hreflang string `json:"hreflang,omitempty"`
	Rel      string `json:"rel,omitempty"`
	Type     string `json:"type,omitempty"`
	Title    string `json:"title,omitempty"`
	Length   string `json:"length,omitempty"`
}

// Content either contains or links to the content of
// the entry
type Content struct {
	Src   string `json:"src,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// Generator identifies the agent used to generate a
// feed, for debugging and other purposes.
type Generator struct {
	Value   string `json:"value,omitempty"`
	URI     string `json:"uri,omitempty"`
	Version string `json:"version,omitempty"`
}

// Source contains the feed information for another
// feed if a given entry came from that feed.
type Source struct {
	Title         string         `json:"title,omitempty"`
	ID            string         `json:"id,omitempty"`
	Updated       string         `json:"updated,omitempty"`
	UpdatedParsed *time.Time     `json:"updatedParsed,omitempty"`
	Subtitle      string         `json:"subtitle,omitempty"`
	Links         []*Link        `json:"links,omitempty"`
	Generator     *Generator     `json:"generator,omitempty"`
	Icon          string         `json:"icon,omitempty"`
	Logo          string         `json:"logo,omitempty"`
	Rights        string         `json:"rights,omitempty"`
	Contributors  []*Person      `json:"contributors,omitempty"`
	Authors       []*Person      `json:"authors,omitempty"`
	Categories    []*Category    `json:"categories,omitempty"`
	Extensions    ext.Extensions `json:"extensions,omitempty"`
}
