package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func main() {
	shell := ishell.New()
	t := service.NewTweetManager()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type help to know commands \n")
	shell.AddCmd(&ishell.Cmd{
		Name: "publishTextTweet",
		Help: "Publishes a text tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Write your tweet: ")
			txt := c.ReadLine()
			Tweet := domain.NewTweetText(user, txt)
			var err error
			var idNew int
			idNew, err = t.PublishTweet(Tweet)
			if err != nil && err.Error() == "user is required" {
				c.Print("User is required, try again")
			} else if err != nil && err.Error() == "text is required" {
				c.Print("Text is required, try again")
			} else if err != nil && err.Error() == "text and user are required" {
				c.Print("text and user are required, try again")
			} else if err != nil && err.Error() == "tweet must be less than 140 chars" {
				c.Print("tweet must be less than 140 chars")
			} else {
				c.Print("Tweet sent with id ", idNew)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishQuoteTweet",
		Help: "Publishes a quote tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Write your tweet: ")
			txt := c.ReadLine()
			c.Print("Write your quote: ")
			quote := c.ReadLine()
			Tweet := domain.NewTweetQuote(user, txt, quote)
			var err error
			var idNew int
			idNew, err = t.PublishTweet(Tweet)
			if err != nil && err.Error() == "user is required" {
				c.Print("User is required, try again")
			} else if err != nil && err.Error() == "text is required" {
				c.Print("Text is required, try again")
			} else if err != nil && err.Error() == "text and user are required" {
				c.Print("text and user are required, try again")
			} else if err != nil && err.Error() == "tweet must be less than 140 chars" {
				c.Print("tweet must be less than 140 chars")
			} else {
				c.Print("Tweet sent with id ", idNew)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishImageTweet",
		Help: "Publishes a image tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Write your tweet: ")
			txt := c.ReadLine()
			c.Print("Write your URL: ")
			url := c.ReadLine()
			Tweet := domain.NewTweetImage(user, txt, url)
			var err error
			var idNew int
			idNew, err = t.PublishTweet(Tweet)
			if err != nil && err.Error() == "user is required" {
				c.Print("User is required, try again")
			} else if err != nil && err.Error() == "text is required" {
				c.Print("Text is required, try again")
			} else if err != nil && err.Error() == "text and user are required" {
				c.Print("text and user are required, try again")
			} else if err != nil && err.Error() == "tweet must be less than 140 chars" {
				c.Print("tweet must be less than 140 chars")
			} else {
				c.Print("Tweet sent with id ", idNew)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweets := t.GetTweets()
			for index := 0; index < len(tweets); index++ {
				c.Println(tweets[index])
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "cleanLastTweet",
		Help: "Cleans last tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Cleans last tweet")
			t.CleanLastTweet()
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsByUser",
		Help: "Counts tweets by user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("The count of user's tweet is ", t.CountTweetsByUser(user))

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsByUser",
		Help: "Shows tweets by user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("User's tweets are: ", t.GetTweetsByUser(user))
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "followUser",
		Help: "Choose who to follow",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Enter the user you want to follow: ")
			user2 := c.ReadLine()
			t.Follow(user, user2)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "SeeMyFollows",
		Help: "My follows",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			follows := t.MyFollows(user)
			c.Print("Your follows: ")
			for i := 0; i < len(follows); i++ {
				c.Println(follows[i] + " ")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "GetTimeline",
		Help: "My timeline",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			timeline := t.GetTimeline(user)
			c.Print("Your timeline: ")
			for i := 0; i < len(timeline); i++ {
				c.Println(timeline[i])
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "TrendingTopic",
		Help: "Trending Topic",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tr := t.GetTrendingTopic()
			c.Print("These are the words more used: ")
			for i := 0; i < len(tr); i++ {
				c.Println(tr[i])
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "SendDirectMessage",
		Help: "Send direct message",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Enter the user you want to send a message: ")
			user2 := c.ReadLine()
			c.Print("Enter the message: ")
			msg := c.ReadLine()
			t.SendDirectMessage(user, user2, msg)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "GetAllDirectMessages",
		Help: "Get all direct message",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			directMessages := t.GetAllDirectMessages(user)
			for i := 0; i < len(directMessages); i++ {
				c.Println(directMessages[i])
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "GetUnreadedDirectMessages",
		Help: "Get unreade direct messages",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			directMessages := t.GetUnreadedDirectMessages(user)
			for i := 0; i < len(directMessages); i++ {
				c.Println(directMessages[i])
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "ReadDirectMessage",
		Help: "Read direct message",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter message's  id: ")
			id := c.ReadLine()
			id2, _ := strconv.Atoi(id)
			message := t.ReadDirectMessage(id2)
			c.Println(message)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "Retwit",
		Help: "Retwit",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Enter twit's  id: ")
			id := c.ReadLine()
			id2, _ := strconv.Atoi(id)
			message := t.Retweet(user, id2)
			c.Println(message)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "AddToFavs",
		Help: "Add to favs",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Enter twit's  id: ")
			id := c.ReadLine()
			id2, _ := strconv.Atoi(id)
			message := t.AddToFavs(user, id2)
			c.Println(message)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "AddToFavs",
		Help: "Add to favs",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Enter twit's  id: ")
			id := c.ReadLine()
			id2, _ := strconv.Atoi(id)
			message := t.AddToFavs(user, id2)
			c.Println(message)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "GetUserFavs",
		Help: "Get user favs",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter a user: ")
			user := c.ReadLine()
			favs := t.GetUserFavs(user)
			c.Println(favs)
			return
		},
	})

	shell.Run()
}
