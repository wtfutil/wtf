package utils

import (
	"fmt"
	"reflect"
)

// Reflective is a convenience wrapper for objects that makes it possible to
// extract property values from the object by property name
type Reflective struct{}

// StringValueForProperty returns a string value for the given property
// If the property doesn't exist, it returns an error
func (ref *Reflective) StringValueForProperty(propName string) (string, error) {
	v := reflect.ValueOf(ref)
	refVal := reflect.Indirect(v).FieldByName(propName)

	if !refVal.IsValid() {
		return "", fmt.Errorf("invalid property name: %s", propName)
	}

	strVal := fmt.Sprintf("%v", refVal)

	return strVal, nil
}
