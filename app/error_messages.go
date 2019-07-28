package app

// This file contains the error messages that get written to the terminal when
// something goes wrong with the configuration process.
//
// As a general rule, if one of these has to be shown the app should then die
// via os.Exit(1)

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

/* -------------------- Unexported Functions -------------------- */

func displayError(err error) {
	fmt.Printf("%s %s\n\n", aurora.Red("Error:"), err.Error())
}
