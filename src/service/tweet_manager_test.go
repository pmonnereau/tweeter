package service_test

import (
	"testing"

	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	user := "rodri"
	text := "holahola"
	var tweet *domain.Tweet
	tweet = domain.NewTweet(user, text)
	service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()

	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \n but is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date cant be nil")
	}
}

func TestCleanTweet(t *testing.T) {
	tweet := domain.NewTweet("priscila", "hola es mi tweet")

	service.PublishTweet(tweet)

	service.CleanTweet()
	tweet_vacio := ""
	if service.GetTweet().Text != tweet_vacio {
		t.Error("Expected tweet is", tweet_vacio)
	}
}
func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	var user string
	text := "this is my first tweet"

	tweet = domain.NewTweet(user, text)
	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error: user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	user := "grupoesfera"
	var text string
	tweet = domain.NewTweet(user, text)
	var err error
	err = service.PublishTweet(tweet)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err.Error() != "text is required" {
		t.Error("Expected error: text is required")
	}

}

func TestTweetWithoutTextAndUserIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	var user string
	var text string
	tweet = domain.NewTweet(user, text)
	var err error
	err = service.PublishTweet(tweet)
	if err == nil {
		t.Error("expected error")
		return
	}
	if err.Error() != "text and user are required" {
		t.Error("Expected error: text and user are required")
	}

}

func TestTweetTextMustHaveLessThan140Chars(t *testing.T) {
	var tweet *domain.Tweet
	user := "pepe"
	text := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("expected error")
		return
	}
	if err.Error() != "tweet must be less than 140 chars" {
		t.Error("Expected error is tweet must be less than 140 chars")
	}

}
