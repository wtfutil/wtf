package t

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Terraform lexer.
var Terraform = internal.Register(MustNewLexer(
	&Config{
		Name:      "Terraform",
		Aliases:   []string{"terraform", "tf"},
		Filenames: []string{"*.tf"},
		MimeTypes: []string{"application/x-tf", "application/x-terraform"},
	},
	Rules{
		"root": {
			Include("string"),
			Include("punctuation"),
			Include("curly"),
			Include("basic"),
			Include("whitespace"),
			{`[0-9]+`, LiteralNumber, nil},
		},
		"basic": {
			{Words(`\b`, `\b`, `true`, `false`), KeywordType, nil},
			{`\s*/\*`, CommentMultiline, Push("comment")},
			{`\s*#.*\n`, CommentSingle, nil},
			{`(.*?)(\s*)(=)`, ByGroups(NameAttribute, Text, Operator), nil},
			{Words(`\b`, `\b`, `variable`, `resource`, `provider`, `provisioner`, `module`), KeywordReserved, Push("function")},
			{Words(`\b`, `\b`, `ingress`, `egress`, `listener`, `default`, `connection`, `alias`), KeywordDeclaration, nil},
			{`\$\{`, LiteralStringInterpol, Push("var_builtin")},
		},
		"function": {
			{`(\s+)(".*")(\s+)`, ByGroups(Text, LiteralString, Text), nil},
			Include("punctuation"),
			Include("curly"),
		},
		"var_builtin": {
			{`\$\{`, LiteralStringInterpol, Push()},
			{Words(`\b`, `\b`, `concat`, `file`, `join`, `lookup`, `element`), NameBuiltin, nil},
			Include("string"),
			Include("punctuation"),
			{`\s+`, Text, nil},
			{`\}`, LiteralStringInterpol, Pop(1)},
		},
		"string": {
			{`(".*")`, ByGroups(LiteralStringDouble), nil},
		},
		"punctuation": {
			{`[\[\](),.]`, Punctuation, nil},
		},
		"curly": {
			{`\{`, TextPunctuation, nil},
			{`\}`, TextPunctuation, nil},
		},
		"comment": {
			{`[^*/]`, CommentMultiline, nil},
			{`/\*`, CommentMultiline, Push()},
			{`\*/`, CommentMultiline, Pop(1)},
			{`[*/]`, CommentMultiline, nil},
		},
		"whitespace": {
			{`\n`, Text, nil},
			{`\s+`, Text, nil},
			{`\\\n`, Text, nil},
		},
	},
))
