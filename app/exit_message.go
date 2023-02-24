package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/logrusorgru/aurora/v4"
	"github.com/olebedev/config"
)

const exitMessageHeader = `
	      ____    __    ____ .___________. _______
	      \   \  /  \  /   / |           ||   ____|
	       \   \/    \/   /  ----|  |-----|  |__
	        \            /       |  |     |   __|
	         \    /\    /        |  |     |  |
	          \__/  \__/         |__|     |__|

    the personal information dashboard for your terminal
`

// DisplayExitMessage displays the onscreen exit message when the app quits
func (wtfApp *WtfApp) DisplayExitMessage() {
	exitMessageIsDisplayable := readDisplayableConfig(wtfApp.config)

	wtfApp.displayExitMsg(exitMessageIsDisplayable)
}

/* -------------------- Unexported Functions -------------------- */

func (wtfApp *WtfApp) displayExitMsg(exitMessageIsDisplayable bool) string {
	// If a sponsor or contributor and opt out of seeing the exit message, do not display it
	if (wtfApp.ghUser.IsContributor || wtfApp.ghUser.IsSponsor) && !exitMessageIsDisplayable {
		return ""
	}

	msgs := []string{}

	msgs = append(msgs, aurora.Magenta(exitMessageHeader).String())

	if wtfApp.ghUser.IsContributor {
		msgs = append(msgs, wtfApp.contributorThankYouMessage())
	}

	if wtfApp.ghUser.IsSponsor {
		msgs = append(msgs, wtfApp.sponsorThankYouMessage())
	}

	if !wtfApp.ghUser.IsContributor && !wtfApp.ghUser.IsSponsor {
		msgs = append(msgs, wtfApp.supportRequestMessage())
	}

	displayMsg := strings.Join(msgs, "\n")

	fmt.Println(displayMsg)

	return displayMsg
}

// readDisplayableConfig figures out whether or not the exit message should be displayed
// per the user's wishes. It allows contributors and sponsors to opt out of the exit message
func readDisplayableConfig(cfg *config.Config) bool {
	displayExitMsg := cfg.UBool("wtf.exitMessage.display", true)
	return displayExitMsg
}

// readGitHubAPIKey attempts to find a GitHub API key somewhere in the configuration file
func readGitHubAPIKey(cfg *config.Config) string {
	apiKey := cfg.UString("wtf.exitMessage.githubAPIKey", os.Getenv("WTF_GITHUB_TOKEN"))
	if apiKey != "" {
		return apiKey
	}

	moduleConfig, err := cfg.Get("wtf.mods.github")
	if err != nil {
		return ""
	}

	return moduleConfig.UString("apiKey", "")
}

/* -------------------- Messaging -------------------- */

func (wtfApp *WtfApp) contributorThankYouMessage() string {
	str := "    On behalf of all the users of WTF, thank you for contributing to the source code."
	str += fmt.Sprintf(" %s", aurora.Green("\n\n    You rock."))

	return str
}

func (wtfApp *WtfApp) sponsorThankYouMessage() string {
	str := "    Your sponsorship of WTF makes a difference. Thank you for sponsoring and supporting WTF."
	str += fmt.Sprintf(" %s", aurora.Green("\n\n    You're awesome."))

	return str
}

func (wtfApp *WtfApp) supportRequestMessage() string {
	str := "    The development and maintenance of WTF is supported by sponsorships.\n"
	str += fmt.Sprintf("    Sponsor the development of WTF at %s\n", aurora.Green("https://github.com/sponsors/senorprogrammer"))

	return str
}
