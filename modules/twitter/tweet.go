package twitter

import (
	"fmt"
	"time"
)

type Tweet struct {
	User      User   `json:"user"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (tweet *Tweet) String() string {
	return fmt.Sprintf("Tweet: %s at %s by %s", tweet.Text, tweet.CreatedAt, tweet.User.ScreenName)
}

/* -------------------- Exported Functions -------------------- */

func (tweet *Tweet) Username() string {
	return tweet.User.ScreenName
}

func (tweet *Tweet) Created() time.Time {
	newTime, _ := time.Parse(time.RubyDate, tweet.CreatedAt)
	return newTime
}

func (tweet *Tweet) PrettyCreatedAt() string {
	newTime := tweet.Created()
	return fmt.Sprint(newTime.Format("Jan 2, 2006"))
}
