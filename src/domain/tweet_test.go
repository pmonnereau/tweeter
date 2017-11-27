package domain_test

import (
	"testing"

	"github.com/tweeter/src/domain"
)

func TestCanGetAPrintableTweet(t *testing.T) {
	//Inicialization
	var tweet domain.Tweet
	tweet = domain.NewTweetText("grupoesfera", "This is my tweet!")
	//Operation
	text := tweet.PrintableTweet()
	//Validation
	expectedText := "@grupoesfera: This is my tweet!"
	if text != expectedText {
		t.Errorf("Expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	var tweet domain.Tweet
	tweet = domain.NewTweetImage("grupoesfera", "This is my image", "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	// Initialization
	var tweet, quotedTweet domain.Tweet
	quotedTweet = domain.NewTweetText("grupoesfera", "This is my tweet")
	tweet = domain.NewTweetQuote("nick", "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTweetText("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
