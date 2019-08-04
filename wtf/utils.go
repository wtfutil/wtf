package wtf

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
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

var OpenFileUtil = "open"

// Init initializes global settings in the wtf package
func Init(openFileUtil string) {
	OpenFileUtil = openFileUtil
}

// CenterText takes a string and a width and pads the left and right of the string with
// empty spaces to ensure that the string is in the middle of the returned value
//
// Example:
//
//    x := CenterText("cat", 11)
//    > "    cat    "
//
func CenterText(str string, width int) string {
	if width < 0 {
		width = 0
	}

	return fmt.Sprintf("%[1]*s", -width, fmt.Sprintf("%[1]*s", (width+len(str))/2, str))
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

// NameFromEmail takes an email address and returns the part that comes before the @ symbol
//
// Example:
//
//    NameFromEmail("test_user@example.com")
//    > "Test_user"
//
func NameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	return strings.Title(strings.Replace(parts[0], ".", " ", -1))
}

// NamesFromEmails takes a slice of email addresses and returns a slice of the parts that
// come before the @ symbol
//
// Example:
//
//    NamesFromEmail("test_user@example.com", "other_user@example.com")
//    > []string{"Test_user", "Other_user"}
//
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
		filePath, _ := utils.ExpandHomeDir(path)
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

// ReadFileBytes reads the contents of a file and returns those contents as a slice of bytes
func ReadFileBytes(filePath string) ([]byte, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}, err
	}

	return fileData, nil
}

/* -------------------- Map Conversion -------------------- */

// MapToStrs takes a map of interfaces and returns a map of strings
func MapToStrs(aMap map[string]interface{}) map[string]string {
	results := make(map[string]string)

	for key, val := range aMap {
		results[key] = val.(string)
	}

	return results
}

/* -------------------- Slice Conversion -------------------- */

// ToInts takes a slice of interfaces and returns a slice of ints
func ToInts(slice []interface{}) []int {
	results := []int{}

	for _, val := range slice {
		results = append(results, val.(int))
	}

	return results
}

// ToStrs takes a slice of interfaces and returns a slice of strings
func ToStrs(slice []interface{}) []string {
	results := []string{}

	for _, val := range slice {
		switch val.(type) {
		case int:
			results = append(results, strconv.Itoa(val.(int)))
		case string:
			results = append(results, val.(string))
		}
	}

	return results
}

func HighlightableHelper(view *tview.TextView, input string, idx, offset int) string {
	fmtStr := fmt.Sprintf(`["%d"][""]`, idx)
	_, _, w, _ := view.GetInnerRect()
	fmtStr += input
	fmtStr += PadRow(offset, w+1)
	fmtStr += `[""]` + "\n"
	return fmtStr
}
