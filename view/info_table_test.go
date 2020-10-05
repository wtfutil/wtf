package view

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeMap() map[string]string {
	m := make(map[string]string)
	m["foo"] = "val1"
	m["bar"] = "val2"
	m["baz"] = "val3"

	return m
}

func newTestTable(height int) *InfoTable {
	var headers [2]string

	headers[0] = "hdr0"
	headers[1] = "hdr1"

	table := NewInfoTable(
		headers[:],
		makeMap(),
		1,
		1,
		height,
	)
	return table
}

func Test_RenderSimpleInfoTable(t *testing.T) {
	table := newTestTable(4).Render()

	assert.Equal(t, " ----- ------ \n  HDR0   HDR1  \n ----- ------ \n  bar   val2  \n  baz   val3  \n  foo   val1  \n ----- ------ \n", table)
}

func Test_RenderPaddedInfoTable(t *testing.T) {
	table := newTestTable(6).Render()

	assert.Equal(t, " ----- ------ \n  HDR0   HDR1  \n ----- ------ \n  bar   val2  \n  baz   val3  \n  foo   val1  \n              \n              \n ----- ------ \n", table)
}
