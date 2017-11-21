package service_test

import (
	"testing"

	"github.com/tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	tweet := "This is my first Tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}
