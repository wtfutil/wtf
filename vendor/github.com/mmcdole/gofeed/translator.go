package gofeed

import (
	"fmt"
	"strings"
	"time"

	"github.com/mmcdole/gofeed/atom"
	"github.com/mmcdole/gofeed/extensions"
	"github.com/mmcdole/gofeed/internal/shared"
	"github.com/mmcdole/gofeed/rss"
)

// Translator converts a particular feed (atom.Feed or rss.Feed)
// into the generic Feed struct
type Translator interface {
	Translate(feed interface{}) (*Feed, error)
}

// DefaultRSSTranslator converts an rss.Feed struct
// into the generic Feed struct.
//
// This default implementation defines a set of
// mapping rules between rss.Feed -> Feed
// for each of the fields in Feed.
type DefaultRSSTranslator struct{}

// Translate converts an RSS feed into the universal
// feed type.
func (t *DefaultRSSTranslator) Translate(feed interface{}) (*Feed, error) {
	rss, found := feed.(*rss.Feed)
	if !found {
		return nil, fmt.Errorf("Feed did not match expected type of *rss.Feed")
	}

	result := &Feed{}
	result.Title = t.translateFeedTitle(rss)
	result.Description = t.translateFeedDescription(rss)
	result.Link = t.translateFeedLink(rss)
	result.FeedLink = t.translateFeedFeedLink(rss)
	result.Updated = t.translateFeedUpdated(rss)
	result.UpdatedParsed = t.translateFeedUpdatedParsed(rss)
	result.Published = t.translateFeedPublished(rss)
	result.PublishedParsed = t.translateFeedPublishedParsed(rss)
	result.Author = t.translateFeedAuthor(rss)
	result.Language = t.translateFeedLanguage(rss)
	result.Image = t.translateFeedImage(rss)
	result.Copyright = t.translateFeedCopyright(rss)
	result.Generator = t.translateFeedGenerator(rss)
	result.Categories = t.translateFeedCategories(rss)
	result.Items = t.translateFeedItems(rss)
	result.ITunesExt = rss.ITunesExt
	result.DublinCoreExt = rss.DublinCoreExt
	result.Extensions = rss.Extensions
	result.FeedVersion = rss.Version
	result.FeedType = "rss"
	return result, nil
}

func (t *DefaultRSSTranslator) translateFeedItem(rssItem *rss.Item) (item *Item) {
	item = &Item{}
	item.Title = t.translateItemTitle(rssItem)
	item.Description = t.translateItemDescription(rssItem)
	item.Content = t.translateItemContent(rssItem)
	item.Link = t.translateItemLink(rssItem)
	item.Published = t.translateItemPublished(rssItem)
	item.PublishedParsed = t.translateItemPublishedParsed(rssItem)
	item.Author = t.translateItemAuthor(rssItem)
	item.GUID = t.translateItemGUID(rssItem)
	item.Image = t.translateItemImage(rssItem)
	item.Categories = t.translateItemCategories(rssItem)
	item.Enclosures = t.translateItemEnclosures(rssItem)
	item.DublinCoreExt = rssItem.DublinCoreExt
	item.ITunesExt = rssItem.ITunesExt
	item.Extensions = rssItem.Extensions
	return
}

func (t *DefaultRSSTranslator) translateFeedTitle(rss *rss.Feed) (title string) {
	if rss.Title != "" {
		title = rss.Title
	} else if rss.DublinCoreExt != nil && rss.DublinCoreExt.Title != nil {
		title = t.firstEntry(rss.DublinCoreExt.Title)
	}
	return
}

func (t *DefaultRSSTranslator) translateFeedDescription(rss *rss.Feed) (desc string) {
	return rss.Description
}

func (t *DefaultRSSTranslator) translateFeedLink(rss *rss.Feed) (link string) {
	if rss.Link != "" {
		link = rss.Link
	} else if rss.ITunesExt != nil && rss.ITunesExt.Subtitle != "" {
		link = rss.ITunesExt.Subtitle
	}
	return
}

