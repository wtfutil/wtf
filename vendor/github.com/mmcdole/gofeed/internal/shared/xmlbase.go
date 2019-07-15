package shared

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"net/url"
	"strings"

	"github.com/mmcdole/goxpp"
)

var (
	// HTML attributes which contain URIs
	// https://pythonhosted.org/feedparser/resolving-relative-links.html
	// To catch every possible URI attribute is non-trivial:
	// https://stackoverflow.com/questions/2725156/complete-list-of-html-tag-attributes-which-have-a-url-value
	htmlURIAttrs = map[string]bool{
		"action":     true,
		"background": true,
		"cite":       true,
		"codebase":   true,
		"data":       true,
		"href":       true,
		"poster":     true,
		"profile":    true,
		"scheme":     true,
		"src":        true,
		"uri":        true,
		"usemap":     true,
	}
)

type urlStack []*url.URL

func (s *urlStack) push(u *url.URL) {
	*s = append([]*url.URL{u}, *s...)
}

func (s *urlStack) pop() *url.URL {
	if s == nil || len(*s) == 0 {
		return nil
	}
	var top *url.URL
	top, *s = (*s)[0], (*s)[1:]
	return top
}

func (s *urlStack) top() *url.URL {
	if s == nil || len(*s) == 0 {
		return nil
	}
	return (*s)[0]
}

type XMLBase struct {
	stack    urlStack
	URIAttrs map[string]bool
}

// FindRoot iterates through the tokens of an xml document until
// it encounters its first StartTag event.  It returns an error
// if it reaches EndDocument before finding a tag.
func (b *XMLBase) FindRoot(p *xpp.XMLPullParser) (event xpp.XMLEventType, err error) {
	for {
		event, err = b.NextTag(p)
		if err != nil {
			return event, err
		}
		if event == xpp.StartTag {
			break
		}

		if event == xpp.EndDocument {
			return event, fmt.Errorf("Failed to find root node before document end.")
		}
	}
	return
}

// XMLBase.NextTag iterates through the tokens until it reaches a StartTag or
// EndTag It maintains the urlStack upon encountering StartTag and EndTags, so
// that the top of the stack (accessible through the CurrentBase() and
// CurrentBaseURL() methods) is the absolute base URI by which relative URIs
// should be resolved.
//
// NextTag is similar to goxpp's NextTag method except it wont throw an error
// if the next immediate token isnt a Start/EndTag.  Instead, it will continue
// to consume tokens until it hits a Start/EndTag or EndDocument.
func (b *XMLBase) NextTag(p *xpp.XMLPullParser) (event xpp.XMLEventType, err error) {
	for {

		if p.Event == xpp.EndTag {
			// Pop xml:base after each end tag
			b.pop()
		}

		event, err = p.Next()
		if err != nil {
			return event, err
		}

		if event == xpp.EndTag {
			break
		}

		if event == xpp.StartTag {
			base := parseBase(p)
			err = b.push(base)
			if err != nil {
				return
			}

			err = b.resolveAttrs(p)
			if err != nil {
				return
			}

			break
		}

		if event == xpp.EndDocument {
			return event, fmt.Errorf("Failed to find NextTag before reaching the end of the document.")
		}

	}
	return
}

func parseBase(p *xpp.XMLPullParser) string {
	xmlURI := "http://www.w3.org/XML/1998/namespace"
	for _, attr := range p.Attrs {
		if attr.Name.Local == "base" && attr.Name.Space == xmlURI {
			return attr.Value
		}
	}
	return ""
}

func (b *XMLBase) push(base string) error {
	newURL, err := url.Parse(base)
	if err != nil {
		return err
	}

	topURL := b.CurrentBaseURL()
	if topURL != nil {
		newURL = topURL.ResolveReference(newURL)
	}
	b.stack.push(newURL)
	return nil
}

// returns the popped base URL
func (b *XMLBase) pop() string {
	url := b.stack.pop()
	if url != nil {
		return url.String()
	}
	return ""
}

func (b *XMLBase) CurrentBaseURL() *url.URL {
	return b.stack.top()
}

func (b *XMLBase) CurrentBase() string {
	if url := b.CurrentBaseURL(); url != nil {
		return url.String()
	}
	return ""
}

// resolve the given string as a URL relative to current base
func (b *XMLBase) ResolveURL(u string) (string, error) {
	if b.CurrentBase() == "" {
		return u, nil
	}

	relURL, err := url.Parse(u)
	if err != nil {
		return u, err
	}
	curr := b.CurrentBaseURL()
	if curr.Path != "" && u != "" && curr.Path[len(curr.Path)-1] != '/' {
		// There's no reason someone would use a path in xml:base if they
		// didn't mean for it to be a directory
		curr.Path = curr.Path + "/"
	}
	absURL := b.CurrentBaseURL().ResolveReference(relURL)
	return absURL.String(), nil
}

// resolve relative URI attributes according to xml:base
func (b *XMLBase) resolveAttrs(p *xpp.XMLPullParser) error {
	for i, attr := range p.Attrs {
		lowerName := strings.ToLower(attr.Name.Local)
		if b.URIAttrs[lowerName] {
			absURL, err := b.ResolveURL(attr.Value)
			if err != nil {
				return err
			}
			p.Attrs[i].Value = absURL
		}
	}
	return nil
}

// Transforms html by resolving any relative URIs in attributes
// if an error occurs during parsing or serialization, then the original string
// is returned along with the error.
func (b *XMLBase) ResolveHTML(relHTML string) (string, error) {
	if b.CurrentBase() == "" {
		return relHTML, nil
	}

	htmlReader := strings.NewReader(relHTML)

	doc, err := html.Parse(htmlReader)
	if err != nil {
		return relHTML, err
	}

	var visit func(*html.Node)

	// recursively traverse HTML resolving any relative URIs in attributes
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for i, a := range n.Attr {
				if htmlURIAttrs[a.Key] {
					absVal, err := b.ResolveURL(a.Val)
					if err == nil {
						n.Attr[i].Val = absVal
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}

	visit(doc)
	var w bytes.Buffer
	err = html.Render(&w, doc)
	if err != nil {
		return relHTML, err
	}

	// html.Render() always writes a complete html5 document, so strip the html
	// and body tags
	absHTML := w.String()
	absHTML = strings.TrimPrefix(absHTML, "<html><head></head><body>")
	absHTML = strings.TrimSuffix(absHTML, "</body></html>")

	return absHTML, err
}
