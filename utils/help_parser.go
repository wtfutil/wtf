package utils

import (
	"reflect"
	"regexp"
	"strconv"
	"unicode"
	"unicode/utf8"

	"github.com/wtfutil/wtf/cfg"
)

func lowercaseTitle(title string) string {
	if title == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(title)
	return string(unicode.ToLower(r)) + title[n:]
}

var (
	openColorRegex = regexp.MustCompile(`\[.*?\]`)
)

func StripColorTags(input string) string {
	return openColorRegex.ReplaceAllString(input, "")
}

func helpFromValue(field reflect.StructField) string {
	result := ""
	var help string = field.Tag.Get("help")
	optional, err := strconv.ParseBool(field.Tag.Get("optional"))
	if err != nil {
		optional = false
	}
	var values string = field.Tag.Get("values")
	if optional {
		help = "Optional " + help
	}
	if help != "" {
		result += "\n\n" + lowercaseTitle(field.Name)
		result += "\n" + help
		if values != "" {
			result += "\nValues: " + values
		}
	}

	return result
}

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
