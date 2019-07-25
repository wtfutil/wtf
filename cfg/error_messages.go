package cfg

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

func displayDefaultConfigWriteError(err error) {
	fmt.Printf("\n%s Could not write the default configuration.\n", aurora.Red("ERROR:"))
	fmt.Println()
	displayError(err)
}

func displayXdgConfigDirCreateError(err error) {
	fmt.Printf("\n%s Could not create the '%s' directory.\n", aurora.Red("ERROR:"), aurora.Yellow(XdgConfigDir))
	fmt.Println()
	displayError(err)
}

func displayWtfConfigDirCreateError(err error) {
	fmt.Printf("\n%s Could not create the '%s' directory.\n", aurora.Red("ERROR:"), aurora.Yellow(WtfConfigDirV2))
	fmt.Println()
	displayError(err)
}

func displayWtfConfigFileLoadError(err error) {
	fmt.Printf("\n%s Could not load '%s'.\n", aurora.Red("ERROR:"), aurora.Yellow(WtfConfigFile))
	fmt.Println()
	fmt.Println("This could mean one of two things:")
	fmt.Println()
	fmt.Printf("    1. Your %s file is missing. Check in %s to see if %s is there.\n", aurora.Yellow(WtfConfigFile), aurora.Yellow("~/.config/wtf/"), aurora.Yellow(WtfConfigFile))
	fmt.Printf("    2. Your %s file has a syntax error. Try running it through http://www.yamllint.com to check for errors.\n", aurora.Yellow(WtfConfigFile))
	fmt.Println()
	displayError(err)
}

func displayWtfCustomConfigFileLoadError(err error) {
	fmt.Printf("\n%s Could not load '%s'.\n", aurora.Red("ERROR:"), aurora.Yellow(WtfConfigFile))
	fmt.Println()
	fmt.Println("This could mean one of two things:")
	fmt.Println()
	fmt.Println("    1. That file doesn't exist.")
	fmt.Println("    2. That file has a YAML syntax error. Try running it through http://www.yamllint.com to check for errors.")
	fmt.Println()
	displayError(err)
}
