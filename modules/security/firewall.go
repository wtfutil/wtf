package security

import (
	"strings"
)

func statusLabel(str string) string {
	label := "off"

	if strings.Contains(str, "enabled") {
		label = "on"
	}

	return label
}
