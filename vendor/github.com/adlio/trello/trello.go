// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func IDToTime(id string) (t time.Time, err error) {
	if id == "" {
		return time.Time{}, nil
	}
	// The first 8 characters in the object ID are a Unix timestamp
	ts, err := strconv.ParseUint(id[:8], 16, 64)
	if err != nil {
		err = errors.Wrapf(err, "ID '%s' failed to convert to timestamp.", id)
	} else {
		t = time.Unix(int64(ts), 0)
	}
	return
}
