package styles

import (
	"github.com/alecthomas/chroma"
)

// Rrt style.
var Rrt = Register(chroma.MustNewStyle("rrt", chroma.StyleEntries{
	chroma.Comment:        "#00ff00",
	chroma.NameFunction:   "#ffff00",
	chroma.NameVariable:   "#eedd82",
	chroma.NameConstant:   "#7fffd4",
	chroma.Keyword:        "#ff0000",
	chroma.CommentPreproc: "#e5e5e5",
	chroma.LiteralString:  "#87ceeb",
	chroma.KeywordType:    "#ee82ee",
	chroma.Background:     " bg:#000000",
}))
