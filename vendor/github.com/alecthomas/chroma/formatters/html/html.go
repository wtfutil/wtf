package html

import (
	"fmt"
	"html"
	"io"
	"sort"
	"strings"

	"github.com/alecthomas/chroma"
)

// Option sets an option of the HTML formatter.
type Option func(f *Formatter)

// Standalone configures the HTML formatter for generating a standalone HTML document.
func Standalone() Option { return func(f *Formatter) { f.standalone = true } }

// ClassPrefix sets the CSS class prefix.
func ClassPrefix(prefix string) Option { return func(f *Formatter) { f.prefix = prefix } }

// WithClasses emits HTML using CSS classes, rather than inline styles.
func WithClasses() Option { return func(f *Formatter) { f.Classes = true } }

// TabWidth sets the number of characters for a tab. Defaults to 8.
func TabWidth(width int) Option { return func(f *Formatter) { f.tabWidth = width } }

// WithLineNumbers formats output with line numbers.
func WithLineNumbers() Option {
	return func(f *Formatter) {
		f.lineNumbers = true
	}
}

// LineNumbersInTable will, when combined with WithLineNumbers, separate the line numbers
// and code in table td's, which make them copy-and-paste friendly.
func LineNumbersInTable() Option {
	return func(f *Formatter) {
		f.lineNumbersInTable = true
	}
}

// HighlightLines higlights the given line ranges with the Highlight style.
//
// A range is the beginning and ending of a range as 1-based line numbers, inclusive.
func HighlightLines(ranges [][2]int) Option {
	return func(f *Formatter) {
		f.highlightRanges = ranges
		sort.Sort(f.highlightRanges)
	}
}

// BaseLineNumber sets the initial number to start line numbering at. Defaults to 1.
func BaseLineNumber(n int) Option {
	return func(f *Formatter) {
		f.baseLineNumber = n
	}
}

// New HTML formatter.
func New(options ...Option) *Formatter {
	f := &Formatter{
		baseLineNumber: 1,
	}
	for _, option := range options {
		option(f)
	}
	return f
}

// Formatter that generates HTML.
type Formatter struct {
	standalone         bool
	prefix             string
	Classes            bool // Exported field to detect when classes are being used
	tabWidth           int
	lineNumbers        bool
	lineNumbersInTable bool
	highlightRanges    highlightRanges
	baseLineNumber     int
}

type highlightRanges [][2]int

func (h highlightRanges) Len() int           { return len(h) }
func (h highlightRanges) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h highlightRanges) Less(i, j int) bool { return h[i][0] < h[j][0] }

func (f *Formatter) Format(w io.Writer, style *chroma.Style, iterator chroma.Iterator) (err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = perr.(error)
		}
	}()
	return f.writeHTML(w, style, iterator.Tokens())
}

func brightenOrDarken(colour chroma.Colour, factor float64) chroma.Colour {
	if colour.Brightness() < 0.5 {
		return colour.Brighten(factor)
	}
	return colour.Brighten(-factor)
}

// Ensure that style entries exist for highlighting, etc.
func (f *Formatter) restyle(style *chroma.Style) (*chroma.Style, error) {
	builder := style.Builder()
	bg := builder.Get(chroma.Background)
	// If we don't have a line highlight colour, make one that is 10% brighter/darker than the background.
	if !style.Has(chroma.LineHighlight) {
		highlight := chroma.StyleEntry{Background: bg.Background}
		highlight.Background = brightenOrDarken(highlight.Background, 0.1)
		builder.AddEntry(chroma.LineHighlight, highlight)
	}
	// If we don't have line numbers, use the text colour but 20% brighter/darker
	if !style.Has(chroma.LineNumbers) {
		text := chroma.StyleEntry{Colour: bg.Colour}
		text.Colour = brightenOrDarken(text.Colour, 0.5)
		builder.AddEntry(chroma.LineNumbers, text)
		builder.AddEntry(chroma.LineNumbersTable, text)
	}
	return builder.Build()
}

