package utils

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

const (
	SimpleDateFormat  = "Jan 2"
	SimpleTimeFormat  = "15:04 MST"
	MinimumTimeFormat = "15:04"

	FullDateFormat         = "Monday, Jan 2"
	FriendlyDateFormat     = "Mon, Jan 2"
	FriendlyDateTimeFormat = "Mon, Jan 2, 15:04"

	TimestampFormat = "2006-01-02T15:04:05-0700"
)

// ExecuteCommand executes an external command on the local machine as the current user
func ExecuteCommand(cmd *exec.Cmd) string {
	if cmd == nil {
		return ""
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	var str string
	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += string(b)
	}

	err = cmd.Wait()
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	return str
}

// Exclude takes a slice of strings and a target string and returns the contents of the original
// slice of strings without the target string in it
//
// Example:
//
//    x := Exclude([]string{"cat", "dog", "rat"}, "dog")
//    > []string{"cat", "rat"}
//
func Exclude(strs []string, val string) bool {
	for _, str := range strs {
		if val == str {
			return false
		}
	}
	return true
}

// FindMatch takes a regex pattern and a string of data and returns back all the matches
// in that string
func FindMatch(pattern string, data string) [][]string {
	r := regexp.MustCompile(pattern)
	return r.FindAllStringSubmatch(data, -1)
}

// OpenFile opens the file defined in `path` via the operating system
func OpenFile(path string) {
	if (strings.HasPrefix(path, "http://")) || (strings.HasPrefix(path, "https://")) {
		switch runtime.GOOS {
		case "linux":
			exec.Command("xdg-open", path).Start()
		case "windows":
			exec.Command("rundll32", "url.dll,FileProtocolHandler", path).Start()
		case "darwin":
			exec.Command("open", path).Start()
		default:
		}
	} else {
		filePath, _ := ExpandHomeDir(path)
		cmd := exec.Command(OpenFileUtil, filePath)
		ExecuteCommand(cmd)
	}
}

// ReadFileBytes reads the contents of a file and returns those contents as a slice of bytes
func ReadFileBytes(filePath string) ([]byte, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, err
	}

	return fileData, nil
}
