package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Powershell lexer.
var Powershell = internal.Register(MustNewLexer(
	&Config{
		Name:            "PowerShell",
		Aliases:         []string{"powershell", "posh", "ps1", "psm1"},
		Filenames:       []string{"*.ps1", "*.psm1"},
		MimeTypes:       []string{"text/x-powershell"},
		DotAll:          true,
		CaseInsensitive: true,
	},
	Rules{
		"root": {
			{`\(`, Punctuation, Push("child")},
			{`\s+`, Text, nil},
			{`^(\s*#[#\s]*)(\.(?:component|description|example|externalhelp|forwardhelpcategory|forwardhelptargetname|functionality|inputs|link|notes|outputs|parameter|remotehelprunspace|role|synopsis))([^\n]*$)`, ByGroups(Comment, LiteralStringDoc, Comment), nil},
			{`#[^\n]*?$`, Comment, nil},
			{`(&lt;|<)#`, CommentMultiline, Push("multline")},
			{`@"\n`, LiteralStringHeredoc, Push("heredoc-double")},
			{`@'\n.*?\n'@`, LiteralStringHeredoc, nil},
			{"`[\\'\"$@-]", Punctuation, nil},
			{`"`, LiteralStringDouble, Push("string")},
			{`'([^']|'')*'`, LiteralStringSingle, nil},
			{`(\$|@@|@)((global|script|private|env):)?\w+`, NameVariable, nil},
			{`(while|validateset|validaterange|validatepattern|validatelength|validatecount|until|trap|switch|return|ref|process|param|parameter|in|if|global:|function|foreach|for|finally|filter|end|elseif|else|dynamicparam|do|default|continue|cmdletbinding|break|begin|alias|\?|%|#script|#private|#local|#global|mandatory|parametersetname|position|valuefrompipeline|valuefrompipelinebypropertyname|valuefromremainingarguments|helpmessage|try|catch|throw)\b`, Keyword, nil},
			{`-(and|as|band|bnot|bor|bxor|casesensitive|ccontains|ceq|cge|cgt|cle|clike|clt|cmatch|cne|cnotcontains|cnotlike|cnotmatch|contains|creplace|eq|exact|f|file|ge|gt|icontains|ieq|ige|igt|ile|ilike|ilt|imatch|ine|inotcontains|inotlike|inotmatch|ireplace|is|isnot|le|like|lt|match|ne|not|notcontains|notlike|notmatch|or|regex|replace|wildcard)\b`, Operator, nil},
			{`(write|where|wait|use|update|unregister|undo|trace|test|tee|take|suspend|stop|start|split|sort|skip|show|set|send|select|scroll|resume|restore|restart|resolve|resize|reset|rename|remove|register|receive|read|push|pop|ping|out|new|move|measure|limit|join|invoke|import|group|get|format|foreach|export|expand|exit|enter|enable|disconnect|disable|debug|cxnew|copy|convertto|convertfrom|convert|connect|complete|compare|clear|checkpoint|aggregate|add)-[a-z_]\w*\b`, NameBuiltin, nil},
			{"\\[[a-z_\\[][\\w. `,\\[\\]]*\\]", NameConstant, nil},
			{`-[a-z_]\w*`, Name, nil},
			{`\w+`, Name, nil},
			{"[.,;@{}\\[\\]$()=+*/\\\\&%!~?^`|<>-]|::", Punctuation, nil},
		},
		"child": {
			{`\)`, Punctuation, Pop(1)},
			Include("root"),
		},
		"multline": {
			{`[^#&.]+`, CommentMultiline, nil},
			{`#(>|&gt;)`, CommentMultiline, Pop(1)},
			{`\.(component|description|example|externalhelp|forwardhelpcategory|forwardhelptargetname|functionality|inputs|link|notes|outputs|parameter|remotehelprunspace|role|synopsis)`, LiteralStringDoc, nil},
			{`[#&.]`, CommentMultiline, nil},
		},
		"string": {
			{"`[0abfnrtv'\\\"$`]", LiteralStringEscape, nil},
			{"[^$`\"]+", LiteralStringDouble, nil},
			{`\$\(`, Punctuation, Push("child")},
			{`""`, LiteralStringDouble, nil},
			{"[`$]", LiteralStringDouble, nil},
			{`"`, LiteralStringDouble, Pop(1)},
		},
		"heredoc-double": {
			{`\n"@`, LiteralStringHeredoc, Pop(1)},
			{`\$\(`, Punctuation, Push("child")},
			{`[^@\n]+"]`, LiteralStringHeredoc, nil},
			{`.`, LiteralStringHeredoc, nil},
		},
	},
))