// We deliberately don't use html/template here because it is two orders of magnitude slower (benchmarked).
//
// OTOH we need to be super careful about correct escaping...
func (f *Formatter) writeHTML(w io.Writer, style *chroma.Style, tokens []*chroma.Token) (err error) { // nolint: gocyclo
	style, err = f.restyle(style)
	if err != nil {
		return err
	}
	css := f.styleToCSS(style)
	if !f.Classes {
		for t, style := range css {
			css[t] = compressStyle(style)
		}
	}
	if f.standalone {
		fmt.Fprint(w, "<html>\n")
		if f.Classes {
			fmt.Fprint(w, "<style type=\"text/css\">\n")
			f.WriteCSS(w, style)
			fmt.Fprintf(w, "body { %s; }\n", css[chroma.Background])
			fmt.Fprint(w, "</style>")
		}
		fmt.Fprintf(w, "<body%s>\n", f.styleAttr(css, chroma.Background))
	}

	wrapInTable := f.lineNumbers && f.lineNumbersInTable

	lines := splitTokensIntoLines(tokens)
	lineDigits := len(fmt.Sprintf("%d", len(lines)))
	highlightIndex := 0

	if wrapInTable {
		// List line numbers in its own <td>
		fmt.Fprintf(w, "<div%s>\n", f.styleAttr(css, chroma.Background))
		fmt.Fprintf(w, "<table%s><tr>", f.styleAttr(css, chroma.LineTable))
		fmt.Fprintf(w, "<td%s>\n", f.styleAttr(css, chroma.LineTableTD))
		fmt.Fprintf(w, "<pre%s>", f.styleAttr(css, chroma.Background))
		for index := range lines {
			line := f.baseLineNumber + index
			highlight, next := f.shouldHighlight(highlightIndex, line)
			if next {
				highlightIndex++
			}
			if highlight {
				fmt.Fprintf(w, "<span%s>", f.styleAttr(css, chroma.LineHighlight))
			}

			fmt.Fprintf(w, "<span%s>%*d\n</span>", f.styleAttr(css, chroma.LineNumbersTable), lineDigits, line)

			if highlight {
				fmt.Fprintf(w, "</span>")
			}
		}
		fmt.Fprint(w, "</pre></td>\n")
		fmt.Fprintf(w, "<td%s>\n", f.styleAttr(css, chroma.LineTableTD))
	}

	fmt.Fprintf(w, "<pre%s>", f.styleAttr(css, chroma.Background))
	highlightIndex = 0
	for index, tokens := range lines {
		// 1-based line number.
		line := f.baseLineNumber + index
		highlight, next := f.shouldHighlight(highlightIndex, line)
		if next {
			highlightIndex++
		}
		if highlight {
			fmt.Fprintf(w, "<span%s>", f.styleAttr(css, chroma.LineHighlight))
		}

		if f.lineNumbers && !wrapInTable {
			fmt.Fprintf(w, "<span%s>%*d</span>", f.styleAttr(css, chroma.LineNumbers), lineDigits, line)
		}

		for _, token := range tokens {
			html := html.EscapeString(token.String())
			attr := f.styleAttr(css, token.Type)
			if attr != "" {
				html = fmt.Sprintf("<span%s>%s</span>", attr, html)
			}
			fmt.Fprint(w, html)
		}
		if highlight {
			fmt.Fprintf(w, "</span>")
		}
	}

	fmt.Fprint(w, "</pre>")

	if wrapInTable {
		fmt.Fprint(w, "</td></tr></table>\n")
		fmt.Fprint(w, "</div>\n")
	}

	if f.standalone {
		fmt.Fprint(w, "\n</body>\n")
		fmt.Fprint(w, "</html>\n")
	}

	return nil
}

func (f *Formatter) shouldHighlight(highlightIndex, line int) (bool, bool) {
	next := false
	for highlightIndex < len(f.highlightRanges) && line > f.highlightRanges[highlightIndex][1] {
		highlightIndex++
		next = true
	}
	if highlightIndex < len(f.highlightRanges) {
		hrange := f.highlightRanges[highlightIndex]
		if line >= hrange[0] && line <= hrange[1] {
			return true, next
		}
	}
	return false, next
}

func (f *Formatter) class(t chroma.TokenType) string {
	for t != 0 {
		if cls, ok := chroma.StandardTypes[t]; ok {
			if cls != "" {
				return f.prefix + cls
			}
			return ""
		}
		t = t.Parent()
	}
	if cls := chroma.StandardTypes[t]; cls != "" {
		return f.prefix + cls
	}
	return ""
}

func (f *Formatter) styleAttr(styles map[chroma.TokenType]string, tt chroma.TokenType) string {
	if f.Classes {
		cls := f.class(tt)
		if cls == "" {
			return ""
		}
		return string(fmt.Sprintf(` class="%s"`, cls))
	}
	if _, ok := styles[tt]; !ok {
		tt = tt.SubCategory()
		if _, ok := styles[tt]; !ok {
			tt = tt.Category()
			if _, ok := styles[tt]; !ok {
				return ""
			}
		}
	}
	return fmt.Sprintf(` style="%s"`, styles[tt])
}

func (f *Formatter) tabWidthStyle() string {
	if f.tabWidth != 0 && f.tabWidth != 8 {
		return fmt.Sprintf("; -moz-tab-size: %[1]d; -o-tab-size: %[1]d; tab-size: %[1]d", f.tabWidth)
	}
	return ""
}

