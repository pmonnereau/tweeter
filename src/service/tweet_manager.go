package service

import (
	"fmt"
	"strings"

	"github.com/tweeter/src/domain"
)

//TweetManager ...
type TweetManager struct {
	TweetsByUser   map[string][]domain.Tweet
	Tweets         []domain.Tweet
	DirectMessages []domain.DirectMessage
	UserFollows    map[string][]string
	FavsByUser     map[string][]domain.Tweet
}

//NewTweetManager ...
func NewTweetManager() *TweetManager {
	Tweets := make([]domain.Tweet, 0)
	TweetsByUser := make(map[string][]domain.Tweet)
	FavsByUser := make(map[string][]domain.Tweet)
	UserFollows := make(map[string][]string)
	DirectMessages := make([]domain.DirectMessage, 0)
	tweetManager := TweetManager{
		TweetsByUser, Tweets, DirectMessages, UserFollows, FavsByUser,
	}
	return &tweetManager
}

//PublishTweet ...
func (t *TweetManager) PublishTweet(tw domain.Tweet) (int, error) {
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
func (t *TweetManager) GetTweets() []domain.Tweet {
	return t.Tweets
}

//GetTweet ...
func (t *TweetManager) GetTweet() domain.Tweet {
	return t.Tweets[len(t.Tweets)-1]
}

//GetTweetByID ...
func (t *TweetManager) GetTweetByID(id int) domain.Tweet {
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
func (t *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	var lista []domain.Tweet
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
func (t *TweetManager) GetTimeline(user string) []domain.Tweet {
	listaFollows := t.UserFollows[user]
	listaTweets := make([]domain.Tweet, 0)
	for i := 0; i < len(listaFollows); i++ {
		listaTweets = append(listaTweets, t.TweetsByUser[listaFollows[i]]...)

	}
	misTweets := t.TweetsByUser[user]
	listaTweets = append(listaTweets, misTweets...)
	return listaTweets
}

//GetTrendingTopic ...
func (t *TweetManager) GetTrendingTopic() []string {
	listaTrending := make(map[string]int)
	listaRes := make([]string, 0)
	//Devuelve un mapa con (clave,sig) = (palabra,cantRep)
	for i := 0; i < len(t.Tweets); i++ {
		listaPalabras := strings.Fields(t.Tweets[i].Text)
		for j := 0; j < len(listaPalabras); j++ {
			elem := listaTrending[listaPalabras[j]]
			listaTrending[listaPalabras[j]] = elem + 1
		}
	}
	maxRepeticion := 0
	palabraMaxima := ""
	palabraMaximaSig := ""
	for i := range listaTrending {
		value := listaTrending[i]
		if listaTrending[i] > maxRepeticion {
			maxRepeticion = value
			palabraMaxima = i
		}
	}

	maxRepeticion = 0

	for i := range listaTrending {
		value := listaTrending[i]
		if listaTrending[i] > maxRepeticion && i != palabraMaxima {
			maxRepeticion = value
			palabraMaximaSig = i
		}
	}

	listaRes = append(listaRes, palabraMaxima)
	listaRes = append(listaRes, palabraMaximaSig)

	return listaRes

}

//SendDirectMessage ...
func (t *TweetManager) SendDirectMessage(user string, user2 string, text string) *domain.DirectMessage {
	msje1 := domain.NewDirectMessages(user, user2, text)
	msje1.ID = len(t.DirectMessages)
	t.DirectMessages = append(t.DirectMessages, msje1)
	return msje1
}

//GetAllDirectMessages ...
func (t *TweetManager) GetAllDirectMessages(user string) []*domain.DirectMessage {
	misDm := make([]*domain.DirectMessage, 0)
	for i := 0; i < len(t.DirectMessages); i++ {
		if t.DirectMessages[i].ToUser == user {
			misDm = append(misDm, t.DirectMessages[i])
		}
	}
	return misDm
}

//GetUnreadedDirectMessages ...
func (t *TweetManager) GetUnreadedDirectMessages(user string) []*domain.DirectMessage {
	misDm := make([]*domain.DirectMessage, 0)
	for i := 0; i < len(t.DirectMessages); i++ {
		if t.DirectMessages[i].ToUser == user && !t.DirectMessages[i].Read {
			misDm = append(misDm, t.DirectMessages[i])
		}
	}
	return misDm
}

//ReadDirectMessage ...
func (t *TweetManager) ReadDirectMessage(msj *domain.DirectMessage) *domain.DirectMessage {
	var msj2 *domain.DirectMessage
	for i := 0; i < len(t.DirectMessages); i++ {
		if msj.ID == t.DirectMessages[i].ID {
			t.DirectMessages[i].Read = true
			msj2 = t.DirectMessages[i]
		}
	}
	return msj2
}

//Retweet ...
func (t *TweetManager) Retweet(user string, twit *domain.Tweet) domain.Tweet {
	twits := t.TweetsByUser[user]
	twits = append(twits, twit)
	t.TweetsByUser[user] = twits

	return twit

}
