package wtf_tests

import (
	"testing"
	"time"

	. "github.com/senorprogrammer/wtf/wtf"
	. "github.com/stretchr/testify/assert"
)

func TestIsToday(t *testing.T) {
	Equal(t, true, IsToday(time.Now().Local()))
	Equal(t, false, IsToday(time.Now().AddDate(0, 0, -1)))
	Equal(t, false, IsToday(time.Now().AddDate(0, 0, +1)))
}

/* -------------------- PrettyDate() -------------------- */

func TestPrettyDate(t *testing.T) {
	Equal(t, "Oct 21, 1999", PrettyDate("1999-10-21"))
}
