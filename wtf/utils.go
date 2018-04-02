package wtf

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func CenterText(str string, width int) string {
	return fmt.Sprintf("%[1]*s", -width, fmt.Sprintf("%[1]*s", (width+len(str))/2, str))
}

func ExecuteCommand(cmd *exec.Cmd) string {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Sprintf("A: %v\n", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Sprintf("B: %v\n", err)
	}

	var str string
	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += string(b)
	}

	cmd.Wait()

	return str
}
