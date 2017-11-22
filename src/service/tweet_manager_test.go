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

	publishedTweet := service.GetTweets()[len(service.GetTweets())-1]

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
	longitudAntesDeEliminar := len(service.GetTweets())

	service.CleanLastTweet()
	if len(service.GetTweets()) == longitudAntesDeEliminar {
		t.Error("Tweet not cleaned")
	}

}
func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	var user string
	text := "this is my first tweet"

	tweet = domain.NewTweet(user, text)
	var err error

	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)
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
	_, err = service.PublishTweet(tweet)
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
	_, err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("expected error")
		return
	}
	if err.Error() != "tweet must be less than 140 chars" {
		t.Error("Expected error is tweet must be less than 140 chars")
	}

}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstPublishedTweet.ID, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondPublishedTweet.ID, user, secondText) {
		return
	}

}

func isValidTweet(t *testing.T, tweet *domain.Tweet, ID int, user, text string) bool {

	if tweet.ID != ID {
		t.Errorf("Expected ID is %v but was %v", ID, tweet.ID)
	}
	if tweet.User != user && tweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.User, tweet.Text)
		return false
	}

	if tweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetByID(id)

	isValidTweet(t, publishedTweet, id, user, text)

}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	// Operation
	count := service.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}
