package app

import (
	"strings"
	"testing"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/support"
	"gotest.tools/assert"
)

func Test_displayExitMessage(t *testing.T) {
	tests := []struct {
		name          string
		isDisplayable bool
		isContributor bool
		isSponsor     bool
		compareWith   string
		expected      string
	}{
		{
			name:          "when not displayable",
			isDisplayable: false,
			isContributor: true,
			isSponsor:     true,
			compareWith:   "equals",
			expected:      "",
		},
		{
			name:          "when contributor",
			isDisplayable: true,
			isContributor: true,
			compareWith:   "contains",
			expected:      "thank you for contributing",
		},
		{
			name:          "when sponsor",
			isDisplayable: true,
			isSponsor:     true,
			compareWith:   "contains",
			expected:      "Thank you for sponsoring",
		},
		{
			name:          "when user",
			isDisplayable: true,
			isContributor: false,
			isSponsor:     false,
			compareWith:   "contains",
			expected:      "supported by sponsorships",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appMan := NewAppManager(&config.Config{}, tview.NewApplication())
			appMan.ghUser = &support.GitHubUser{
				IsContributor: tt.isContributor,
				IsSponsor:     tt.isSponsor,
			}

			actual := appMan.displayExitMsg(tt.isDisplayable)

			if tt.compareWith == "equals" {
				assert.Equal(t, actual, tt.expected)
			}

			if tt.compareWith == "contains" {
				assert.Equal(t, true, strings.Contains(actual, tt.expected))
			}
		})
	}
}
