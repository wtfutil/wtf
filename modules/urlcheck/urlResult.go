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
