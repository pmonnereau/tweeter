package main

import (
	"github.com/abiosoft/ishell"
	"github.com/tweeter/src/domain"
	"github.com/tweeter/src/service"
)

func main() {
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type help to know commands \n")
	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Write your tweet: ")
			txt := c.ReadLine()
			Tweet := domain.NewTweet(user, txt)
			var err error
			var idNew int
			idNew, _ = service.TweetManager.PublishTweet(Tweet)
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
		Name: "showTweet",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweets := service.GetTweets()
			for index := 0; index < len(tweets); index++ {
				c.Println(tweets[index].User + " " + tweets[index].Text)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "cleanLastTweet",
		Help: "Cleans last tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Cleans your last tweet")
			service.CleanLastTweet()
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsByUser",
		Help: "Count tweets by user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("The count of user's tweet is ", service.CountTweetsByUser(user))

			return
		},
	})
	shell.Run()
}
