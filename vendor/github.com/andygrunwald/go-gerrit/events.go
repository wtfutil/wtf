package gerrit

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"time"
)

// PatchSet contains detailed information about a specific patch set.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/json.html#patchSet
type PatchSet struct {
	Number    Number      `json:"number"`
	Revision  string      `json:"revision"`
	Parents   []string    `json:"parents"`
	Ref       string      `json:"ref"`
	Uploader  AccountInfo `json:"uploader"`
	Author    AccountInfo `json:"author"`
	CreatedOn int         `json:"createdOn"`
	IsDraft   bool        `json:"isDraft"`
	Kind      string      `json:"kind"`
}

// RefUpdate contains data about a reference update.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/json.html#refUpdate
type RefUpdate struct {
	OldRev  string `json:"oldRev"`
	NewRev  string `json:"newRev"`
	RefName string `json:"refName"`
	Project string `json:"project"`
}

// EventInfo contains information about an event emitted by Gerrit.  This
// structure can be used either when parsing streamed events or when reading
// the output of the events-log plugin.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/cmd-stream-events.html#events
type EventInfo struct {
	Type           string        `json:"type"`
	Change         ChangeInfo    `json:"change,omitempty"`
	ChangeKey      ChangeInfo    `json:"changeKey,omitempty"`
	PatchSet       PatchSet      `json:"patchSet,omitempty"`
	EventCreatedOn int           `json:"eventCreatedOn,omitempty"`
	Reason         string        `json:"reason,omitempty"`
	Abandoner      AccountInfo   `json:"abandoner,omitempty"`
	Restorer       AccountInfo   `json:"restorer,omitempty"`
	Submitter      AccountInfo   `json:"submitter,omitempty"`
	Author         AccountInfo   `json:"author,omitempty"`
	Uploader       AccountInfo   `json:"uploader,omitempty"`
	Approvals      []AccountInfo `json:"approvals,omitempty"`
	Comment        string        `json:"comment,omitempty"`
	Editor         AccountInfo   `json:"editor,omitempty"`
	Added          []string      `json:"added,omitempty"`
	Removed        []string      `json:"removed,omitempty"`
	Hashtags       []string      `json:"hashtags,omitempty"`
	RefUpdate      RefUpdate     `json:"refUpdate,omitempty"`
	Project        ProjectInfo   `json:"project,omitempty"`
	Reviewer       AccountInfo   `json:"reviewer,omitempty"`
	OldTopic       string        `json:"oldTopic,omitempty"`
	Changer        AccountInfo   `json:"changer,omitempty"`
}

// EventsLogService contains functions for querying the API provided
// by the optional events-log plugin.
type EventsLogService struct {
	client *Client
}

// EventsLogOptions contains options for querying events from the events-logs
// plugin.
type EventsLogOptions struct {
	From time.Time
	To   time.Time

	// IgnoreUnmarshalErrors will cause GetEvents to ignore any errors
	// that come up when calling json.Unmarshal. This can be useful in
	// cases where the events-log plugin was not kept up to date with
	// the Gerrit version for some reason. In these cases the events-log
	// plugin will return data structs that don't match the EventInfo
	// struct which in turn causes issues for json.Unmarshal.
	IgnoreUnmarshalErrors bool
}

// getURL returns the url that should be used in the request.  This will vary
// depending on the options provided to GetEvents.
func (events *EventsLogService) getURL(options *EventsLogOptions) (string, error) {
	parsed, err := url.Parse("/plugins/events-log/events/")
	if err != nil {
		return "", err
	}

	query := parsed.Query()

	if !options.From.IsZero() {
		query.Set("t1", options.From.Format("2006-01-02 15:04:05"))
	}

	if !options.To.IsZero() {
		query.Set("t2", options.To.Format("2006-01-02 15:04:05"))
	}

	encoded := query.Encode()
	if len(encoded) > 0 {
		parsed.RawQuery = encoded
	}

	return parsed.String(), nil
}

// GetEvents returns a list of events for the given input options.  Use of this
// function requires an authenticated user and for the events-log plugin to be
// installed. This function returns the unmarshalled EventInfo structs, response,
// failed lines and errors. Marshaling errors will cause this function to return
// before processing is complete unless you set EventsLogOptions.IgnoreUnmarshalErrors
// to true. This can be useful in cases where the events-log plugin got out of sync
// with the Gerrit version which in turn produced events which can't be transformed
// unmarshalled into EventInfo.
//
// Gerrit API docs: https://<yourserver>/plugins/events-log/Documentation/rest-api-events.html
func (events *EventsLogService) GetEvents(options *EventsLogOptions) ([]EventInfo, *Response, [][]byte, error) {
	info := []EventInfo{}
	failures := [][]byte{}
	requestURL, err := events.getURL(options)

	if err != nil {
		return info, nil, failures, err
	}

	request, err := events.client.NewRequest("GET", requestURL, nil)
	if err != nil {
		return info, nil, failures, err
	}

	// Perform the request but do not pass in a structure to unpack
	// the response into.  The format of the response is one EventInfo
	// object per line so we need to manually handle the response here.
	response, err := events.client.Do(request, nil)
	if err != nil {
		return info, response, failures, err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close() // nolint: errcheck
	if err != nil {
		return info, response, failures, err
	}

	for _, line := range bytes.Split(body, []byte("\n")) {
		if len(line) > 0 {
			event := EventInfo{}
			if err := json.Unmarshal(line, &event); err != nil { // nolint: vetshadow
				failures = append(failures, line)

				if !options.IgnoreUnmarshalErrors {
					return info, response, failures, err
				}
				continue
			}
			info = append(info, event)
		}
	}
	return info, response, failures, err
}
