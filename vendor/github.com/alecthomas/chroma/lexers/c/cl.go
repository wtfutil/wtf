package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Common Lisp lexer.
var CommonLisp = internal.Register(MustNewLexer(
	&Config{
		Name:            "Common Lisp",
		Aliases:         []string{"common-lisp", "cl", "lisp"},
		Filenames:       []string{"*.cl", "*.lisp"},
		MimeTypes:       []string{"text/x-common-lisp"},
		CaseInsensitive: true,
	},
	Rules{
		"root": {
			Default(Push("body")),
		},
		"multiline-comment": {
			{`#\|`, CommentMultiline, Push()},
			{`\|#`, CommentMultiline, Pop(1)},
			{`[^|#]+`, CommentMultiline, nil},
			{`[|#]`, CommentMultiline, nil},
		},
		"commented-form": {
			{`\(`, CommentPreproc, Push()},
			{`\)`, CommentPreproc, Pop(1)},
			{`[^()]+`, CommentPreproc, nil},
		},
		"body": {
			{`\s+`, Text, nil},
			{`;.*$`, CommentSingle, nil},
			{`#\|`, CommentMultiline, Push("multiline-comment")},
			{`#\d*Y.*$`, CommentSpecial, nil},
			{`"(\\.|\\\n|[^"\\])*"`, LiteralString, nil},
			{`:(\|[^|]+\||(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~])(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~]|[#.:])*)`, LiteralStringSymbol, nil},
			{`::(\|[^|]+\||(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~])(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~]|[#.:])*)`, LiteralStringSymbol, nil},
			{`:#(\|[^|]+\||(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~])(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~]|[#.:])*)`, LiteralStringSymbol, nil},
			{`'(\|[^|]+\||(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~])(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~]|[#.:])*)`, LiteralStringSymbol, nil},
			{`'`, Operator, nil},
			{"`", Operator, nil},
			{"[-+]?\\d+\\.?(?=[ \"()\\'\\n,;`])", LiteralNumberInteger, nil},
			{"[-+]?\\d+/\\d+(?=[ \"()\\'\\n,;`])", LiteralNumber, nil},
			{"[-+]?(\\d*\\.\\d+([defls][-+]?\\d+)?|\\d+(\\.\\d*)?[defls][-+]?\\d+)(?=[ \"()\\'\\n,;`])", LiteralNumberFloat, nil},
			{"#\\\\.(?=[ \"()\\'\\n,;`])", LiteralStringChar, nil},
			{`#\\(\|[^|]+\||(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~])(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~]|[#.:])*)`, LiteralStringChar, nil},
			{`#\(`, Operator, Push("body")},
			{`#\d*\*[01]*`, LiteralOther, nil},
			{`#:(\|[^|]+\||(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~])(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~]|[#.:])*)`, LiteralStringSymbol, nil},
			{`#[.,]`, Operator, nil},
			{`#\'`, NameFunction, nil},
			{`#b[+-]?[01]+(/[01]+)?`, LiteralNumberBin, nil},
			{`#o[+-]?[0-7]+(/[0-7]+)?`, LiteralNumberOct, nil},
			{`#x[+-]?[0-9a-f]+(/[0-9a-f]+)?`, LiteralNumberHex, nil},
			{`#\d+r[+-]?[0-9a-z]+(/[0-9a-z]+)?`, LiteralNumber, nil},
			{`(#c)(\()`, ByGroups(LiteralNumber, Punctuation), Push("body")},
			{`(#\d+a)(\()`, ByGroups(LiteralOther, Punctuation), Push("body")},
			{`(#s)(\()`, ByGroups(LiteralOther, Punctuation), Push("body")},
			{`#p?"(\\.|[^"])*"`, LiteralOther, nil},
			{`#\d+=`, Operator, nil},
			{`#\d+#`, Operator, nil},
			{"#+nil(?=[ \"()\\'\\n,;`])\\s*\\(", CommentPreproc, Push("commented-form")},
			{`#[+-]`, Operator, nil},
			{`(,@|,|\.)`, Operator, nil},
			{"(t|nil)(?=[ \"()\\'\\n,;`])", NameConstant, nil},
			{`\*(\|[^|]+\||(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~])(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~]|[#.:])*)\*`, NameVariableGlobal, nil},
			{`(\|[^|]+\||(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~])(?:\\.|[\w!$%&*+-/<=>?@\[\]^{}~]|[#.:])*)`, NameVariable, nil},
			{`\(`, Punctuation, Push("body")},
			{`\)`, Punctuation, Pop(1)},
		},
	},
))
