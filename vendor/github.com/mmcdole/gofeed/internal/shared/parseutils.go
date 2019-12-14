package shared

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mmcdole/goxpp"
)

var (
	emailNameRgx = regexp.MustCompile(`^([^@]+@[^\s]+)\s+\(([^@]+)\)$`)
	nameEmailRgx = regexp.MustCompile(`^([^@]+)\s+\(([^@]+@[^)]+)\)$`)
	nameOnlyRgx  = regexp.MustCompile(`^([^@()]+)$`)
	emailOnlyRgx = regexp.MustCompile(`^([^@()]+@[^@()]+)$`)

	TruncatedEntity         = errors.New("truncated entity")
	InvalidNumericReference = errors.New("invalid numeric reference")
)

// FindRoot iterates through the tokens of an xml document until
// it encounters its first StartTag event.  It returns an error
// if it reaches EndDocument before finding a tag.
func FindRoot(p *xpp.XMLPullParser) (event xpp.XMLEventType, err error) {
	for {
		event, err = p.Next()
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

// NextTag iterates through the tokens until it reaches a StartTag or EndTag
// It is similar to goxpp's NextTag method except it wont throw an error if
// the next immediate token isnt a Start/EndTag.  Instead, it will continue to
// consume tokens until it hits a Start/EndTag or EndDocument.
func NextTag(p *xpp.XMLPullParser) (event xpp.XMLEventType, err error) {
	for {
		event, err = p.Next()
		if err != nil {
			return event, err
		}

		if event == xpp.StartTag || event == xpp.EndTag {
			break
		}

		if event == xpp.EndDocument {
			return event, fmt.Errorf("Failed to find NextTag before reaching the end of the document.")
		}

	}
	return
}

// ParseText is a helper function for parsing the text
// from the current element of the XMLPullParser.
// This function can handle parsing naked XML text from
// an element.
func ParseText(p *xpp.XMLPullParser) (string, error) {
	var text struct {
		Type     string `xml:"type,attr"`
		InnerXML string `xml:",innerxml"`
	}

	err := p.DecodeElement(&text)
	if err != nil {
		return "", err
	}

	result := text.InnerXML
	result = strings.TrimSpace(result)

	if strings.HasPrefix(result, "<![CDATA[") &&
		strings.HasSuffix(result, "]]>") {
		result = strings.TrimPrefix(result, "<![CDATA[")
		result = strings.TrimSuffix(result, "]]>")
		return result, nil
	}

	return DecodeEntities(result)
}

// DecodeEntities decodes escaped XML entities
// in a string and returns the unescaped string
func DecodeEntities(str string) (string, error) {
	data := []byte(str)
	buf := bytes.NewBuffer([]byte{})

	for len(data) > 0 {
		// Find the next entity
		idx := bytes.IndexByte(data, '&')
		if idx == -1 {
			buf.Write(data)
			break
		}

		// Write and skip everything before it
		buf.Write(data[:idx])
		data = data[idx+1:]

		if len(data) == 0 {
			return "", TruncatedEntity
		}

		// Find the end of the entity
		end := bytes.IndexByte(data, ';')
		if end == -1 {
			return "", TruncatedEntity
		}

		if data[0] == '#' {
			// Numerical character reference
			var str string
			base := 10

			if len(data) > 1 && data[1] == 'x' {
				str = string(data[2:end])
				base = 16
			} else {
				str = string(data[1:end])
			}

			i, err := strconv.ParseUint(str, base, 32)
			if err != nil {
				return "", InvalidNumericReference
			}

			buf.WriteRune(rune(i))
		} else {
			// Predefined entity
			name := string(data[:end])

			var c byte
			switch name {
			case "lt":
				c = '<'
			case "gt":
				c = '>'
			case "quot":
				c = '"'
			case "apos":
				c = '\''
			case "amp":
				c = '&'
			default:
				return "", fmt.Errorf("unknown predefined "+
					"entity &%s;", name)
			}

			buf.WriteByte(c)
		}

		// Skip the entity
		data = data[end+1:]
	}

	return buf.String(), nil
}

// ParseNameAddress parses name/email strings commonly
// found in RSS feeds of the format "Example Name (example@site.com)"
// and other variations of this format.
func ParseNameAddress(nameAddressText string) (name string, address string) {
	if nameAddressText == "" {
		return
	}

	if emailNameRgx.MatchString(nameAddressText) {
		result := emailNameRgx.FindStringSubmatch(nameAddressText)
		address = result[1]
		name = result[2]
	} else if nameEmailRgx.MatchString(nameAddressText) {
		result := nameEmailRgx.FindStringSubmatch(nameAddressText)
		name = result[1]
		address = result[2]
	} else if nameOnlyRgx.MatchString(nameAddressText) {
		result := nameOnlyRgx.FindStringSubmatch(nameAddressText)
		name = result[1]
	} else if emailOnlyRgx.MatchString(nameAddressText) {
		result := emailOnlyRgx.FindStringSubmatch(nameAddressText)
		address = result[1]
	}
	return
}
