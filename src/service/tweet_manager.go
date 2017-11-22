package service

import "github.com/tweeter/src/domain"
import "fmt"

var tweets []*domain.Tweet

//PublishTweet ...
func PublishTweet(tw *domain.Tweet) (int, error) {
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
		tw.ID = len(tweets)
		tweets = append(tweets, tw)
	}
	return tw.ID, err
}

//CleanLastTweet ...
func CleanLastTweet() {
	tweets = tweets[0 : len(tweets)-1]
}

//InitializeService ...
func InitializeService() {
	tweets = make([]*domain.Tweet, 0)

}

//GetTweets ....
func GetTweets() []*domain.Tweet {
	return tweets
}

//GetTweetByID ...
func GetTweetByID(id int) *domain.Tweet {
	return tweets[id]
}
