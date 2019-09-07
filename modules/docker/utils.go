package docker

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func padSlice(padLeft bool, slice interface{}, getter func(i int) string, setter func(i int, newVal string)) {
	rv := reflect.ValueOf(slice)
	length := rv.Len()
	maxLen := 0
	for i := 0; i < length; i++ {
		val := getter(i)
		maxLen = int(math.Max(float64(len(val)), float64(maxLen)))
	}

	sign := "-"
	if padLeft {
		sign = ""
	}

	for i := 0; i < length; i++ {
		val := getter(i)
		val = fmt.Sprintf("%"+sign+strconv.Itoa(maxLen)+"s", val)
		setter(i, val)
	}
}
