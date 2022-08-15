package hibp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	apiURL            = "https://haveibeenpwned.com/api/v3/breachedaccount/"
	clientTimeoutSecs = 2
	userAgent         = "WTFUtil"
)

type hibpError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (widget *Widget) fullURL(account string, truncated bool) string {
	truncStr := "false"
	if truncated {
		truncStr = "true"
	}

	return apiURL + account + fmt.Sprintf("?truncateResponse=%s", truncStr)
}

func (widget *Widget) fetchForAccount(account string, since string) (*Status, error) {
	if account == "" {
		return nil, nil
	}

	hibpClient := http.Client{
		Timeout: time.Second * clientTimeoutSecs,
	}

	asTruncated := true
	if since != "" {
		asTruncated = false
	}

	request, err := http.NewRequest(http.MethodGet, widget.fullURL(account, asTruncated), http.NoBody)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("hibp-api-key", widget.settings.apiKey)

	response, getErr := hibpClient.Do(request)
	if getErr != nil {
		return nil, err
	}

	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return nil, err
	}

	hibpErr := widget.validateHTTPResponse(response.StatusCode, body)
	if hibpErr != nil {
		return nil, errors.New(hibpErr.Message)
	}

	stat, err := widget.parseResponseBody(account, body)
	if err != nil {
		return nil, err
	}

	return stat, nil
}

func (widget *Widget) parseResponseBody(account string, body []byte) (*Status, error) {
	breaches := []Breach{}
	stat := NewStatus(account, breaches)

	if len(body) == 0 {
		// If the body is empty then there's no breaches
		return stat, nil
	}

	jsonErr := json.Unmarshal(body, &breaches)
	if jsonErr != nil {
		return stat, jsonErr
	}

	breaches = widget.filterBreaches(breaches)
	stat.Breaches = breaches

	return stat, nil
}

func (widget *Widget) filterBreaches(breaches []Breach) []Breach {
	// If there's no valid since value in the settings, there's no point in trying to filter
	// the breaches on that value, they'll all pass
	if !widget.settings.HasSince() {
		return breaches
	}

	sinceDate, err := widget.settings.SinceDate()
	if err != nil {
		return breaches
	}

	latestBreaches := []Breach{}

	for _, breach := range breaches {
		breachDate, err := breach.BreachDate()
		if err != nil {
			// Append the erring breach here because a failing breach date doesn't mean that
			// the breach itself isn't applicable. The date could be missing or malformed,
			// in which case we err on the side of caution and assume that the breach is valid
			latestBreaches = append(latestBreaches, breach)
			continue
		}

		if breachDate.After(sinceDate) {
			latestBreaches = append(latestBreaches, breach)
		}
	}

	return latestBreaches
}

func (widget *Widget) validateHTTPResponse(responseCode int, body []byte) *hibpError {
	hibpErr := &hibpError{}

	switch responseCode {
	case 401, 402:
		err := json.Unmarshal(body, hibpErr)
		if err != nil {
			return nil
		}
	default:
		hibpErr = nil
	}

	return hibpErr
}
