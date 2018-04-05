package wtf

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
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

func Today() string {
	localNow := time.Now().Local()
	return localNow.Format("2006-01-02")
}

func UnixTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}
