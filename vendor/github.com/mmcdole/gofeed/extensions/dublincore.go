package ext

// DublinCoreExtension represents a feed extension
// for the Dublin Core specification.
type DublinCoreExtension struct {
	Title       []string `json:"title,omitempty"`
	Creator     []string `json:"creator,omitempty"`
	Author      []string `json:"author,omitempty"`
	Subject     []string `json:"subject,omitempty"`
	Description []string `json:"description,omitempty"`
	Publisher   []string `json:"publisher,omitempty"`
	Contributor []string `json:"contributor,omitempty"`
	Date        []string `json:"date,omitempty"`
	Type        []string `json:"type,omitempty"`
	Format      []string `json:"format,omitempty"`
	Identifier  []string `json:"identifier,omitempty"`
	Source      []string `json:"source,omitempty"`
	Language    []string `json:"language,omitempty"`
	Relation    []string `json:"relation,omitempty"`
	Coverage    []string `json:"coverage,omitempty"`
	Rights      []string `json:"rights,omitempty"`
}

// NewDublinCoreExtension creates a new DublinCoreExtension
// given the generic extension map for the "dc" prefix.
func NewDublinCoreExtension(extensions map[string][]Extension) *DublinCoreExtension {
	dc := &DublinCoreExtension{}
	dc.Title = parseTextArrayExtension("title", extensions)
	dc.Creator = parseTextArrayExtension("creator", extensions)
	dc.Author = parseTextArrayExtension("author", extensions)
	dc.Subject = parseTextArrayExtension("subject", extensions)
	dc.Description = parseTextArrayExtension("description", extensions)
	dc.Publisher = parseTextArrayExtension("publisher", extensions)
	dc.Contributor = parseTextArrayExtension("contributor", extensions)
	dc.Date = parseTextArrayExtension("date", extensions)
	dc.Type = parseTextArrayExtension("type", extensions)
	dc.Format = parseTextArrayExtension("format", extensions)
	dc.Identifier = parseTextArrayExtension("identifier", extensions)
	dc.Source = parseTextArrayExtension("source", extensions)
	dc.Language = parseTextArrayExtension("language", extensions)
	dc.Relation = parseTextArrayExtension("relation", extensions)
	dc.Coverage = parseTextArrayExtension("coverage", extensions)
	dc.Rights = parseTextArrayExtension("rights", extensions)
	return dc
}
