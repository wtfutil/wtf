# gofeed 

[![Build Status](https://travis-ci.org/mmcdole/gofeed.svg?branch=master)](https://travis-ci.org/mmcdole/gofeed) [![Coverage Status](https://coveralls.io/repos/github/mmcdole/gofeed/badge.svg?branch=master)](https://coveralls.io/github/mmcdole/gofeed?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/mmcdole/gofeed)](https://goreportcard.com/report/github.com/mmcdole/gofeed) [![](https://godoc.org/github.com/mmcdole/gofeed?status.svg)](http://godoc.org/github.com/mmcdole/gofeed) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

The `gofeed` library is a robust feed parser that supports parsing both [RSS](https://en.wikipedia.org/wiki/RSS) and [Atom](https://en.wikipedia.org/wiki/Atom_(standard)) feeds.  The universal `gofeed.Parser` will parse and convert all feed types into a hybrid `gofeed.Feed` model.  You also have the option of parsing them into their respective `atom.Feed` and `rss.Feed` models using the feed specific `atom.Parser` or `rss.Parser`.  

##### Supported feed types:
* RSS 0.90
* Netscape RSS 0.91
* Userland RSS 0.91
* RSS 0.92
* RSS 0.93
* RSS 0.94
* RSS 1.0
* RSS 2.0
* Atom 0.3
* Atom 1.0

It also provides support for parsing several popular predefined extension modules, including [Dublin Core](http://dublincore.org/documents/dces/) and [Apple’s iTunes](https://help.apple.com/itc/podcasts_connect/#/itcb54353390), as well as arbitrary extensions.  See the [Extensions](#extensions) section for more details.

## Table of Contents
- [Overview](#overview)
- [Basic Usage](#basic-usage)
- [Advanced Usage](#advanced-usage)
- [Extensions](#extensions)
- [Invalid Feeds](#invalid-feeds)
- [Default Mappings](#default-mappings)
- [Dependencies](#dependencies)
- [License](#license)
- [Donate](#donate)
- [Credits](#credits)

## Overview

#### Universal Feed Parser

The universal `gofeed.Parser` works in 3 stages: detection, parsing and translation.  It first detects the feed type that it is currently parsing.  Then it uses a feed specific parser to parse the feed into its true representation which will be either a `rss.Feed` or `atom.Feed`.  These models cover every field possible for their respective feed types.  Finally, they are *translated* into a `gofeed.Feed` model that is a hybrid of both feed types.  Performing the universal feed parsing in these 3 stages allows for more flexibility and keeps the code base more maintainable by separating RSS and Atom parsing into seperate packages.

![Diagram](docs/sequence.png)

The translation step is done by anything which adheres to the `gofeed.Translator` interface.  The `DefaultRSSTranslator` and `DefaultAtomTranslator` are used behind the scenes when you use the `gofeed.Parser` with its default settings.  You can see how they translate fields from ```atom.Feed``` or ```rss.Feed``` to the universal ```gofeed.Feed``` struct in the [Default Mappings](#default-mappings) section.  However, should you disagree with the way certain fields are translated you can easily supply your own `gofeed.Translator` and override this behavior.  See the [Advanced Usage](#advanced-usage) section for an example how to do this.

#### Feed Specific Parsers

The `gofeed` library provides two feed specific parsers: `atom.Parser` and `rss.Parser`.  If the hybrid `gofeed.Feed` model that the universal `gofeed.Parser` produces does not contain a field from the `atom.Feed` or `rss.Feed` model that you require, it might be beneficial to use the feed specific parsers.  When using the `atom.Parser` or `rss.Parser` directly, you can access all of fields found in the `atom.Feed` and `rss.Feed` models.  It is also marginally faster because you are able to skip the translation step.

However, for the *vast* majority of users, the universal `gofeed.Parser` is the best way to parse feeds.  This allows the user of `gofeed` library to not care about the differences between RSS or Atom feeds.

## Basic Usage

#### Universal Feed Parser

The most common usage scenario will be to use ```gofeed.Parser``` to parse an arbitrary RSS or Atom feed into the hybrid ```gofeed.Feed``` model.  This hybrid model allows you to treat RSS and Atom feeds the same.

##### Parse a feed from an URL:

```go
fp := gofeed.NewParser()
feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
fmt.Println(feed.Title)
```

##### Parse a feed from a string:

```go
feedData := `<rss version="2.0">
<channel>
<title>Sample Feed</title>
</channel>
</rss>`
fp := gofeed.NewParser()
feed, _ := fp.ParseString(feedData)
fmt.Println(feed.Title)
```

##### Parse a feed from an io.Reader:

```go
file, _ := os.Open("/path/to/a/file.xml")
defer file.Close()
fp := gofeed.NewParser()
feed, _ := fp.Parse(file)
fmt.Println(feed.Title)
```

#### Feed Specific Parsers

You can easily use the `rss.Parser` and `atom.Parser` directly if you have a usage scenario that requires it:

##### Parse a RSS feed into a `rss.Feed`

```go
feedData := `<rss version="2.0">
<channel>
<webMaster>example@site.com (Example Name)</webMaster>
</channel>
</rss>`
fp := rss.Parser{}
rssFeed, _ := fp.Parse(strings.NewReader(feedData))
fmt.Println(rssFeed.WebMaster)
```

##### Parse an Atom feed into a `atom.Feed`

```go
feedData := `<feed xmlns="http://www.w3.org/2005/Atom">
<subtitle>Example Atom</subtitle>
</feed>`
fp := atom.Parser{}
atomFeed, _ := fp.Parse(strings.NewReader(feedData))
fmt.Println(atomFeed.Subtitle)
```

## Advanced Usage

##### Parse a feed while using a custom translator

The mappings and precedence order that are outlined in the [Default Mappings](#default-mappings) section are provided by the following two structs: `DefaultRSSTranslator` and `DefaultAtomTranslator`.  If you have fields that you think should have a different precedence, or if you want to make a translator that is aware of an unsupported extension you can do this by specifying your own RSS or Atom translator when using the `gofeed.Parser`.

Here is a simple example of creating a custom `Translator` that makes the `/rss/channel/itunes:author` field have a higher precedence than the `/rss/channel/managingEditor` field in RSS feeds.  We will wrap the existing `DefaultRSSTranslator` since we only want to change the behavior for a single field.

First we must define a custom translator:

```go

import (
    "fmt"

    "github.com/mmcdole/gofeed"
    "github.com/mmcdole/gofeed/rss"
)

type MyCustomTranslator struct {
    defaultTranslator *gofeed.DefaultRSSTranslator
}

func NewMyCustomTranslator() *MyCustomTranslator {
  t := &MyCustomTranslator{}
  
  // We create a DefaultRSSTranslator internally so we can wrap its Translate
  // call since we only want to modify the precedence for a single field.
  t.defaultTranslator = &gofeed.DefaultRSSTranslator{}
  return t
}

func (ct* MyCustomTranslator) Translate(feed interface{}) (*gofeed.Feed, error) {
	rss, found := feed.(*rss.Feed)
	if !found {
		return nil, fmt.Errorf("Feed did not match expected type of *rss.Feed")
	}

  f, err := ct.defaultTranslator.Translate(rss)
  if err != nil {
    return nil, err
  }
  
  if rss.ITunesExt != nil && rss.ITunesExt.Author != "" {
      f.Author = rss.ITunesExt.Author
  } else {
      f.Author = rss.ManagingEditor
  }
  return f
}
```

Next you must configure your `gofeed.Parser` to utilize the new `gofeed.Translator`:

```go
feedData := `<rss version="2.0">
<channel>
<managingEditor>Ender Wiggin</managingEditor>
<itunes:author>Valentine Wiggin</itunes:author>
</channel>
</rss>`
    
fp := gofeed.NewParser()
fp.RSSTranslator = NewMyCustomTranslator()
feed, _ := fp.ParseString(feedData)
fmt.Println(feed.Author) // Valentine Wiggin
```

## Extensions 

Every element which does not belong to the feed's default namespace is considered an extension by `gofeed`.  These are parsed and stored in a tree-like structure located at `Feed.Extensions` and `Item.Extensions`.  These fields should allow you to access and read any custom extension elements.

In addition to the generic handling of extensions, `gofeed` also has built in support for parsing certain popular extensions into their own structs for convenience.  It currently supports the [Dublin Core](http://dublincore.org/documents/dces/) and [Apple iTunes](https://help.apple.com/itc/podcasts_connect/#/itcb54353390) extensions which you can access at `Feed.ItunesExt`, `feed.DublinCoreExt` and `Item.ITunesExt` and `Item.DublinCoreExt`

## Invalid Feeds

A best-effort attempt is made at parsing broken and invalid XML feeds.  Currently, `gofeed` can succesfully parse feeds with the following issues:
- Unescaped/Naked Markup in feed elements
- Undeclared namespace prefixes
- Missing closing tags on certain elements
- Illegal tags within feed elements without namespace prefixes
- Missing "required" elements as specified by the respective feed specs.
- Incorrect date formats 

## Default Mappings

The ```DefaultRSSTranslator``` and the ```DefaultAtomTranslator``` map the following ```rss.Feed``` and ```atom.Feed``` fields to their respective ```gofeed.Feed``` fields.  They are listed in order of precedence (highest to lowest):


`gofeed.Feed` | RSS | Atom
--- | --- | ---
Title | /rss/channel/title<br>/rdf:RDF/channel/title<br>/rss/channel/dc:title<br>/rdf:RDF/channel/dc:title | /feed/title
Description | /rss/channel/description<br>/rdf:RDF/channel/description<br>/rss/channel/itunes:subtitle | /feed/subtitle<br>/feed/tagline
Link | /rss/channel/link<br>/rdf:RDF/channel/link | /feed/link[@rel=”alternate”]/@href<br>/feed/link[not(@rel)]/@href
FeedLink | /rss/channel/atom:link[@rel="self"]/@href<br>/rdf:RDF/channel/atom:link[@rel="self"]/@href | /feed/link[@rel="self"]/@href
Updated | /rss/channel/lastBuildDate<br>/rss/channel/dc:date<br>/rdf:RDF/channel/dc:date | /feed/updated<br>/feed/modified
Published | /rss/channel/pubDate |
Author | /rss/channel/managingEditor<br>/rss/channel/webMaster<br>/rss/channel/dc:author<br>/rdf:RDF/channel/dc:author<br>/rss/channel/dc:creator<br>/rdf:RDF/channel/dc:creator<br>/rss/channel/itunes:author | /feed/author
Language | /rss/channel/language<br>/rss/channel/dc:language<br>/rdf:RDF/channel/dc:language | /feed/@xml:lang
Image | /rss/channel/image<br>/rdf:RDF/image<br>/rss/channel/itunes:image | /feed/logo
Copyright | /rss/channel/copyright<br>/rss/channel/dc:rights<br>/rdf:RDF/channel/dc:rights | /feed/rights<br>/feed/copyright
Generator | /rss/channel/generator | /feed/generator
Categories | /rss/channel/category<br>/rss/channel/itunes:category<br>/rss/channel/itunes:keywords<br>/rss/channel/dc:subject<br>/rdf:RDF/channel/dc:subject | /feed/category


`gofeed.Item` | RSS | Atom
--- | --- | ---
Title | /rss/channel/item/title<br>/rdf:RDF/item/title<br>/rdf:RDF/item/dc:title<br>/rss/channel/item/dc:title | /feed/entry/title
Description | /rss/channel/item/description<br>/rdf:RDF/item/description<br>/rss/channel/item/dc:description<br>/rdf:RDF/item/dc:description | /feed/entry/summary
Content | | /feed/entry/content
Link | /rss/channel/item/link<br>/rdf:RDF/item/link | /feed/entry/link[@rel=”alternate”]/@href<br>/feed/entry/link[not(@rel)]/@href
Updated | /rss/channel/item/dc:date<br>/rdf:RDF/rdf:item/dc:date | /feed/entry/modified<br>/feed/entry/updated
Published | /rss/channel/item/pubDate<br>/rss/channel/item/dc:date | /feed/entry/published<br>/feed/entry/issued
Author | /rss/channel/item/author<br>/rss/channel/item/dc:author<br>/rdf:RDF/item/dc:author<br>/rss/channel/item/dc:creator<br>/rdf:RDF/item/dc:creator<br>/rss/channel/item/itunes:author | /feed/entry/author
Guid |  /rss/channel/item/guid | /feed/entry/id
Image | /rss/channel/item/itunes:image<br>/rss/channel/item/media:image |
Categories | /rss/channel/item/category<br>/rss/channel/item/dc:subject<br>/rss/channel/item/itunes:keywords<br>/rdf:RDF/channel/item/dc:subject | /feed/entry/category
Enclosures | /rss/channel/item/enclosure | /feed/entry/link[@rel=”enclosure”]

## Dependencies

* [goxpp](https://github.com/mmcdole/goxpp) - XML Pull Parser
* [goquery](https://github.com/PuerkitoBio/goquery) - Go jQuery-like interface
* [testify](https://github.com/stretchr/testify) - Unit test enhancements

## License

This project is licensed under the [MIT License](https://raw.githubusercontent.com/mmcdole/gofeed/master/LICENSE)

## Donate

I write open source software for fun. However, if you want to buy me a beer because you found something I wrote useful, feel free!

Bitcoin: 1CXrjBBkxgVNgKXRAq5MnsR7zzZbHvUHkJ

## Credits

* [Mark Pilgrim](https://en.wikipedia.org/wiki/Mark_Pilgrim) and [Kurt McKee](http://kurtmckee.org) for their work on the excellent [Universal Feed Parser](https://github.com/kurtmckee/feedparser) Python library.  This library was the inspiration for the `gofeed` library.
* [Dan MacTough](http://blog.mact.me) for his work on [node-feedparser](https://github.com/danmactough/node-feedparser).  It provided inspiration for the set of fields that should be covered in the hybrid `gofeed.Feed` model.
* [Matt Jibson](https://mattjibson.com/) for his date parsing function in the [goread](https://github.com/mjibson/goread) project.
* [Jim Teeuwen](https://github.com/jteeuwen) for his method of representing arbitrary feed extensions in the [go-pkg-rss](https://github.com/jteeuwen/go-pkg-rss) library.
