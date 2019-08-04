package wtf

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func Test_ASCIItoTviewColors(t *testing.T) {
	Equal(t, "", ASCIItoTviewColors(""))
	Equal(t, "cat", ASCIItoTviewColors("cat"))
	Equal(t, "[38;5;226mcat/[-]", ASCIItoTviewColors("[38;5;226mcat/[0m"))
}
