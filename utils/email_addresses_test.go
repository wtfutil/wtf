package utils

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func Test_NameFromEmail(t *testing.T) {
	Equal(t, "", NameFromEmail(""))
	Equal(t, "Chris Cummer", NameFromEmail("chris.cummer@me.com"))
}

func Test_NamesFromEmails(t *testing.T) {
	var result []string

	result = NamesFromEmails([]string{})
	Equal(t, []string{}, result)

	result = NamesFromEmails([]string{"chris.cummer@me.com", "chriscummer@me.com"})
	Equal(t, []string{"Chris Cummer", "Chriscummer"}, result)
}