// WriteCSS writes CSS style definitions (without any surrounding HTML).
func (f *Formatter) WriteCSS(w io.Writer, style *chroma.Style) error {
	css := f.styleToCSS(style)
	// Special-case background as it is mapped to the outer ".chroma" class.
	if _, err := fmt.Fprintf(w, "/* %s */ .%schroma { %s }\n", chroma.Background, f.prefix, css[chroma.Background]); err != nil {
		return err
	}
	// Special-case code column of table to expand width.
	if f.lineNumbers && f.lineNumbersInTable {
		if _, err := fmt.Fprintf(w, "/* %s */ .%schroma .%s:last-child { width: 100%%; }",
			chroma.LineTableTD, f.prefix, f.class(chroma.LineTableTD)); err != nil {
			return err
		}
	}
	tts := []int{}
	for tt := range css {
		tts = append(tts, int(tt))
	}
	sort.Ints(tts)
	for _, ti := range tts {
		tt := chroma.TokenType(ti)
		if tt == chroma.Background {
			continue
		}
		styles := css[tt]
		if _, err := fmt.Fprintf(w, "/* %s */ .%schroma .%s { %s }\n", tt, f.prefix, f.class(tt), styles); err != nil {
			return err
		}
	}
	return nil
}

func (f *Formatter) styleToCSS(style *chroma.Style) map[chroma.TokenType]string {
	classes := map[chroma.TokenType]string{}
	bg := style.Get(chroma.Background)
	// Convert the style.
	for t := range chroma.StandardTypes {
		entry := style.Get(t)
		if t != chroma.Background {
			entry = entry.Sub(bg)
		}
		if entry.IsZero() {
			continue
		}
		classes[t] = StyleEntryToCSS(entry)
	}
	classes[chroma.Background] += f.tabWidthStyle()
	lineNumbersStyle := "margin-right: 0.4em; padding: 0 0.4em 0 0.4em;"
	// All rules begin with default rules followed by user provided rules
	classes[chroma.LineNumbers] = lineNumbersStyle + classes[chroma.LineNumbers]
	classes[chroma.LineNumbersTable] = lineNumbersStyle + classes[chroma.LineNumbersTable]
	classes[chroma.LineHighlight] = "display: block; width: 100%;" + classes[chroma.LineHighlight]
	classes[chroma.LineTable] = "border-spacing: 0; padding: 0; margin: 0; border: 0; width: auto; overflow: auto; display: block;" + classes[chroma.LineTable]
	classes[chroma.LineTableTD] = "vertical-align: top; padding: 0; margin: 0; border: 0;" + classes[chroma.LineTableTD]
	return classes
}

// StyleEntryToCSS converts a chroma.StyleEntry to CSS attributes.
func StyleEntryToCSS(e chroma.StyleEntry) string {
	styles := []string{}
	if e.Colour.IsSet() {
		styles = append(styles, "color: "+e.Colour.String())
	}
	if e.Background.IsSet() {
		styles = append(styles, "background-color: "+e.Background.String())
	}
	if e.Bold == chroma.Yes {
		styles = append(styles, "font-weight: bold")
	}
	if e.Italic == chroma.Yes {
		styles = append(styles, "font-style: italic")
	}
	if e.Underline == chroma.Yes {
		styles = append(styles, "text-decoration: underline")
	}
	return strings.Join(styles, "; ")
}

// Compress CSS attributes - remove spaces, transform 6-digit colours to 3.
func compressStyle(s string) string {
	parts := strings.Split(s, ";")
	out := []string{}
	for _, p := range parts {
		p = strings.Join(strings.Fields(p), " ")
		p = strings.Replace(p, ": ", ":", 1)
		if strings.Contains(p, "#") {
			c := p[len(p)-6:]
			if c[0] == c[1] && c[2] == c[3] && c[4] == c[5] {
				p = p[:len(p)-6] + c[0:1] + c[2:3] + c[4:5]
			}
		}
		out = append(out, p)
	}
	return strings.Join(out, ";")
}

func splitTokensIntoLines(tokens []*chroma.Token) (out [][]*chroma.Token) {
	line := []*chroma.Token{}
	for _, token := range tokens {
		for strings.Contains(token.Value, "\n") {
			parts := strings.SplitAfterN(token.Value, "\n", 2)
			// Token becomes the tail.
			token.Value = parts[1]

			// Append the head to the line and flush the line.
			clone := token.Clone()
			clone.Value = parts[0]
			line = append(line, clone)
			out = append(out, line)
			line = nil
		}
		line = append(line, token)
	}
	if len(line) > 0 {
		out = append(out, line)
	}
	return
}