func (t *DefaultRSSTranslator) translateFeedFeedLink(rss *rss.Feed) (link string) {
	atomExtensions := t.extensionsForKeys([]string{"atom", "atom10", "atom03"}, rss.Extensions)
	for _, ex := range atomExtensions {
		if links, ok := ex["link"]; ok {
			for _, l := range links {
				if l.Attrs["Rel"] == "self" {
					link = l.Value
				}
			}
		}
	}
	return
}

func (t *DefaultRSSTranslator) translateFeedUpdated(rss *rss.Feed) (updated string) {
	if rss.LastBuildDate != "" {
		updated = rss.LastBuildDate
	} else if rss.DublinCoreExt != nil && rss.DublinCoreExt.Date != nil {
		updated = t.firstEntry(rss.DublinCoreExt.Date)
	}
	return
}

func (t *DefaultRSSTranslator) translateFeedUpdatedParsed(rss *rss.Feed) (updated *time.Time) {
	if rss.LastBuildDateParsed != nil {
		updated = rss.LastBuildDateParsed
	} else if rss.DublinCoreExt != nil && rss.DublinCoreExt.Date != nil {
		dateText := t.firstEntry(rss.DublinCoreExt.Date)
		date, err := shared.ParseDate(dateText)
		if err == nil {
			updated = &date
		}
	}
	return
}

func (t *DefaultRSSTranslator) translateFeedPublished(rss *rss.Feed) (published string) {
	return rss.PubDate
}

func (t *DefaultRSSTranslator) translateFeedPublishedParsed(rss *rss.Feed) (published *time.Time) {
	return rss.PubDateParsed
}

func (t *DefaultRSSTranslator) translateFeedAuthor(rss *rss.Feed) (author *Person) {
	if rss.ManagingEditor != "" {
		name, address := shared.ParseNameAddress(rss.ManagingEditor)
		author = &Person{}
		author.Name = name
		author.Email = address
	} else if rss.WebMaster != "" {
		name, address := shared.ParseNameAddress(rss.WebMaster)
		author = &Person{}
		author.Name = name
		author.Email = address
	} else if rss.DublinCoreExt != nil && rss.DublinCoreExt.Author != nil {
		dcAuthor := t.firstEntry(rss.DublinCoreExt.Author)
		name, address := shared.ParseNameAddress(dcAuthor)
		author = &Person{}
		author.Name = name
		author.Email = address
	} else if rss.DublinCoreExt != nil && rss.DublinCoreExt.Creator != nil {
		dcCreator := t.firstEntry(rss.DublinCoreExt.Creator)
		name, address := shared.ParseNameAddress(dcCreator)
		author = &Person{}
		author.Name = name
		author.Email = address
	} else if rss.ITunesExt != nil && rss.ITunesExt.Author != "" {
		name, address := shared.ParseNameAddress(rss.ITunesExt.Author)
		author = &Person{}
		author.Name = name
		author.Email = address
	}
	return
}

func (t *DefaultRSSTranslator) translateFeedLanguage(rss *rss.Feed) (language string) {
	if rss.Language != "" {
		language = rss.Language
	} else if rss.DublinCoreExt != nil && rss.DublinCoreExt.Language != nil {
		language = t.firstEntry(rss.DublinCoreExt.Language)
	}
	return
}

func (t *DefaultRSSTranslator) translateFeedImage(rss *rss.Feed) (image *Image) {
	if rss.Image != nil {
		image = &Image{}
		image.Title = rss.Image.Title
		image.URL = rss.Image.URL
	} else if rss.ITunesExt != nil && rss.ITunesExt.Image != "" {
		image = &Image{}
		image.URL = rss.ITunesExt.Image
	}
	return
}

func (t *DefaultRSSTranslator) translateFeedCopyright(rss *rss.Feed) (rights string) {
	if rss.Copyright != "" {
		rights = rss.Copyright
	} else if rss.DublinCoreExt != nil && rss.DublinCoreExt.Rights != nil {
		rights = t.firstEntry(rss.DublinCoreExt.Rights)
	}
	return
}

