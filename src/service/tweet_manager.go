package service

import "github.com/tweeter/src/domain"
import "fmt"

//TweetManager ...
type TweetManager struct {
	TweetsByUser map[string][]*domain.Tweet
	Tweets       []*domain.Tweet
	UserFollows  map[string][]string
}

//NewTweetManager ...
func NewTweetManager() *TweetManager {
	Tweets := make([]*domain.Tweet, 0)
	TweetsByUser := make(map[string][]*domain.Tweet)
	UserFollows := make(map[string][]string)
	tweetManager := TweetManager{
		TweetsByUser, Tweets, UserFollows,
	}
	return &tweetManager
}

//PublishTweet ...
func (t *TweetManager) PublishTweet(tw *domain.Tweet) (int, error) {
	var err error
	if tw.User == "" && tw.Text == "" {
		err = fmt.Errorf("text and user are required")
	} else if tw.User == "" {
		err = fmt.Errorf("user is required")
	} else if tw.Text == "" {
		err = fmt.Errorf("text is required")
	} else if len(tw.Text) > 140 {
		err = fmt.Errorf("text exceeds 140 characters")
	} else {
		tw.ID = len(t.Tweets)
		t.Tweets = append(t.Tweets, tw)
		elem, _ := t.TweetsByUser[tw.User]
		elem = append(elem, tw)
		t.TweetsByUser[tw.User] = elem
	}
	return tw.ID, err
}

//CleanLastTweet ...
func (t *TweetManager) CleanLastTweet() {
	t.Tweets = t.Tweets[0 : len(t.Tweets)-1]
}

//GetTweets ....
func (t *TweetManager) GetTweets() []*domain.Tweet {
	return t.Tweets
}

//GetTweet ...
func (t *TweetManager) GetTweet() *domain.Tweet {
	return t.Tweets[len(t.Tweets)-1]
}

//GetTweetByID ...
func (t *TweetManager) GetTweetByID(id int) *domain.Tweet {
	return t.Tweets[id]
}

//CountTweetsByUser ...
func (t *TweetManager) CountTweetsByUser(user string) int {
	elem, _ := t.TweetsByUser[user]
	var suma int
	for i := 0; i < len(elem); i++ {
		suma++
	}
	return suma
}

//GetTweetsByUser ...
func (t *TweetManager) GetTweetsByUser(user string) []*domain.Tweet {
	var lista []*domain.Tweet
	elem, ok := t.TweetsByUser[user]
	if ok {
		lista = elem
	}
	return lista
}

//Follow ...
func (t *TweetManager) Follow(user string, userFollow string) {
	elem := t.UserFollows[user]
	t.UserFollows[user] = append(elem, userFollow)
}

//GetTimeline ...
func (t *TweetManager) GetTimeline(user string) []*domain.Tweet {
	listaFollows := t.UserFollows[user]
	listaTweets := make([]*domain.Tweet, 0)
	for i := 0; i < len(listaFollows); i++ {
		listaTweets = append(listaTweets, t.TweetsByUser[listaFollows[i]]...)

	}
	misTweets := t.TweetsByUser[user]
	listaTweets = append(listaTweets, misTweets...)
	return listaTweets
}
