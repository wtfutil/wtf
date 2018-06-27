package gerrit

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// Timestamp represents an instant in time with nanosecond precision, in UTC time zone.
// It encodes to and from JSON in Gerrit's timestamp format.
// All exported methods of time.Time can be called on Timestamp.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api.html#timestamp
type Timestamp struct {
	// Time is an instant in time. Its time zone must be UTC.
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in Gerrit's timestamp format.
// An error is returned if t.Time time zone is not UTC.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	if t.Location() != time.UTC {
		return nil, errors.New("Timestamp.MarshalJSON: time zone must be UTC")
	}
	if y := t.Year(); y < 0 || 9999 < y {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#issuecomment-66073163 for more discussion.
		return nil, errors.New("Timestamp.MarshalJSON: year outside of range [0,9999]")
	}
	b := make([]byte, 0, len(timeLayout)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, timeLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in Gerrit's timestamp format.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the main JSON package.
	if string(b) == "null" {
		return nil
	}
	var err error
	t.Time, err = time.Parse(`"`+timeLayout+`"`, string(b))
	return err
}

// Gerrit's timestamp layout is like time.RFC3339Nano, but with a space instead
// of the "T", without a timezone (it's always in UTC), and always includes nanoseconds.
// See https://gerrit-review.googlesource.com/Documentation/rest-api.html#timestamp.
const timeLayout = "2006-01-02 15:04:05.000000000"

// Number is a string representing a number. This type is only used in cases
// where the API being queried may return an inconsistent result.
type Number string

// String returns the string representing the current number.
func (n *Number) String() string {
	return string(*n)
}

// Int returns the current number as an integer
func (n *Number) Int() (int, error) {
	return strconv.Atoi(n.String())
}

// UnmarshalJSON will marshal the provided data into the current *Number struct.
func (n *Number) UnmarshalJSON(data []byte) error {
	// `data` is a number represented as a string (ex. "5").
	var stringNumber string
	if err := json.Unmarshal(data, &stringNumber); err == nil {
		*n = Number(stringNumber)
		return nil
	}

	// `data` is a number represented as an integer (ex. 5). Here
	// we're using json.Unmarshal to convert bytes -> number which
	// we then convert to our own Number type.
	var number int
	if err := json.Unmarshal(data, &number); err == nil {
		*n = Number(strconv.Itoa(number))
		return nil
	}
	return errors.New("cannot convert data to number")
}
