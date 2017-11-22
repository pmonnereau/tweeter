package service

import "github.com/tweeter/src/domain"
import "fmt"

var tweet *domain.Tweet

//GetTweet ...
func GetTweet() *domain.Tweet {
	return tweet
}

//PublishTweet ...
func PublishTweet(tw *domain.Tweet) error {
	var err error
	if tw.User == "" && tw.Text == "" {
		err = fmt.Errorf("text and user are required")
	} else if tw.User == "" {
		err = fmt.Errorf("user is required")
	} else if tw.Text == "" {
		err = fmt.Errorf("text is required")
	} else if len(tw.Text) > 140 {
		err = fmt.Errorf("tweet must be less than 140 chars")
	} else {
		tweet = tw
	}
	return err
}

//CleanTweet ...
func CleanTweet() {
	tweet.Text = ""
	tweet.User = ""
	tweet.Date = nil
}
