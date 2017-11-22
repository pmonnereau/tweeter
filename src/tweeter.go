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
			c.Print("Ingrese su usuario: ")
			user := c.ReadLine()
			c.Print("Ingrese su tweet: ")
			txt := c.ReadLine()
			Tweet := domain.NewTweet(user, txt)
			var err error
			err = service.PublishTweet(Tweet)
			if err != nil && err.Error() == "user is required" {
				c.Print("User is required")
			} else {
				c.Print("Tweet sent \n")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweet := service.GetTweet()
			c.Println(tweet.User + " " + tweet.Text)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "cleanTweet",
		Help: "Cleans tweeter",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Cleans your tweets")
			service.CleanTweet()
			return
		},
	})

	shell.Run()
}
