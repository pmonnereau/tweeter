package service

import "github.com/tweeter/src/domain"

var tweet *domain.Tweet

//GetTweet ...
func GetTweet() *domain.Tweet {
	return tweet
}

//PublishTweet ...
func PublishTweet(tw *domain.Tweet) {
	tweet = tw
}

//CleanTweet ...
func CleanTweet() {
	tweet.Text = ""
}
