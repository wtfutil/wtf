package atom

import (
	"encoding/base64"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	ext "github.com/mmcdole/gofeed/extensions"
	"github.com/mmcdole/gofeed/internal/shared"
	xpp "github.com/mmcdole/goxpp"
)

var (
	// Atom elements which contain URIs
	// https://tools.ietf.org/html/rfc4287
	uriElements = map[string]bool{
		"icon": true,
		"id":   true,
		"logo": true,
		"uri":  true,
		"url":  true, // atom 0.3
	}

	// Atom attributes which contain URIs
	// https://tools.ietf.org/html/rfc4287
	atomURIAttrs = map[string]bool{
		"href":   true,
		"scheme": true,
		"src":    true,
		"uri":    true,
	}
)

// Parser is an Atom Parser
type Parser struct {
	base *shared.XMLBase
}

// Parse parses an xml feed into an atom.Feed
func (ap *Parser) Parse(feed io.Reader) (*Feed, error) {
	p := xpp.NewXMLPullParser(feed, false, shared.NewReaderLabel)
	ap.base = &shared.XMLBase{URIAttrs: atomURIAttrs}

	_, err := ap.base.FindRoot(p)
	if err != nil {
		return nil, err
	}

	return ap.parseRoot(p)
}

func (ap *Parser) parseRoot(p *xpp.XMLPullParser) (*Feed, error) {
	if err := p.Expect(xpp.StartTag, "feed"); err != nil {
		return nil, err
	}

	atom := &Feed{}
	atom.Entries = []*Entry{}
	atom.Version = ap.parseVersion(p)
	atom.Language = ap.parseLanguage(p)

	contributors := []*Person{}
	authors := []*Person{}
	categories := []*Category{}
	links := []*Link{}
	extensions := ext.Extensions{}

	for {
		tok, err := ap.base.NextTag(p)
		if err != nil {
			return nil, err
		}

		if tok == xpp.EndTag {
			break
		}

		if tok == xpp.StartTag {

			name := strings.ToLower(p.Name)

			if shared.IsExtension(p) {
				e, err := shared.ParseExtension(extensions, p)
				if err != nil {
					return nil, err
				}
				extensions = e
			} else if name == "title" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				atom.Title = result
			} else if name == "id" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				atom.ID = result
			} else if name == "updated" ||
				name == "modified" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				atom.Updated = result
				date, err := shared.ParseDate(result)
				if err == nil {
					utcDate := date.UTC()
					atom.UpdatedParsed = &utcDate
				}
			} else if name == "subtitle" ||
				name == "tagline" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				atom.Subtitle = result
			} else if name == "link" {
				result, err := ap.parseLink(p)
				if err != nil {
					return nil, err
				}
				links = append(links, result)
			} else if name == "generator" {
				result, err := ap.parseGenerator(p)
				if err != nil {
					return nil, err
				}
				atom.Generator = result
			} else if name == "icon" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				atom.Icon = result
			} else if name == "logo" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				atom.Logo = result
			} else if name == "rights" ||
				name == "copyright" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				atom.Rights = result
			} else if name == "contributor" {
				result, err := ap.parsePerson("contributor", p)
				if err != nil {
					return nil, err
				}
				contributors = append(contributors, result)
			} else if name == "author" {
				result, err := ap.parsePerson("author", p)
				if err != nil {
					return nil, err
				}
				authors = append(authors, result)
			} else if name == "category" {
				result, err := ap.parseCategory(p)
				if err != nil {
					return nil, err
				}
				categories = append(categories, result)
			} else if name == "entry" {
				result, err := ap.parseEntry(p)
				if err != nil {
					return nil, err
				}
				atom.Entries = append(atom.Entries, result)
			} else {
				err := p.Skip()
				if err != nil {
					return nil, err
				}
			}
		}
	}

	if len(categories) > 0 {
		atom.Categories = categories
	}

	if len(authors) > 0 {
		atom.Authors = authors
	}

	if len(contributors) > 0 {
		atom.Contributors = contributors
	}

	if len(links) > 0 {
		atom.Links = links
	}

	if len(extensions) > 0 {
		atom.Extensions = extensions
	}

	if err := p.Expect(xpp.EndTag, "feed"); err != nil {
		return nil, err
	}

	return atom, nil
}

