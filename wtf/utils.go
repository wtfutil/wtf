package wtf

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

const SimpleDateFormat = "Jan 2"
const SimpleTimeFormat = "15:04 MST"
const MinimumTimeFormat = "15:04"
const FullDateFormat = "Monday, Jan 2"
const FriendlyDateFormat = "Mon, Jan 2"
const FriendlyDateTimeFormat = "Mon, Jan 2, 15:04"
const TimestampFormat = "2006-01-02T15:04:05-0700"

var OpenFileUtil = "open"

func CenterText(str string, width int) string {
	if width < 0 {
		width = 0
	}

	return fmt.Sprintf("%[1]*s", -width, fmt.Sprintf("%[1]*s", (width+len(str))/2, str))
}

func ExecuteCommand(cmd *exec.Cmd) string {
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

func Exclude(strs []string, val string) bool {
	for _, str := range strs {
		if val == str {
			return false
		}
	}
	return true
}

func FindMatch(pattern string, data string) [][]string {
	r := regexp.MustCompile(pattern)
	return r.FindAllStringSubmatch(data, -1)
}

func NameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	return strings.Title(strings.Replace(parts[0], ".", " ", -1))
}

func NamesFromEmails(emails []string) []string {
	names := []string{}

	for _, email := range emails {
		names = append(names, NameFromEmail(email))
	}

	return names
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

// PadRow returns a padding for a row to make it the full width of the containing widget.
// Useful for ensurig row highlighting spans the full width (I suspect tcell has a better
// way to do this, but I haven't yet found it)
func PadRow(offset int, max int) string {
	padSize := max - offset
	if padSize < 0 {
		padSize = 0
	}

	return strings.Repeat(" ", padSize)
}

func ReadFileBytes(filePath string) ([]byte, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, err
	}

	return fileData, nil
}

/* -------------------- Map Conversion -------------------- */

func MapToStrs(aMap map[string]interface{}) map[string]string {
	results := make(map[string]string)

	for key, val := range aMap {
		results[key] = val.(string)
	}

	return results
}

/* -------------------- Slice Conversion -------------------- */

func ToInts(slice []interface{}) []int {
	results := []int{}

	for _, val := range slice {
		results = append(results, val.(int))
	}

	return results
}

func ToStrs(slice []interface{}) []string {
	results := []string{}

	for _, val := range slice {
		results = append(results, val.(string))
	}

	return results
}
