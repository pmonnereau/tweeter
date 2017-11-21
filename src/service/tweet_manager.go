package service

var tweet string

func GetTweet() string {
	return tweet
}

func PublishTweet(tw string) {
	tweet = tw
}
