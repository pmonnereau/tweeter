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
	DirectMessages []*domain.DirectMessage
	UserFollows    map[string][]string
	FavsByUser     map[string][]domain.Tweet
	ChannelTW      ChannelTweetWriter
}

//NewTweetManager ...
func NewTweetManager(ChannelTW ChannelTweetWriter) *TweetManager {
	Tweets := make([]domain.Tweet, 0)
	TweetsByUser := make(map[string][]domain.Tweet)
	FavsByUser := make(map[string][]domain.Tweet)
	UserFollows := make(map[string][]string)
	DirectMessages := make([]*domain.DirectMessage, 0)
	tweetManager := TweetManager{
		TweetsByUser, Tweets, DirectMessages, UserFollows, FavsByUser, ChannelTW,
	}
	return &tweetManager
}

//PublishTweet ...
func (t *TweetManager) PublishTweet(tw domain.Tweet, quit chan bool) (int, error) {
	var err error
	if tw.GetUser() == "" && tw.GetText() == "" {
		err = fmt.Errorf("text and user are required")
	} else if tw.GetUser() == "" {
		err = fmt.Errorf("user is required")
	} else if tw.GetText() == "" {
		err = fmt.Errorf("text is required")
	} else if len(tw.GetText()) > 140 {
		err = fmt.Errorf("text exceeds 140 characters")
	} else {
		tweetsToWrite := make(chan domain.Tweet)
		go t.ChannelTW.WriteTweet(tweetsToWrite, quit)
		tweetsToWrite <- tw
		close(tweetsToWrite)
		tw.SetId(len(t.Tweets))
		t.Tweets = append(t.Tweets, tw)
		elem, _ := t.TweetsByUser[tw.GetUser()]
		elem = append(elem, tw)
		t.TweetsByUser[tw.GetUser()] = elem
	}
	return tw.GetId(), err
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

//Follow ...
func (t *TweetManager) MyFollows(user string) []string {
	elem := t.UserFollows[user]
	return elem
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
		listaPalabras := strings.Fields(t.Tweets[i].GetText())
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
func (t *TweetManager) ReadDirectMessage(msj int) *domain.DirectMessage {
	var msj2 *domain.DirectMessage
	for i := 0; i < len(t.DirectMessages); i++ {
		if msj == t.DirectMessages[i].ID {
			t.DirectMessages[i].Read = true
			msj2 = t.DirectMessages[i]
		}
	}
	return msj2
}

//Retweet ...
func (t *TweetManager) Retweet(user string, idTwit int) domain.Tweet {
	twits := t.TweetsByUser[user]
	twits = append(twits, t.GetTweetByID(idTwit))
	t.TweetsByUser[user] = twits

	return t.GetTweetByID(idTwit)

}

//AddToFavs ...
func (t *TweetManager) AddToFavs(user string, idTwit int) domain.Tweet {
	favs := t.FavsByUser[user]
	favs = append(favs, t.GetTweetByID(idTwit))
	t.FavsByUser[user] = favs

	return t.GetTweetByID(idTwit)

}

//GetUserFavs ...
func (t *TweetManager) GetUserFavs(user string) []domain.Tweet {
	listaFavs := t.FavsByUser[user]
	return listaFavs
}
