package shared

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	xpp "github.com/mmcdole/goxpp"
)

var (
	emailNameRgx = regexp.MustCompile(`^([^@]+@[^\s]+)\s+\(([^@]+)\)$`)
	nameEmailRgx = regexp.MustCompile(`^([^@]+)\s+\(([^@]+@[^)]+)\)$`)
	nameOnlyRgx  = regexp.MustCompile(`^([^@()]+)$`)
	emailOnlyRgx = regexp.MustCompile(`^([^@()]+@[^@()]+)$`)

	TruncatedEntity         = errors.New("truncated entity")
	InvalidNumericReference = errors.New("invalid numeric reference")
)

const CDATA_START = "<![CDATA["
const CDATA_END = "]]>"

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

	if strings.Contains(result, CDATA_START) {
		return StripCDATA(result), nil
	}

	return DecodeEntities(result)
}

// StripCDATA removes CDATA tags from the string
// content outside of CDATA tags is passed via DecodeEntities
func StripCDATA(str string) string {
	buf := bytes.NewBuffer([]byte{})

	curr := 0

	for curr < len(str) {

		start := indexAt(str, CDATA_START, curr)

		if start == -1 {
			dec, _ := DecodeEntities(str[curr:])
			buf.Write([]byte(dec))
			return buf.String()
		}

		end := indexAt(str, CDATA_END, start)

		if end == -1 {
			dec, _ := DecodeEntities(str[curr:])
			buf.Write([]byte(dec))
			return buf.String()
		}

		buf.Write([]byte(str[start+len(CDATA_START) : end]))

		curr = curr + end + len(CDATA_END)
	}

	return buf.String()
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

func indexAt(str, substr string, start int) int {
	idx := strings.Index(str[start:], substr)
	if idx > -1 {
		idx += start
	}
	return idx
}
