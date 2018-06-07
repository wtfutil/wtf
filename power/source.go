// +build !linux

package power

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

const SingleQuotesRegExp = "'(.*)'"

// powerSource returns the name of the current power source, probably one of
// "AC Power" or "Battery Power"
func powerSource() string {
	cmd := exec.Command("pmset", []string{"-g", "ps"}...)
	result := wtf.ExecuteCommand(cmd)

	r, _ := regexp.Compile(SingleQuotesRegExp)

	source := r.FindString(result)
	source = strings.Replace(source, "'", "", -1)

	return source
}