func (ap *Parser) parseEntry(p *xpp.XMLPullParser) (*Entry, error) {
	if err := p.Expect(xpp.StartTag, "entry"); err != nil {
		return nil, err
	}
	entry := &Entry{}

	contributors := []*Person{}
	authors := []*Person{}
	categories := []*Category{}
	links := []*Link{}
	extensions := ext.Extensions{}

	for {
		tok, err := ap.base.NextTag(p)
		if err != nil {
			return nil, err
		}

		if tok == xpp.EndTag {
			break
		}

		if tok == xpp.StartTag {

			name := strings.ToLower(p.Name)

			if shared.IsExtension(p) {
				e, err := shared.ParseExtension(extensions, p)
				if err != nil {
					return nil, err
				}
				extensions = e
			} else if name == "title" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				entry.Title = result
			} else if name == "id" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				entry.ID = result
			} else if name == "rights" ||
				name == "copyright" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				entry.Rights = result
			} else if name == "summary" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				entry.Summary = result
			} else if name == "source" {
				result, err := ap.parseSource(p)
				if err != nil {
					return nil, err
				}
				entry.Source = result
			} else if name == "updated" ||
				name == "modified" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				entry.Updated = result
				date, err := shared.ParseDate(result)
				if err == nil {
					utcDate := date.UTC()
					entry.UpdatedParsed = &utcDate
				}
			} else if name == "contributor" {
				result, err := ap.parsePerson("contributor", p)
				if err != nil {
					return nil, err
				}
				contributors = append(contributors, result)
			} else if name == "author" {
				result, err := ap.parsePerson("author", p)
				if err != nil {
					return nil, err
				}
				authors = append(authors, result)
			} else if name == "category" {
				result, err := ap.parseCategory(p)
				if err != nil {
					return nil, err
				}
				categories = append(categories, result)
			} else if name == "link" {
				result, err := ap.parseLink(p)
				if err != nil {
					return nil, err
				}
				links = append(links, result)
			} else if name == "published" ||
				name == "issued" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				entry.Published = result
				date, err := shared.ParseDate(result)
				if err == nil {
					utcDate := date.UTC()
					entry.PublishedParsed = &utcDate
				}
			} else if name == "content" {
				result, err := ap.parseContent(p)
				if err != nil {
					return nil, err
				}
				entry.Content = result
			} else {
				err := p.Skip()
				if err != nil {
					return nil, err
				}
			}
		}
	}

	if len(categories) > 0 {
		entry.Categories = categories
	}

	if len(authors) > 0 {
		entry.Authors = authors
	}

	if len(links) > 0 {
		entry.Links = links
	}

	if len(contributors) > 0 {
		entry.Contributors = contributors
	}

	if len(extensions) > 0 {
		entry.Extensions = extensions
	}

	if err := p.Expect(xpp.EndTag, "entry"); err != nil {
		return nil, err
	}

	return entry, nil
}

