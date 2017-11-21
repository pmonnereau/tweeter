package service_test

import (
	"testing"

	"github.com/tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	var tweet string = "This is my first Tweet"

	service.PublishTweet(tweet)

	if service.Tweet != tweet {
		t.Error("Expected tweet is", tweet)
	}
}
