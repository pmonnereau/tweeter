package domain

import (
	"time"
)

//Tweet ... struct
type Tweet struct {
	User string
	Text string
	Date *time.Time
}

//NewTweet ...
func NewTweet(user string, text string) *Tweet {
	date := time.Now()
	tw := Tweet{
		user, text, &date,
	}
	return &tw
}
