package j

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var JSXRules = func() Rules {
	rules := JavascriptRules.Clone()
	rules["jsx"] = []Rule{
		{`(<)([\w_\-]+)`, ByGroups(Punctuation, NameTag), Push("tag")},
		{`(<)(/)(\s*)([\w_\-]+)(\s*)(>)`, ByGroups(Punctuation, Punctuation, Text, NameTag, Text, Punctuation), nil},
	}
	rules["tag"] = []Rule{
		{`\s+`, Text, nil},
		{`([\w]+\s*)(=)(\s*)`, ByGroups(NameAttribute, Operator, Text), Push("attr")},
		{`[{}]+`, Punctuation, nil},
		{`[\w\.]+`, NameAttribute, nil},
		{`(/?)(\s*)(>)`, ByGroups(Punctuation, Text, Punctuation), Pop(1)},
	}
	rules["attr"] = []Rule{
		{`\s+`, Text, nil},
		{`".*?"`, String, Pop(1)},
		{`'.*?'`, String, Pop(1)},
		{`[^\s>]+`, String, Pop(1)},
	}

	rules["root"] = append([]Rule{Include("jsx")}, rules["root"]...)
	return rules
}()

// JSX lexer.
var JSX = internal.Register(MustNewLexer(
	&Config{
		Name:      "JSX",
		Aliases:   []string{"react"},
		Filenames: []string{"*.jsx", "*.react"},
		MimeTypes: []string{"text/jsx", "text/typescript-jsx"},
		DotAll:    true,
	},
	JSXRules,
))