func (ap *Parser) parseSource(p *xpp.XMLPullParser) (*Source, error) {

	if err := p.Expect(xpp.StartTag, "source"); err != nil {
		return nil, err
	}

	source := &Source{}

	contributors := []*Person{}
	authors := []*Person{}
	categories := []*Category{}
	links := []*Link{}
	extensions := ext.Extensions{}

	for {
		tok, err := ap.base.NextTag(p)
		if err != nil {
			return nil, err
		}

		if tok == xpp.EndTag {
			break
		}

		if tok == xpp.StartTag {

			name := strings.ToLower(p.Name)

			if shared.IsExtension(p) {
				e, err := shared.ParseExtension(extensions, p)
				if err != nil {
					return nil, err
				}
				extensions = e
			} else if name == "title" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				source.Title = result
			} else if name == "id" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				source.ID = result
			} else if name == "updated" ||
				name == "modified" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				source.Updated = result
				date, err := shared.ParseDate(result)
				if err == nil {
					utcDate := date.UTC()
					source.UpdatedParsed = &utcDate
				}
			} else if name == "subtitle" ||
				name == "tagline" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				source.Subtitle = result
			} else if name == "link" {
				result, err := ap.parseLink(p)
				if err != nil {
					return nil, err
				}
				links = append(links, result)
			} else if name == "generator" {
				result, err := ap.parseGenerator(p)
				if err != nil {
					return nil, err
				}
				source.Generator = result
			} else if name == "icon" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				source.Icon = result
			} else if name == "logo" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				source.Logo = result
			} else if name == "rights" ||
				name == "copyright" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				source.Rights = result
			} else if name == "contributor" {
				result, err := ap.parsePerson("contributor", p)
				if err != nil {
					return nil, err
				}
				contributors = append(contributors, result)
			} else if name == "author" {
				result, err := ap.parsePerson("author", p)
				if err != nil {
					return nil, err
				}
				authors = append(authors, result)
			} else if name == "category" {
				result, err := ap.parseCategory(p)
				if err != nil {
					return nil, err
				}
				categories = append(categories, result)
			} else {
				err := p.Skip()
				if err != nil {
					return nil, err
				}
			}
		}
	}

	if len(categories) > 0 {
		source.Categories = categories
	}

	if len(authors) > 0 {
		source.Authors = authors
	}

	if len(contributors) > 0 {
		source.Contributors = contributors
	}

	if len(links) > 0 {
		source.Links = links
	}

	if len(extensions) > 0 {
		source.Extensions = extensions
	}

	if err := p.Expect(xpp.EndTag, "source"); err != nil {
		return nil, err
	}

	return source, nil
}

func (ap *Parser) parseContent(p *xpp.XMLPullParser) (*Content, error) {
	c := &Content{}
	c.Type = p.Attribute("type")
	c.Src = p.Attribute("src")

	text, err := ap.parseAtomText(p)
	if err != nil {
		return nil, err
	}
	c.Value = text

	return c, nil
}

func (ap *Parser) parsePerson(name string, p *xpp.XMLPullParser) (*Person, error) {

	if err := p.Expect(xpp.StartTag, name); err != nil {
		return nil, err
	}

	person := &Person{}

	for {
		tok, err := ap.base.NextTag(p)
		if err != nil {
			return nil, err
		}

		if tok == xpp.EndTag {
			break
		}

		if tok == xpp.StartTag {

			name := strings.ToLower(p.Name)

			if name == "name" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				person.Name = result
			} else if name == "email" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				person.Email = result
			} else if name == "uri" ||
				name == "url" ||
				name == "homepage" {
				result, err := ap.parseAtomText(p)
				if err != nil {
					return nil, err
				}
				person.URI = result
			} else {
				err := p.Skip()
				if err != nil {
					return nil, err
				}
			}
		}
	}

	if err := p.Expect(xpp.EndTag, name); err != nil {
		return nil, err
	}

	return person, nil
}

func (ap *Parser) parseLink(p *xpp.XMLPullParser) (*Link, error) {
	if err := p.Expect(xpp.StartTag, "link"); err != nil {
		return nil, err
	}

	l := &Link{}
	l.Href = p.Attribute("href")
	l.Hreflang = p.Attribute("hreflang")
	l.Type = p.Attribute("type")
	l.Length = p.Attribute("length")
	l.Title = p.Attribute("title")
	l.Rel = p.Attribute("rel")
	if l.Rel == "" {
		l.Rel = "alternate"
	}

	if err := p.Skip(); err != nil {
		return nil, err
	}

	if err := p.Expect(xpp.EndTag, "link"); err != nil {
		return nil, err
	}
	return l, nil
}

