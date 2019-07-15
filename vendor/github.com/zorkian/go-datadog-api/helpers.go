/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2017 by authors and contributors.
 */

package datadog

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// GetBool is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetBool(v *bool) (bool, bool) {
	if v != nil {
		return *v, true
	}

	return false, false
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// GetIntOk is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetIntOk(v *int) (int, bool) {
	if v != nil {
		return *v, true
	}

	return 0, false
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// GetStringOk is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetStringOk(v *string) (string, bool) {
	if v != nil {
		return *v, true
	}

	return "", false
}

// JsonNumber is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func JsonNumber(v json.Number) *json.Number { return &v }

// GetJsonNumberOk is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetJsonNumberOk(v *json.Number) (json.Number, bool) {
	if v != nil {
		return *v, true
	}

	return "", false
}

// Precision is a helper routine that allocates a new precision value
// to store v and returns a pointer to it.
func Precision(v PrecisionT) *PrecisionT { return &v }

// GetPrecision is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetPrecision(v *PrecisionT) (PrecisionT, bool) {
	if v != nil {
		return *v, true
	}

	return PrecisionT(""), false
}

// GetStringId is a helper routine that allows screenboards and timeboards to be retrieved
// by either the legacy numerical format or the new string format.
// It returns the id as is if it is a string, converts it to a string if it is an integer.
// It return an error if the type is neither string or an integer
func GetStringId(id interface{}) (string, error) {
	switch v := id.(type) {
	case int:
		return strconv.Itoa(v), nil
	case string:
		return v, nil
	default:
		return "", errors.New("unsupported id type")
	}
}
