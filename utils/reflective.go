package utils

import (
	"fmt"
	"reflect"
)

// StringValueForProperty returns a string value for the given property
// If the property doesn't exist, it returns an error
func StringValueForProperty(ref interface{}, propName string) (string, error) {
	v := reflect.ValueOf(ref)
	refVal := reflect.Indirect(v).FieldByName(propName)

	if !refVal.IsValid() {
		return "", fmt.Errorf("invalid property name: %s", propName)
	}

	strVal := fmt.Sprintf("%v", refVal)

	return strVal, nil
}
