//go:build !linux && !freebsd

package power

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

const SingleQuotesRegExp = "'(.*)'"

// powerSource returns the name of the current power source, probably one of
// "AC Power" or "Battery Power"
func powerSource() string {
	cmd := exec.Command("pmset", []string{"-g", "ps"}...)
	result := utils.ExecuteCommand(cmd)

	r, _ := regexp.Compile(SingleQuotesRegExp)

	source := r.FindString(result)
	source = strings.Replace(source, "'", "", -1)

	return source
}
