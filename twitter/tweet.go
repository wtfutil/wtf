package twitter

import (
	"fmt"

	"github.com/senorprogrammer/wtf/wtf"
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
	return fmt.Sprint(tweet.User.ScreenName)
}
func (tweet *Tweet) PrettyStart() string {
	return wtf.PrettyDate(tweet.CreatedAt)
}
