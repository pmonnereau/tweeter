package service

import (
	"github.com/tweeter/src/domain"
)

type ChannelTweetWriter struct {
	writer TweetWriter
}

func NewChannelTweetWriter(writer TweetWriter) ChannelTweetWriter {
	channelTweetWriter := new(ChannelTweetWriter)
	channelTweetWriter.writer = writer
	return *channelTweetWriter
}

func (channel ChannelTweetWriter) WriteTweet(tweetsToWrite chan domain.Tweet, quit chan bool) {
	tweet, open := <-tweetsToWrite
	for open {
		channel.writer.WriteTweet(tweet)
		tweet, open = <-tweetsToWrite

	}

	quit <- true
}
