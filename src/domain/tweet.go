package domain

import (
	"time"
)

//Tweet ... struct
type Tweet interface {
	GetUser() string
	GetText() string
	GetId() int
	SetId() int
	PrintableTweet() string
}

//TweetText ...
type TweetText struct {
	ID   int
	User string
	Text string
	Date *time.Time
}

//TweetQuote ...
type TweetQuote struct {
	TweetText
	Quote string
}

//TweetImage ...
type TweetImage struct {
	TweetText
	URL string
}

//DirectMessage ... struct
type DirectMessage struct {
	ID       int
	FromUser string
	ToUser   string
	Read     bool
	Text     string
}

//NewTweetText ...
func NewTweetText(user string, text string) *TweetText {
	date := time.Now()
	var ID int
	tw := TweetText{
		ID, user, text, &date,
	}
	return &tw
}

//NewTweetQuote ...
func NewTweetQuote(user string, text string, quote string) *TweetQuote {
	date := time.Now()
	var ID int
	tw := TweetQuote{
		TweetText: TweetText{
			ID:   ID,
			User: user,
			Text: text,
			Date: &date,
		},
		Quote: quote,
	}
	return &tw
}

//NewTweetImage ...
func NewTweetImage(user string, text string, url string) *TweetImage {
	date := time.Now()
	var ID int
	tw := TweetImage{
		TweetText: TweetText{
			ID:   ID,
			User: user,
			Text: text,
			Date: &date,
		},
		URL: url,
	}
	return &tw
}

//PrintableTweet ...
func (t *TweetImage) PrintableTweet() string {
	var tweet string
	tweet = "@" + t.User + ": " + t.Text + t.URL
	return tweet
}

//PrintableTweet ...
func (t *TweetQuote) PrintableTweet() string {
	var tweet string
	tweet = "@" + t.User + ": " + t.Text + t.Quote
	return tweet
}

//PrintableTweet ...
func (t *TweetText) PrintableTweet() string {
	var tweet string
	tweet = "@" + t.User + ": " + t.Text
	return tweet
}

func (t *Tweet) String() string {
	return t.PrintableTweet()
}

//NewDirectMessages ...
func NewDirectMessages(fromUser string, toUser string, text string) *DirectMessage {
	var ID int
	var read bool
	read = false
	dm := DirectMessage{
		ID, fromUser, toUser, read, text,
	}
	return &dm
}
