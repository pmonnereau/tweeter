package service

var tweet string

//GetTweet ...
func GetTweet() string {
	return tweet
}

//PublishTweet ...
func PublishTweet(tw string) {
	tweet = tw
}

//CleanTweet ...
func CleanTweet() {
	tweet = ""
}