func (t *DefaultRSSTranslator) translateFeedGenerator(rss *rss.Feed) (generator string) {
	return rss.Generator
}

func (t *DefaultRSSTranslator) translateFeedCategories(rss *rss.Feed) (categories []string) {
	cats := []string{}
	if rss.Categories != nil {
		for _, c := range rss.Categories {
			cats = append(cats, c.Value)
		}
	}

	if rss.ITunesExt != nil && rss.ITunesExt.Keywords != "" {
		keywords := strings.Split(rss.ITunesExt.Keywords, ",")
		for _, k := range keywords {
			cats = append(cats, k)
		}
	}

	if rss.ITunesExt != nil && rss.ITunesExt.Categories != nil {
		for _, c := range rss.ITunesExt.Categories {
			cats = append(cats, c.Text)
			if c.Subcategory != nil {
				cats = append(cats, c.Subcategory.Text)
			}
		}
	}

	if rss.DublinCoreExt != nil && rss.DublinCoreExt.Subject != nil {
		for _, c := range rss.DublinCoreExt.Subject {
			cats = append(cats, c)
		}
	}

	if len(cats) > 0 {
		categories = cats
	}

	return
}

func (t *DefaultRSSTranslator) translateFeedItems(rss *rss.Feed) (items []*Item) {
	items = []*Item{}
	for _, i := range rss.Items {
		items = append(items, t.translateFeedItem(i))
	}
	return
}

func (t *DefaultRSSTranslator) translateItemTitle(rssItem *rss.Item) (title string) {
	if rssItem.Title != "" {
		title = rssItem.Title
	} else if rssItem.DublinCoreExt != nil && rssItem.DublinCoreExt.Title != nil {
		title = t.firstEntry(rssItem.DublinCoreExt.Title)
	}
	return
}

func (t *DefaultRSSTranslator) translateItemDescription(rssItem *rss.Item) (desc string) {
	if rssItem.Description != "" {
		desc = rssItem.Description
	} else if rssItem.DublinCoreExt != nil && rssItem.DublinCoreExt.Description != nil {
		desc = t.firstEntry(rssItem.DublinCoreExt.Description)
	}
	return
}

func (t *DefaultRSSTranslator) translateItemContent(rssItem *rss.Item) (content string) {
	return rssItem.Content
}

func (t *DefaultRSSTranslator) translateItemLink(rssItem *rss.Item) (link string) {
	return rssItem.Link
}

func (t *DefaultRSSTranslator) translateItemUpdated(rssItem *rss.Item) (updated string) {
	if rssItem.DublinCoreExt != nil && rssItem.DublinCoreExt.Date != nil {
		updated = t.firstEntry(rssItem.DublinCoreExt.Date)
	}
	return updated
}

func (t *DefaultRSSTranslator) translateItemUpdatedParsed(rssItem *rss.Item) (updated *time.Time) {
	if rssItem.DublinCoreExt != nil && rssItem.DublinCoreExt.Date != nil {
		updatedText := t.firstEntry(rssItem.DublinCoreExt.Date)
		updatedDate, err := shared.ParseDate(updatedText)
		if err == nil {
			updated = &updatedDate
		}
	}
	return
}

func (t *DefaultRSSTranslator) translateItemPublished(rssItem *rss.Item) (pubDate string) {
	if rssItem.PubDate != "" {
		return rssItem.PubDate
	} else if rssItem.DublinCoreExt != nil && rssItem.DublinCoreExt.Date != nil {
		return t.firstEntry(rssItem.DublinCoreExt.Date)
	}
	return
}

func (t *DefaultRSSTranslator) translateItemPublishedParsed(rssItem *rss.Item) (pubDate *time.Time) {
	if rssItem.PubDateParsed != nil {
		return rssItem.PubDateParsed
	} else if rssItem.DublinCoreExt != nil && rssItem.DublinCoreExt.Date != nil {
		pubDateText := t.firstEntry(rssItem.DublinCoreExt.Date)
		pubDateParsed, err := shared.ParseDate(pubDateText)
		if err == nil {
			pubDate = &pubDateParsed
		}
	}
	return
}

