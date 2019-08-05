package utils

import (
	"reflect"
	"regexp"
	"strconv"
	"unicode"
	"unicode/utf8"

	"github.com/wtfutil/wtf/cfg"
)

/* -------------------- Exported Functions -------------------- */

func HelpFromInterface(item interface{}) string {
	result := ""
	t := reflect.TypeOf(item)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		kind := field.Type.Kind()
		if field.Type.Kind() == reflect.Ptr {
			kind = field.Type.Elem().Kind()
		}

		if field.Name == "common" {
			result += HelpFromInterface(cfg.Common{})
		}

		switch kind {
		case reflect.Interface:
			result += HelpFromInterface(field.Type.Elem())
		default:
			result += helpFromValue(field)
		}
	}

	return result
}

// StripColorTags removes tcell color tags from a given string
func StripColorTags(input string) string {
	openColorRegex := regexp.MustCompile(`\[.*?\]`)
	return openColorRegex.ReplaceAllString(input, "")
}

/* -------------------- Unexported Functions -------------------- */

func helpFromValue(field reflect.StructField) string {
	result := ""

	optional, err := strconv.ParseBool(field.Tag.Get("optional"))
	if err != nil {
		optional = false
	}

	help := field.Tag.Get("help")
	if optional {
		help = "Optional " + help
	}

	values := field.Tag.Get("values")
	if help != "" {
		result += "\n\n " + lowercaseTitle(field.Name)
		result += "\n " + help

		if values != "" {
			result += "\n Values: " + values
		}
	}

	return result
}

func lowercaseTitle(title string) string {
	if title == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(title)
	return string(unicode.ToLower(r)) + title[n:]
}
