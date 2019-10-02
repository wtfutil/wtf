package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

const (
	SimpleDateFormat    = "Jan 2"
	SimpleTimeFormat    = "15:04 MST"
	MinimumTimeFormat12 = "3:04 PM"
	MinimumTimeFormat24 = "15:04"

	FullDateFormat         = "Monday, Jan 2"
	FriendlyDateFormat     = "Mon, Jan 2"
	FriendlyDateTimeFormat = "Mon, Jan 2, 15:04"

	TimestampFormat = "2006-01-02T15:04:05-0700"
)

// DoesNotInclude takes a slice of strings and a target string and returns
// TRUE if the slice does not include the target, FALSE if it does
//
// Example:
//
//    x := DoesNotInclude([]string{"cat", "dog", "rat"}, "dog")
//    > false
//
//    x := DoesNotInclude([]string{"cat", "dog", "rat"}, "pig")
//    > true
//
func DoesNotInclude(strs []string, val string) bool {
	return !Includes(strs, val)
}

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

// FindMatch takes a regex pattern and a string of data and returns back all the matches
// in that string
func FindMatch(pattern string, data string) [][]string {
	r := regexp.MustCompile(pattern)
	return r.FindAllStringSubmatch(data, -1)
}

// Includes takes a slice of strings and a target string and returns
// TRUE if the slice includes the target, FALSE if it does not
//
// Example:
//
//    x := Includes([]string{"cat", "dog", "rat"}, "dog")
//    > true
//
//    x := Includes([]string{"cat", "dog", "rat"}, "pig")
//    > false
//
func Includes(strs []string, val string) bool {
	for _, str := range strs {
		if val == str {
			return true
		}
	}
	return false
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
			// for the BSDs
			exec.Command("xdg-open", path).Start()
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

// ParseJson is a standard JSON reader from text
func ParseJson(obj interface{}, text io.Reader) error {
	jsonStream, err := ioutil.ReadAll(text)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonStream))

	for {
		if err := decoder.Decode(obj); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}