func (t *DefaultRSSTranslator) translateItemAuthor(rssItem *rss.Item) (author *Person) {
	if rssItem.Author != "" {
		name, address := shared.ParseNameAddress(rssItem.Author)
		author = &Person{}
		author.Name = name
		author.Email = address
	} else if rssItem.DublinCoreExt != nil && rssItem.DublinCoreExt.Author != nil {
		dcAuthor := t.firstEntry(rssItem.DublinCoreExt.Author)
		name, address := shared.ParseNameAddress(dcAuthor)
		author = &Person{}
		author.Name = name
		author.Email = address
	} else if rssItem.DublinCoreExt != nil && rssItem.DublinCoreExt.Creator != nil {
		dcCreator := t.firstEntry(rssItem.DublinCoreExt.Creator)
		name, address := shared.ParseNameAddress(dcCreator)
		author = &Person{}
		author.Name = name
		author.Email = address
	} else if rssItem.ITunesExt != nil && rssItem.ITunesExt.Author != "" {
		name, address := shared.ParseNameAddress(rssItem.ITunesExt.Author)
		author = &Person{}
		author.Name = name
		author.Email = address
	}
	return
}

func (t *DefaultRSSTranslator) translateItemGUID(rssItem *rss.Item) (guid string) {
	if rssItem.GUID != nil {
		guid = rssItem.GUID.Value
	}
	return
}

func (t *DefaultRSSTranslator) translateItemImage(rssItem *rss.Item) (image *Image) {
	if rssItem.ITunesExt != nil && rssItem.ITunesExt.Image != "" {
		image = &Image{}
		image.URL = rssItem.ITunesExt.Image
	}
	return
}

func (t *DefaultRSSTranslator) translateItemCategories(rssItem *rss.Item) (categories []string) {
	cats := []string{}
	if rssItem.Categories != nil {
		for _, c := range rssItem.Categories {
			cats = append(cats, c.Value)
		}
	}

	if rssItem.ITunesExt != nil && rssItem.ITunesExt.Keywords != "" {
		keywords := strings.Split(rssItem.ITunesExt.Keywords, ",")
		for _, k := range keywords {
			cats = append(cats, k)
		}
	}

	if rssItem.DublinCoreExt != nil && rssItem.DublinCoreExt.Subject != nil {
		for _, c := range rssItem.DublinCoreExt.Subject {
			cats = append(cats, c)
		}
	}

	if len(cats) > 0 {
		categories = cats
	}

	return
}

func (t *DefaultRSSTranslator) translateItemEnclosures(rssItem *rss.Item) (enclosures []*Enclosure) {
	if rssItem.Enclosure != nil {
		e := &Enclosure{}
		e.URL = rssItem.Enclosure.URL
		e.Type = rssItem.Enclosure.Type
		e.Length = rssItem.Enclosure.Length
		enclosures = []*Enclosure{e}
	}
	return
}

func (t *DefaultRSSTranslator) extensionsForKeys(keys []string, extensions ext.Extensions) (matches []map[string][]ext.Extension) {
	matches = []map[string][]ext.Extension{}

	if extensions == nil {
		return
	}

	for _, key := range keys {
		if match, ok := extensions[key]; ok {
			matches = append(matches, match)
		}
	}
	return
}

func (t *DefaultRSSTranslator) firstEntry(entries []string) (value string) {
	if entries == nil {
		return
	}

	if len(entries) == 0 {
		return
	}

	return entries[0]
}

// DefaultAtomTranslator converts an atom.Feed struct
// into the generic Feed struct.
//
// This default implementation defines a set of
// mapping rules between atom.Feed -> Feed
// for each of the fields in Feed.
type DefaultAtomTranslator struct{}

