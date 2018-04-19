package clocks

import(
	"time"
)

type Clock struct {
	Label     string
	LocalTime time.Time
	Timezone  string
}
