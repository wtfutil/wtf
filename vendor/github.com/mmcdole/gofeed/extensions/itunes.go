package ext

// ITunesFeedExtension is a set of extension
// fields for RSS feeds.
type ITunesFeedExtension struct {
	Author     string            `json:"author,omitempty"`
	Block      string            `json:"block,omitempty"`
	Categories []*ITunesCategory `json:"categories,omitempty"`
	Explicit   string            `json:"explicit,omitempty"`
	Keywords   string            `json:"keywords,omitempty"`
	Owner      *ITunesOwner      `json:"owner,omitempty"`
	Subtitle   string            `json:"subtitle,omitempty"`
	Summary    string            `json:"summary,omitempty"`
	Image      string            `json:"image,omitempty"`
	Complete   string            `json:"complete,omitempty"`
	NewFeedURL string            `json:"newFeedUrl,omitempty"`
}

// ITunesItemExtension is a set of extension
// fields for RSS items.
type ITunesItemExtension struct {
	Author            string `json:"author,omitempty"`
	Block             string `json:"block,omitempty"`
	Duration          string `json:"duration,omitempty"`
	Explicit          string `json:"explicit,omitempty"`
	Keywords          string `json:"keywords,omitempty"`
	Subtitle          string `json:"subtitle,omitempty"`
	Summary           string `json:"summary,omitempty"`
	Image             string `json:"image,omitempty"`
	IsClosedCaptioned string `json:"isClosedCaptioned,omitempty"`
	Order             string `json:"order,omitempty"`
}

// ITunesCategory is a category element for itunes feeds.
type ITunesCategory struct {
	Text        string          `json:"text,omitempty"`
	Subcategory *ITunesCategory `json:"subcategory,omitempty"`
}

// ITunesOwner is the owner of a particular itunes feed.
type ITunesOwner struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// NewITunesFeedExtension creates an ITunesFeedExtension given an
// extension map for the "itunes" key.
func NewITunesFeedExtension(extensions map[string][]Extension) *ITunesFeedExtension {
	feed := &ITunesFeedExtension{}
	feed.Author = parseTextExtension("author", extensions)
	feed.Block = parseTextExtension("block", extensions)
	feed.Explicit = parseTextExtension("explicit", extensions)
	feed.Keywords = parseTextExtension("keywords", extensions)
	feed.Subtitle = parseTextExtension("subtitle", extensions)
	feed.Summary = parseTextExtension("summary", extensions)
	feed.Image = parseImage(extensions)
	feed.Complete = parseTextExtension("complete", extensions)
	feed.NewFeedURL = parseTextExtension("new-feed-url", extensions)
	feed.Categories = parseCategories(extensions)
	feed.Owner = parseOwner(extensions)
	return feed
}

// NewITunesItemExtension creates an ITunesItemExtension given an
// extension map for the "itunes" key.
func NewITunesItemExtension(extensions map[string][]Extension) *ITunesItemExtension {
	entry := &ITunesItemExtension{}
	entry.Author = parseTextExtension("author", extensions)
	entry.Block = parseTextExtension("block", extensions)
	entry.Duration = parseTextExtension("duration", extensions)
	entry.Explicit = parseTextExtension("explicit", extensions)
	entry.Subtitle = parseTextExtension("subtitle", extensions)
	entry.Summary = parseTextExtension("summary", extensions)
	entry.Image = parseImage(extensions)
	entry.IsClosedCaptioned = parseTextExtension("isClosedCaptioned", extensions)
	entry.Order = parseTextExtension("order", extensions)
	return entry
}

func parseImage(extensions map[string][]Extension) (image string) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["image"]
	if !ok || len(matches) == 0 {
		return
	}

	image = matches[0].Attrs["href"]
	return
}

func parseOwner(extensions map[string][]Extension) (owner *ITunesOwner) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["owner"]
	if !ok || len(matches) == 0 {
		return
	}

	owner = &ITunesOwner{}
	if name, ok := matches[0].Children["name"]; ok {
		owner.Name = name[0].Value
	}
	if email, ok := matches[0].Children["email"]; ok {
		owner.Email = email[0].Value
	}
	return
}

func parseCategories(extensions map[string][]Extension) (categories []*ITunesCategory) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["category"]
	if !ok || len(matches) == 0 {
		return
	}

	categories = []*ITunesCategory{}
	for _, cat := range matches {
		c := &ITunesCategory{}
		if text, ok := cat.Attrs["text"]; ok {
			c.Text = text
		}

		if subs, ok := cat.Children["category"]; ok {
			s := &ITunesCategory{}
			if text, ok := subs[0].Attrs["text"]; ok {
				s.Text = text
			}
			c.Subcategory = s
		}
		categories = append(categories, c)
	}
	return
}
