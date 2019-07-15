package ext

// Extensions is the generic extension map for Feeds and Items.
// The first map is for the element namespace prefix (e.g., itunes).
// The second map is for the element name (e.g., author).
type Extensions map[string]map[string][]Extension

// Extension represents a single XML element that was in a non
// default namespace in a Feed or Item/Entry.
type Extension struct {
	Name     string                 `json:"name"`
	Value    string                 `json:"value"`
	Attrs    map[string]string      `json:"attrs"`
	Children map[string][]Extension `json:"children"`
}

func parseTextExtension(name string, extensions map[string][]Extension) (value string) {
	if extensions == nil {
		return
	}

	matches, ok := extensions[name]
	if !ok || len(matches) == 0 {
		return
	}

	match := matches[0]
	return match.Value
}

func parseTextArrayExtension(name string, extensions map[string][]Extension) (values []string) {
	if extensions == nil {
		return
	}

	matches, ok := extensions[name]
	if !ok || len(matches) == 0 {
		return
	}

	values = []string{}
	for _, m := range matches {
		values = append(values, m.Value)
	}
	return
}
