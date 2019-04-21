package wtf_tests

import (
	"testing"

	. "github.com/stretchr/testify/assert"
	. "github.com/wtfutil/wtf/wtf"
)

func Test_ASCIItoTviewColors(t *testing.T) {
	Equal(t, "", ASCIItoTviewColors(""))
	Equal(t, "cat", ASCIItoTviewColors("cat"))
	Equal(t, "[38;5;226mcat/[-]", ASCIItoTviewColors("[38;5;226mcat/[0m"))
}
