package wtf_tests

import (
	"testing"

	. "github.com/senorprogrammer/wtf/wtf"
	. "github.com/stretchr/testify/assert"
)

func TestASCIItoTviewColors(t *testing.T) {
	Equal(t, "", ASCIItoTviewColors(""))
	Equal(t, "cat", ASCIItoTviewColors("cat"))
	Equal(t, "[38;5;226mcat/[-]", ASCIItoTviewColors("[38;5;226mcat/[0m"))
}
