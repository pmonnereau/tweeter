package service_test

import (
	"testing"

	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.TweetText

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweetText(user, text)

	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweet()

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweetText(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTweetText(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members 
	all over the world. To date all community oriented activities have been organized by the community
	with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTweetText(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}
func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweetText(user, text)
	secondTweet = domain.NewTweetText(user, secondText)

	// Operation
	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweetText(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetByID(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweetText(user, text)
	secondTweet = domain.NewTweetText(user, secondText)
	thirdTweet = domain.NewTweetText(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweetText(user, text)
	secondTweet = domain.NewTweetText(user, secondText)
	thirdTweet = domain.NewTweetText(anotherUser, text)

	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(tweets))
		return
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func isValidTweet(t *testing.T, tweet domain.Tweet, id int, user, text string) bool {

	if tweet.GetId() != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.GetId())
	}

	if tweet.GetUser() != user && tweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.GetUser(), tweet.GetText())
		return false
	}

	if tweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

func TestTrendingTopics(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "rodri"
	anotherUser := "pri"
	text := "hola soy rodri"
	secondText := "hola soy pro"

	tweet = domain.NewTweetText(user, text)
	secondTweet = domain.NewTweetText(user, secondText)
	thirdTweet = domain.NewTweetText(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	trendingTopics := tweetManager.GetTrendingTopic()
	println(trendingTopics[0])
	println(trendingTopics[1])
	if trendingTopics[0] != "hola" && trendingTopics[0] != "soy" {
		t.Error("Expected hola soy")
		return
	}

	if trendingTopics[1] != "hola" && trendingTopics[1] != "soy" {
		t.Error("Expected hola soy")
		return
	}
}

func TestSendDirectMessage(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	var firstMes, secondMes, thirdMes *domain.DirectMessage

	user := "rodri"
	anotherUser := "pri"
	text := "hola soy rodri"
	text3 := "hola soy pro"
	text2 := "holissss"

	firstMes = tweetManager.SendDirectMessage(user, anotherUser, text)
	secondMes = tweetManager.SendDirectMessage(user, anotherUser, text2)
	thirdMes = tweetManager.SendDirectMessage(anotherUser, user, text3)

	if tweetManager.DirectMessages[0].ID != firstMes.ID {
		t.Error("No es el mensaje directo")
	}
	if tweetManager.DirectMessages[1].ID != secondMes.ID {
		t.Error("No es el mensaje directo")
	}

	if tweetManager.DirectMessages[2].ID != thirdMes.ID {
		t.Error("No es el mensaje directo")
	}
}

func TestGetAllDirectMessagesByUser(t *testing.T) {
	tweetManager := service.NewTweetManager()

	var firstMes, secondMes *domain.DirectMessage

	user := "rodri"
	anotherUser := "pri"
	text := "hola soy rodri"
	text3 := "hola soy pro"
	text2 := "holissss"

	firstMes = tweetManager.SendDirectMessage(user, anotherUser, text)
	secondMes = tweetManager.SendDirectMessage(user, anotherUser, text2)
	_ = tweetManager.SendDirectMessage(anotherUser, user, text3)

	dms := tweetManager.GetAllDirectMessages("pri")
	if dms[0].ID != firstMes.ID {
		t.Error("No es el mensaje directo")
	}
	if dms[1].ID != secondMes.ID {
		t.Error("No es el mensaje directo")
	}

}

func TestGetUnreadedMessagesByUser(t *testing.T) {
	tweetManager := service.NewTweetManager()

	var firstMes, secondMes *domain.DirectMessage

	user := "rodri"
	anotherUser := "pri"
	text := "hola soy rodri"
	text3 := "hola soy pro"
	text2 := "holissss"

	firstMes = tweetManager.SendDirectMessage(user, anotherUser, text)
	secondMes = tweetManager.SendDirectMessage(user, anotherUser, text2)
	thirdMes := tweetManager.SendDirectMessage(anotherUser, user, text3)

	dms := tweetManager.GetUnreadedDirectMessages("pri")
	if dms[0].ID != firstMes.ID {
		t.Error("No es el mensaje directo")
	}
	if dms[1].ID != secondMes.ID {
		t.Error("No es el mensaje directo")
	}

	_ = tweetManager.ReadDirectMessage(firstMes.ID)
	_ = tweetManager.ReadDirectMessage(secondMes.ID)
	_ = tweetManager.ReadDirectMessage(thirdMes.ID)

	dms = tweetManager.GetUnreadedDirectMessages("pri")
	dms2 := tweetManager.GetUnreadedDirectMessages("rodri")
	if len(dms) != 0 {
		t.Error("Hay mensajes sin leer y no deberia")
	}

	if len(dms2) != 0 {
		t.Error("Hay mensajes sin leer y no deberia")
	}

}
func TestReTwitear(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "rodri"
	anotherUser := "pri"
	text := "hola soy rodri"
	secondText := "hola soy pro"

	tweet = domain.NewTweetText(user, text)
	secondTweet = domain.NewTweetText(user, secondText)
	thirdTweet = domain.NewTweetText(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	tweetManager.Retweet("pri", secondTweet.GetId())

	tweetsByUser := tweetManager.GetTweetsByUser("pri")

	println(tweetsByUser[len(tweetsByUser)-1].GetId())
	println(tweetsByUser[len(tweetsByUser)-1].GetText())
	println(secondTweet.GetId())
	println(secondTweet.GetText())
	if tweetsByUser[len(tweetsByUser)-1].GetId() != secondTweet.GetId() {
		t.Error("Hay mensajes sin leer y no deberia")
	}

}
