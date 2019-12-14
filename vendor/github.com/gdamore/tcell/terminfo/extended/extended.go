// Copyright 2019 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package extended contains an extended set of terminal descriptions.
// Applications desiring to have a better chance of Just Working by
// default should include this package.  This will significantly increase
// the size of the program.
package extended

import (
	// The following imports just register themselves --
	// these are the terminal types we aggregate in this package.
	_ "github.com/gdamore/tcell/terminfo/a/adm3a"
	_ "github.com/gdamore/tcell/terminfo/a/aixterm"
	_ "github.com/gdamore/tcell/terminfo/a/alacritty"
	_ "github.com/gdamore/tcell/terminfo/a/ansi"
	_ "github.com/gdamore/tcell/terminfo/a/aterm"
	_ "github.com/gdamore/tcell/terminfo/b/beterm"
	_ "github.com/gdamore/tcell/terminfo/b/bsdos_pc"
	_ "github.com/gdamore/tcell/terminfo/c/cygwin"
	_ "github.com/gdamore/tcell/terminfo/d/d200"
	_ "github.com/gdamore/tcell/terminfo/d/d210"
	_ "github.com/gdamore/tcell/terminfo/d/dtterm"
	_ "github.com/gdamore/tcell/terminfo/e/emacs"
	_ "github.com/gdamore/tcell/terminfo/e/eterm"
	_ "github.com/gdamore/tcell/terminfo/g/gnome"
	_ "github.com/gdamore/tcell/terminfo/h/hpterm"
	_ "github.com/gdamore/tcell/terminfo/h/hz1500"
	_ "github.com/gdamore/tcell/terminfo/k/konsole"
	_ "github.com/gdamore/tcell/terminfo/k/kterm"
	_ "github.com/gdamore/tcell/terminfo/l/linux"
	_ "github.com/gdamore/tcell/terminfo/p/pcansi"
	_ "github.com/gdamore/tcell/terminfo/r/rxvt"
	_ "github.com/gdamore/tcell/terminfo/s/screen"
	_ "github.com/gdamore/tcell/terminfo/s/simpleterm"
	_ "github.com/gdamore/tcell/terminfo/s/sun"
	_ "github.com/gdamore/tcell/terminfo/t/termite"
	_ "github.com/gdamore/tcell/terminfo/t/tvi910"
	_ "github.com/gdamore/tcell/terminfo/t/tvi912"
	_ "github.com/gdamore/tcell/terminfo/t/tvi921"
	_ "github.com/gdamore/tcell/terminfo/t/tvi925"
	_ "github.com/gdamore/tcell/terminfo/t/tvi950"
	_ "github.com/gdamore/tcell/terminfo/t/tvi970"
	_ "github.com/gdamore/tcell/terminfo/v/vt100"
	_ "github.com/gdamore/tcell/terminfo/v/vt102"
	_ "github.com/gdamore/tcell/terminfo/v/vt220"
	_ "github.com/gdamore/tcell/terminfo/v/vt320"
	_ "github.com/gdamore/tcell/terminfo/v/vt400"
	_ "github.com/gdamore/tcell/terminfo/v/vt420"
	_ "github.com/gdamore/tcell/terminfo/v/vt52"
	_ "github.com/gdamore/tcell/terminfo/w/wy50"
	_ "github.com/gdamore/tcell/terminfo/w/wy60"
	_ "github.com/gdamore/tcell/terminfo/w/wy99_ansi"
	_ "github.com/gdamore/tcell/terminfo/x/xfce"
	_ "github.com/gdamore/tcell/terminfo/x/xnuppc"
	_ "github.com/gdamore/tcell/terminfo/x/xterm"
	_ "github.com/gdamore/tcell/terminfo/x/xterm_kitty"
)
