package devto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// nonSuccessfulResponse indicates that the status code for
// the HTTP response passed in was not one of the "success"
// (2XX) status codes
func nonSuccessfulResponse(res *http.Response) bool { return res.StatusCode/100 != 2 }

// attempt to deserialize the error response; if it succeeds,
// the error will be an ErrorResponse, otherwise it will be
// an error indicating that the error response could not be
// deserialized.
func unmarshalErrorResponse(res *http.Response) error {
	var e ErrorResponse
	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		return fmt.Errorf(
			`unexpected error deserializing %d response: "%v"`,
			res.StatusCode,
			err,
		)
	}
	return &e
}

func decodeResponse(r *http.Response) []byte {
	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte("")
	}
	defer r.Body.Close()
	return c
}
