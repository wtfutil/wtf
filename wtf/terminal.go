package wtf

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora/v4"
	"github.com/olebedev/config"
)

// SetTerminal sets the TERM environment variable, defaulting to whatever the OS
// has as the current value if none is specified.
// See https://www.gnu.org/software/gettext/manual/html_node/The-TERM-variable.html for
// more details.
func SetTerminal(config *config.Config) {
	term := config.UString("wtf.term", os.Getenv("TERM"))
	err := os.Setenv("TERM", term)
	if err != nil {
		fmt.Printf("\n%s Failed to set $TERM to %s.\n", aurora.Red("ERROR"), aurora.Yellow(term))
		os.Exit(1)
	}
}
