package utils

import (
	"strconv"
)

/* -------------------- Map Conversion -------------------- */

// MapToStrs takes a map of interfaces and returns a map of strings
func MapToStrs(aMap map[string]interface{}) map[string]string {
	results := make(map[string]string)

	for key, val := range aMap {
		results[key] = val.(string)
	}

	return results
}

/* -------------------- Slice Conversion -------------------- */

// ToInts takes a slice of interfaces and returns a slice of ints
func ToInts(slice []interface{}) []int {
	results := []int{}

	for _, val := range slice {
		results = append(results, val.(int))
	}

	return results
}

// ToStrs takes a slice of interfaces and returns a slice of strings
func ToStrs(slice []interface{}) []string {
	results := []string{}

	for _, val := range slice {
		switch val.(type) {
		case int:
			results = append(results, strconv.Itoa(val.(int)))
		case string:
			results = append(results, val.(string))
		}
	}

	return results
}
