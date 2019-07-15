package transmissionrpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

/*
	Torrent Mutators
	https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L105
*/

// TorrentSet apply a list of mutator(s) to a list of torrent ids.
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L107
func (c *Client) TorrentSet(payload *TorrentSetPayload) (err error) {
	// Validate
	if payload == nil {
		return errors.New("payload can't be nil")
	}
	if len(payload.IDs) == 0 {
		return errors.New("there must be at least one ID")
	}
	// Send payload
	if err = c.rpcCall("torrent-set", payload, nil); err != nil {
		err = fmt.Errorf("'torrent-set' rpc method failed: %v", err)
	}
	return
}

// TorrentSetPayload contains all the mutators appliable on one torrent.
// https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L111
type TorrentSetPayload struct {
	BandwidthPriority   *int64         `json:"bandwidthPriority"`   // this torrent's bandwidth tr_priority_t
	DownloadLimit       *int64         `json:"downloadLimit"`       // maximum download speed (KBps)
	DownloadLimited     *bool          `json:"downloadLimited"`     // true if "downloadLimit" is honored
	FilesWanted         []int64        `json:"files-wanted"`        // indices of file(s) to download
	FilesUnwanted       []int64        `json:"files-unwanted"`      // indices of file(s) to not download
	HonorsSessionLimits *bool          `json:"honorsSessionLimits"` // true if session upload limits are honored
	IDs                 []int64        `json:"ids"`                 // torrent list
	Location            *string        `json:"location"`            // new location of the torrent's content
	Peerlimit           *int64         `json:"peer-limit"`          // maximum number of peers
	PriorityHigh        []int64        `json:"priority-high"`       // indices of high-priority file(s)
	PriorityLow         []int64        `json:"priority-low"`        // indices of low-priority file(s)
	PriorityNormal      []int64        `json:"priority-normal"`     // indices of normal-priority file(s)
	QueuePosition       *int64         `json:"queuePosition"`       // position of this torrent in its queue [0...n)
	SeedIdleLimit       *time.Duration `json:"seedIdleLimit"`       // torrent-level number of minutes of seeding inactivity
	SeedIdleMode        *int64         `json:"seedIdleMode"`        // which seeding inactivity to use
	SeedRatioLimit      *float64       `json:"seedRatioLimit"`      // torrent-level seeding ratio
	SeedRatioMode       *SeedRatioMode `json:"seedRatioMode"`       // which ratio mode to use
	TrackerAdd          []string       `json:"trackerAdd"`          // strings of announce URLs to add
	TrackerRemove       []int64        `json:"trackerRemove"`       // ids of trackers to remove
	TrackerReplace      []string       `json:"trackerReplace"`      // pairs of <trackerId/new announce URLs> (TODO: validate string value usable as is)
	UploadLimit         *int64         `json:"uploadLimit"`         // maximum upload speed (KBps)
	UploadLimited       *bool          `json:"uploadLimited"`       // true if "uploadLimit" is honored
}

// MarshalJSON allows to marshall into JSON only the non nil fields.
// It differs from 'omitempty' which also skip default values
// (as 0 or false which can be valid here).
func (tsp *TorrentSetPayload) MarshalJSON() (data []byte, err error) {
	// Build an intermediary payload with base types
	type baseTorrentSetPayload TorrentSetPayload
	tmp := struct {
		SeedIdleLimit *int64 `json:"seedIdleLimit"`
		*baseTorrentSetPayload
	}{
		baseTorrentSetPayload: (*baseTorrentSetPayload)(tsp),
	}
	if tsp.SeedIdleLimit != nil {
		sil := int64(*tsp.SeedIdleLimit / time.Minute)
		tmp.SeedIdleLimit = &sil
	}
	// Build a payload with only the non nil fields
	tspv := reflect.ValueOf(tmp)
	tspt := tspv.Type()
	cleanPayload := make(map[string]interface{}, tspt.NumField())
	var currentValue, nestedStruct, currentNestedValue reflect.Value
	var currentStructField, currentNestedStructField reflect.StructField
	var j int
	for i := 0; i < tspv.NumField(); i++ {
		currentValue = tspv.Field(i)
		currentStructField = tspt.Field(i)
		if !currentValue.IsNil() {
			if currentStructField.Name == "baseTorrentSetPayload" {
				// inherited/nested struct
				nestedStruct = reflect.Indirect(currentValue)
				for j = 0; j < nestedStruct.NumField(); j++ {
					currentNestedValue = nestedStruct.Field(j)
					currentNestedStructField = nestedStruct.Type().Field(j)
					if !currentNestedValue.IsNil() {
						cleanPayload[currentNestedStructField.Tag.Get("json")] = currentNestedValue.Interface()
					}
				}
			} else {
				// Overloaded field
				cleanPayload[currentStructField.Tag.Get("json")] = currentValue.Interface()
			}
		}
	}
	// Marshall the clean payload
	return json.Marshal(cleanPayload)
}