// Translate converts an Atom feed into the universal
// feed type.
func (t *DefaultAtomTranslator) Translate(feed interface{}) (*Feed, error) {
	atom, found := feed.(*atom.Feed)
	if !found {
		return nil, fmt.Errorf("Feed did not match expected type of *atom.Feed")
	}

	result := &Feed{}
	result.Title = t.translateFeedTitle(atom)
	result.Description = t.translateFeedDescription(atom)
	result.Link = t.translateFeedLink(atom)
	result.FeedLink = t.translateFeedFeedLink(atom)
	result.Updated = t.translateFeedUpdated(atom)
	result.UpdatedParsed = t.translateFeedUpdatedParsed(atom)
	result.Author = t.translateFeedAuthor(atom)
	result.Language = t.translateFeedLanguage(atom)
	result.Image = t.translateFeedImage(atom)
	result.Copyright = t.translateFeedCopyright(atom)
	result.Categories = t.translateFeedCategories(atom)
	result.Generator = t.translateFeedGenerator(atom)
	result.Items = t.translateFeedItems(atom)
	result.Extensions = atom.Extensions
	result.FeedVersion = atom.Version
	result.FeedType = "atom"
	return result, nil
}

func (t *DefaultAtomTranslator) translateFeedItem(entry *atom.Entry) (item *Item) {
	item = &Item{}
	item.Title = t.translateItemTitle(entry)
	item.Description = t.translateItemDescription(entry)
	item.Content = t.translateItemContent(entry)
	item.Link = t.translateItemLink(entry)
	item.Updated = t.translateItemUpdated(entry)
	item.UpdatedParsed = t.translateItemUpdatedParsed(entry)
	item.Published = t.translateItemPublished(entry)
	item.PublishedParsed = t.translateItemPublishedParsed(entry)
	item.Author = t.translateItemAuthor(entry)
	item.GUID = t.translateItemGUID(entry)
	item.Image = t.translateItemImage(entry)
	item.Categories = t.translateItemCategories(entry)
	item.Enclosures = t.translateItemEnclosures(entry)
	item.Extensions = entry.Extensions
	return
}

func (t *DefaultAtomTranslator) translateFeedTitle(atom *atom.Feed) (title string) {
	return atom.Title
}

func (t *DefaultAtomTranslator) translateFeedDescription(atom *atom.Feed) (desc string) {
	return atom.Subtitle
}

func (t *DefaultAtomTranslator) translateFeedLink(atom *atom.Feed) (link string) {
	l := t.firstLinkWithType("alternate", atom.Links)
	if l != nil {
		link = l.Href
	}
	return
}

func (t *DefaultAtomTranslator) translateFeedFeedLink(atom *atom.Feed) (link string) {
	feedLink := t.firstLinkWithType("self", atom.Links)
	if feedLink != nil {
		link = feedLink.Href
	}
	return
}

func (t *DefaultAtomTranslator) translateFeedUpdated(atom *atom.Feed) (updated string) {
	return atom.Updated
}

func (t *DefaultAtomTranslator) translateFeedUpdatedParsed(atom *atom.Feed) (updated *time.Time) {
	return atom.UpdatedParsed
}

func (t *DefaultAtomTranslator) translateFeedAuthor(atom *atom.Feed) (author *Person) {
	a := t.firstPerson(atom.Authors)
	if a != nil {
		feedAuthor := Person{}
		feedAuthor.Name = a.Name
		feedAuthor.Email = a.Email
		author = &feedAuthor
	}
	return
}

func (t *DefaultAtomTranslator) translateFeedLanguage(atom *atom.Feed) (language string) {
	return atom.Language
}

func (t *DefaultAtomTranslator) translateFeedImage(atom *atom.Feed) (image *Image) {
	if atom.Logo != "" {
		feedImage := Image{}
		feedImage.URL = atom.Logo
		image = &feedImage
	}
	return
}

func (t *DefaultAtomTranslator) translateFeedCopyright(atom *atom.Feed) (rights string) {
	return atom.Rights
}

