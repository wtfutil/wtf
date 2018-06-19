package trello

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type notFoundError interface {
	IsNotFound() bool
}

type rateLimitError interface {
	IsRateLimit() bool
}

type permissionDeniedError interface {
	IsPermissionDenied() bool
}

type httpClientError struct {
	msg  string
	code int
}

func makeHttpClientError(url string, resp *http.Response) error {

	body, _ := ioutil.ReadAll(resp.Body)
	msg := fmt.Sprintf("HTTP request failure on %s:\n%d: %s", url, resp.StatusCode, string(body))

	return &httpClientError{
		msg:  msg,
		code: resp.StatusCode,
	}
}

func (e *httpClientError) Error() string            { return e.msg }
func (e *httpClientError) IsRateLimit() bool        { return e.code == 429 }
func (e *httpClientError) IsNotFound() bool         { return e.code == 404 }
func (e *httpClientError) IsPermissionDenied() bool { return e.code == 401 }

func IsRateLimit(err error) bool {
	re, ok := err.(rateLimitError)
	return ok && re.IsRateLimit()
}

func IsNotFound(err error) bool {
	nf, ok := err.(notFoundError)
	return ok && nf.IsNotFound()
}

func IsPermissionDenied(err error) bool {
	pd, ok := err.(permissionDeniedError)
	return ok && pd.IsPermissionDenied()
}
