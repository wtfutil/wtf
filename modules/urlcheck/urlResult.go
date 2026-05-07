package urlcheck

import (
	"net/url"
)

const InvalidResultCode = 999

// Collect useful properties of each given URL
type urlResult struct {
	Url           string
	ResultCode    int
	ResultMessage string
	IsValid       bool
}

// DisplayURL returns the URL with any Basic Auth password redacted.
func (ur urlResult) DisplayURL() string {
	u, err := url.Parse(ur.Url)
	if err != nil || u.User == nil {
		return ur.Url
	}

	username := u.User.Username()
	if _, hasPassword := u.User.Password(); !hasPassword {
		return ur.Url
	}

	u.User = url.UserPassword(username, "xxxxx")
	return u.String()
}

// Create a UrlResult instance from an urls occurence in the settings
func newUrlResult(urlString string) *urlResult {

	uResult := urlResult{
		Url: urlString,
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		uResult.ResultMessage = err.Error()
		uResult.ResultCode = InvalidResultCode
		uResult.IsValid = false
		return &uResult
	}

	uResult.IsValid = true
	return &uResult
}
