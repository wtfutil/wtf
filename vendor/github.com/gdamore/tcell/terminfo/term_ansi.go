// Generated automatically.  DO NOT HAND-EDIT.

package terminfo

func init() {
	// ansi/pc-term compatible with color
	AddTerminfo(&Terminfo{
		Name:         "ansi",
		Columns:      80,
		Lines:        24,
		Colors:       8,
		Bell:         "\a",
		Clear:        "\x1b[H\x1b[J",
		AttrOff:      "\x1b[0;10m",
		Underline:    "\x1b[4m",
		Bold:         "\x1b[1m",
		Blink:        "\x1b[5m",
		Reverse:      "\x1b[7m",
		SetFg:        "\x1b[3%p1%dm",
		SetBg:        "\x1b[4%p1%dm",
		SetFgBg:      "\x1b[3%p1%d;4%p2%dm",
		PadChar:      "\x00",
		AltChars:     "+\x0020,\x0021-\x0030.\x190333`\x0004a261f370g361h260j331k277l332m300n305o~p304q304r304s_t303u264v301w302x263y363z362{343|330}234~376",
		EnterAcs:     "\x1b[11m",
		ExitAcs:      "\x1b[10m",
		SetCursor:    "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:  "\x1b[D",
		CursorUp1:    "\x1b[A",
		KeyUp:        "\x1b[A",
		KeyDown:      "\x1b[B",
		KeyRight:     "\x1b[C",
		KeyLeft:      "\x1b[D",
		KeyInsert:    "\x1b[L",
		KeyBackspace: "\b",
		KeyHome:      "\x1b[H",
		KeyBacktab:   "\x1b[Z",
	})
}
