// Generated automatically.  DO NOT HAND-EDIT.

package terminfo

func init() {
	// gnu emacs term.el terminal emulation
	AddTerminfo(&Terminfo{
		Name:        "eterm",
		Columns:     80,
		Lines:       24,
		Bell:        "\a",
		Clear:       "\x1b[H\x1b[J",
		EnterCA:     "\x1b7\x1b[?47h",
		ExitCA:      "\x1b[2J\x1b[?47l\x1b8",
		AttrOff:     "\x1b[m",
		Underline:   "\x1b[4m",
		Bold:        "\x1b[1m",
		Reverse:     "\x1b[7m",
		PadChar:     "\x00",
		SetCursor:   "\x1b[%i%p1%d;%p2%dH",
		CursorBack1: "\b",
		CursorUp1:   "\x1b[A",
	})
}
