package domain

import (
	"time"
)

//Tweet ... struct
type Tweet struct {
	ID   int
	User string
	Text string
	Date *time.Time
}

//NewTweet ...
func NewTweet(user string, text string) *Tweet {
	date := time.Now()
	var ID int
	tw := Tweet{
		ID, user, text, &date,
	}
	return &tw
}
