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
	if tw.User == "" {
		err = fmt.Errorf("user is required")
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