func (ap *Parser) parseCategory(p *xpp.XMLPullParser) (*Category, error) {
	if err := p.Expect(xpp.StartTag, "category"); err != nil {
		return nil, err
	}

	c := &Category{}
	c.Term = p.Attribute("term")
	c.Scheme = p.Attribute("scheme")
	c.Label = p.Attribute("label")

	if err := p.Skip(); err != nil {
		return nil, err
	}

	if err := p.Expect(xpp.EndTag, "category"); err != nil {
		return nil, err
	}
	return c, nil
}

func (ap *Parser) parseGenerator(p *xpp.XMLPullParser) (*Generator, error) {

	if err := p.Expect(xpp.StartTag, "generator"); err != nil {
		return nil, err
	}

	g := &Generator{}

	uri := p.Attribute("uri") // Atom 1.0
	url := p.Attribute("url") // Atom 0.3

	if uri != "" {
		g.URI = uri
	} else if url != "" {
		g.URI = url
	}

	g.Version = p.Attribute("version")

	result, err := ap.parseAtomText(p)
	if err != nil {
		return nil, err
	}

	g.Value = result

	if err := p.Expect(xpp.EndTag, "generator"); err != nil {
		return nil, err
	}

	return g, nil
}

func (ap *Parser) parseAtomText(p *xpp.XMLPullParser) (string, error) {

	var text struct {
		Type     string `xml:"type,attr"`
		Mode     string `xml:"mode,attr"`
		InnerXML string `xml:",innerxml"`
	}

	err := p.DecodeElement(&text)
	if err != nil {
		return "", err
	}

	result := text.InnerXML
	result = strings.TrimSpace(result)

	lowerType := strings.ToLower(text.Type)
	lowerMode := strings.ToLower(text.Mode)

	if strings.Contains(result, "<![CDATA[") {
		result = shared.StripCDATA(result)
		if lowerType == "html" || strings.Contains(lowerType, "xhtml") {
			result, _ = ap.base.ResolveHTML(result)
		}
	} else {
		// decode non-CDATA contents depending on type

		if lowerType == "text" ||
			strings.HasPrefix(lowerType, "text/") ||
			(lowerType == "" && lowerMode == "") {
			result, err = shared.DecodeEntities(result)
		} else if strings.Contains(lowerType, "xhtml") {
			result = ap.stripWrappingDiv(result)
			result, _ = ap.base.ResolveHTML(result)
		} else if lowerType == "html" {
			result = ap.stripWrappingDiv(result)
			result, err = shared.DecodeEntities(result)
			if err == nil {
				result, _ = ap.base.ResolveHTML(result)
			}
		} else {
			decodedStr, err := base64.StdEncoding.DecodeString(result)
			if err == nil {
				result = string(decodedStr)
			}
		}
	}

	// resolve relative URIs in URI-containing elements according to xml:base
	name := strings.ToLower(p.Name)
	if uriElements[name] {
		resolved, err := ap.base.ResolveURL(result)
		if err == nil {
			result = resolved
		}
	}

	return result, err
}

func (ap *Parser) parseLanguage(p *xpp.XMLPullParser) string {
	return p.Attribute("lang")
}

func (ap *Parser) parseVersion(p *xpp.XMLPullParser) string {
	ver := p.Attribute("version")
	if ver != "" {
		return ver
	}

	ns := p.Attribute("xmlns")
	if ns == "http://purl.org/atom/ns#" {
		return "0.3"
	}

	if ns == "http://www.w3.org/2005/Atom" {
		return "1.0"
	}

	return ""
}

func (ap *Parser) stripWrappingDiv(content string) (result string) {
	result = content
	r := strings.NewReader(result)
	doc, err := goquery.NewDocumentFromReader(r)
	if err == nil {
		root := doc.Find("body").Children()
		if root.Is("div") && root.Siblings().Size() == 0 {
			html, err := root.Unwrap().Html()
			if err == nil {
				result = html
			}
		}
	}
	return
}
