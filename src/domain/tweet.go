package domain

import (
	"time"
)

//Tweet ... struct
type Tweet interface {
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetId() int
	SetId(id int)
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
	Quote Tweet
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
func NewTweetQuote(user string, text string, quote Tweet) *TweetQuote {
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

func (t *TweetImage) String() string {
	return t.PrintableTweet()
}

func (t *TweetText) String() string {
	return t.PrintableTweet()
}

func (t *TweetQuote) String() string {
	return t.PrintableTweet()
}

//PrintableTweet ...
func (t *TweetImage) PrintableTweet() string {
	var tweet string
	tweet = "@" + t.User + ": " + t.Text + " " + t.URL
	return tweet
}

//PrintableTweet ...
func (t *TweetQuote) PrintableTweet() string {
	var tweet string
	tweet = "@" + t.User + ": " + t.Text + " " + `"` + t.Quote.PrintableTweet() + `"`
	return tweet
}

//PrintableTweet ...
func (t *TweetText) PrintableTweet() string {
	var tweet string
	tweet = "@" + t.User + ": " + t.Text
	return tweet
}

//PrintableTweet ...
func (t *DirectMessage) PrintableMessage() string {
	var message string
	message = t.FromUser + ": @" + t.ToUser + t.Text
	return message
}

func (t *DirectMessage) String() string {
	return t.PrintableMessage()
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

func (t *TweetText) GetId() int {
	return t.ID
}

func (t *TweetImage) GetId() int {
	return t.ID
}

func (t *TweetQuote) GetId() int {
	return t.ID
}

func (t *TweetText) GetUser() string {
	return t.User
}

func (t *TweetImage) GetUser() string {
	return t.User
}

func (t *TweetQuote) GetUser() string {
	return t.User
}

func (t *TweetText) GetText() string {
	return t.Text
}

func (t *TweetImage) GetText() string {
	return t.Text
}

func (t *TweetQuote) GetText() string {
	return t.Text
}

func (t *TweetText) SetId(i int) {
	t.ID = i
}

func (t *TweetImage) SetId(i int) {
	t.ID = i
}

func (t *TweetQuote) SetId(i int) {
	t.ID = i
}

func (t *TweetText) GetDate() *time.Time {
	return t.Date
}

func (t *TweetImage) GetDate() *time.Time {
	return t.Date
}

func (t *TweetQuote) GetDate() *time.Time {
	return t.Date
}
