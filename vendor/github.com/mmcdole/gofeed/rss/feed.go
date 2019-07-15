package rss

import (
	"encoding/json"
	"time"

	"github.com/mmcdole/gofeed/extensions"
)

// Feed is an RSS Feed
type Feed struct {
	Title               string                   `json:"title,omitempty"`
	Link                string                   `json:"link,omitempty"`
	Description         string                   `json:"description,omitempty"`
	Language            string                   `json:"language,omitempty"`
	Copyright           string                   `json:"copyright,omitempty"`
	ManagingEditor      string                   `json:"managingEditor,omitempty"`
	WebMaster           string                   `json:"webMaster,omitempty"`
	PubDate             string                   `json:"pubDate,omitempty"`
	PubDateParsed       *time.Time               `json:"pubDateParsed,omitempty"`
	LastBuildDate       string                   `json:"lastBuildDate,omitempty"`
	LastBuildDateParsed *time.Time               `json:"lastBuildDateParsed,omitempty"`
	Categories          []*Category              `json:"categories,omitempty"`
	Generator           string                   `json:"generator,omitempty"`
	Docs                string                   `json:"docs,omitempty"`
	TTL                 string                   `json:"ttl,omitempty"`
	Image               *Image                   `json:"image,omitempty"`
	Rating              string                   `json:"rating,omitempty"`
	SkipHours           []string                 `json:"skipHours,omitempty"`
	SkipDays            []string                 `json:"skipDays,omitempty"`
	Cloud               *Cloud                   `json:"cloud,omitempty"`
	TextInput           *TextInput               `json:"textInput,omitempty"`
	DublinCoreExt       *ext.DublinCoreExtension `json:"dcExt,omitempty"`
	ITunesExt           *ext.ITunesFeedExtension `json:"itunesExt,omitempty"`
	Extensions          ext.Extensions           `json:"extensions,omitempty"`
	Items               []*Item                  `json:"items"`
	Version             string                   `json:"version"`
}

func (f Feed) String() string {
	json, _ := json.MarshalIndent(f, "", "    ")
	return string(json)
}

// Item is an RSS Item
type Item struct {
	Title         string                   `json:"title,omitempty"`
	Link          string                   `json:"link,omitempty"`
	Description   string                   `json:"description,omitempty"`
	Content       string                   `json:"content,omitempty"`
	Author        string                   `json:"author,omitempty"`
	Categories    []*Category              `json:"categories,omitempty"`
	Comments      string                   `json:"comments,omitempty"`
	Enclosure     *Enclosure               `json:"enclosure,omitempty"`
	GUID          *GUID                    `json:"guid,omitempty"`
	PubDate       string                   `json:"pubDate,omitempty"`
	PubDateParsed *time.Time               `json:"pubDateParsed,omitempty"`
	Source        *Source                  `json:"source,omitempty"`
	DublinCoreExt *ext.DublinCoreExtension `json:"dcExt,omitempty"`
	ITunesExt     *ext.ITunesItemExtension `json:"itunesExt,omitempty"`
	Extensions    ext.Extensions           `json:"extensions,omitempty"`
}

// Image is an image that represents the feed
type Image struct {
	URL         string `json:"url,omitempty"`
	Link        string `json:"link,omitempty"`
	Title       string `json:"title,omitempty"`
	Width       string `json:"width,omitempty"`
	Height      string `json:"height,omitempty"`
	Description string `json:"description,omitempty"`
}

// Enclosure is a media object that is attached to
// the item
type Enclosure struct {
	URL    string `json:"url,omitempty"`
	Length string `json:"length,omitempty"`
	Type   string `json:"type,omitempty"`
}

// GUID is a unique identifier for an item
type GUID struct {
	Value       string `json:"value,omitempty"`
	IsPermalink string `json:"isPermalink,omitempty"`
}

// Source contains feed information for another
// feed if a given item came from that feed
type Source struct {
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
}

// Category is category metadata for Feeds and Entries
type Category struct {
	Domain string `json:"domain,omitempty"`
	Value  string `json:"value,omitempty"`
}

// TextInput specifies a text input box that
// can be displayed with the channel
type TextInput struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Link        string `json:"link,omitempty"`
}

// Cloud allows processes to register with a
// cloud to be notified of updates to the channel,
// implementing a lightweight publish-subscribe protocol
// for RSS feeds
type Cloud struct {
	Domain            string `json:"domain,omitempty"`
	Port              string `json:"port,omitempty"`
	Path              string `json:"path,omitempty"`
	RegisterProcedure string `json:"registerProcedure,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
}