func (t *DefaultAtomTranslator) translateFeedGenerator(atom *atom.Feed) (generator string) {
	if atom.Generator != nil {
		if atom.Generator.Value != "" {
			generator += atom.Generator.Value
		}
		if atom.Generator.Version != "" {
			generator += " v" + atom.Generator.Version
		}
		if atom.Generator.URI != "" {
			generator += " " + atom.Generator.URI
		}
		generator = strings.TrimSpace(generator)
	}
	return
}

func (t *DefaultAtomTranslator) translateFeedCategories(atom *atom.Feed) (categories []string) {
	if atom.Categories != nil {
		categories = []string{}
		for _, c := range atom.Categories {
			categories = append(categories, c.Term)
		}
	}
	return
}

func (t *DefaultAtomTranslator) translateFeedItems(atom *atom.Feed) (items []*Item) {
	items = []*Item{}
	for _, entry := range atom.Entries {
		items = append(items, t.translateFeedItem(entry))
	}
	return
}

func (t *DefaultAtomTranslator) translateItemTitle(entry *atom.Entry) (title string) {
	return entry.Title
}

func (t *DefaultAtomTranslator) translateItemDescription(entry *atom.Entry) (desc string) {
	return entry.Summary
}

func (t *DefaultAtomTranslator) translateItemContent(entry *atom.Entry) (content string) {
	if entry.Content != nil {
		content = entry.Content.Value
	}
	return
}

func (t *DefaultAtomTranslator) translateItemLink(entry *atom.Entry) (link string) {
	l := t.firstLinkWithType("alternate", entry.Links)
	if l != nil {
		link = l.Href
	}
	return
}

func (t *DefaultAtomTranslator) translateItemUpdated(entry *atom.Entry) (updated string) {
	return entry.Updated
}

func (t *DefaultAtomTranslator) translateItemUpdatedParsed(entry *atom.Entry) (updated *time.Time) {
	return entry.UpdatedParsed
}

func (t *DefaultAtomTranslator) translateItemPublished(entry *atom.Entry) (updated string) {
	return entry.Published
}

func (t *DefaultAtomTranslator) translateItemPublishedParsed(entry *atom.Entry) (updated *time.Time) {
	return entry.PublishedParsed
}

func (t *DefaultAtomTranslator) translateItemAuthor(entry *atom.Entry) (author *Person) {
	a := t.firstPerson(entry.Authors)
	if a != nil {
		author = &Person{}
		author.Name = a.Name
		author.Email = a.Email
	}
	return
}

func (t *DefaultAtomTranslator) translateItemGUID(entry *atom.Entry) (guid string) {
	return entry.ID
}

func (t *DefaultAtomTranslator) translateItemImage(entry *atom.Entry) (image *Image) {
	return nil
}

func (t *DefaultAtomTranslator) translateItemCategories(entry *atom.Entry) (categories []string) {
	if entry.Categories != nil {
		categories = []string{}
		for _, c := range entry.Categories {
			categories = append(categories, c.Term)
		}
	}
	return
}

func (t *DefaultAtomTranslator) translateItemEnclosures(entry *atom.Entry) (enclosures []*Enclosure) {
	if entry.Links != nil {
		enclosures = []*Enclosure{}
		for _, e := range entry.Links {
			if e.Rel == "enclosure" {
				enclosure := &Enclosure{}
				enclosure.URL = e.Href
				enclosure.Length = e.Length
				enclosure.Type = e.Type
				enclosures = append(enclosures, enclosure)
			}
		}

		if len(enclosures) == 0 {
			enclosures = nil
		}
	}
	return
}

func (t *DefaultAtomTranslator) firstLinkWithType(linkType string, links []*atom.Link) *atom.Link {
	if links == nil {
		return nil
	}

	for _, link := range links {
		if link.Rel == linkType {
			return link
		}
	}
	return nil
}

func (t *DefaultAtomTranslator) firstPerson(persons []*atom.Person) (person *atom.Person) {
	if persons == nil || len(persons) == 0 {
		return
	}

	person = persons[0]
	return
}
