package bamboohr

import (
	"time"
)

func Fetch() []Item {
	client := NewClient()
	result := client.Away("timeOff", today(), today())

	return result
}

func today() string {
	localNow := time.Now().Local()
	return localNow.Format("2006-01-02")
}
