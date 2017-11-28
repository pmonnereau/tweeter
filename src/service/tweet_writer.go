package service

import (
	"os"

	"github.com/tweeter/src/domain"
)

type TweetWriter interface {
	WriteTweet(domain.Tweet)
}

type MemoryTweetWriter struct {
	Tweets []domain.Tweet
}

type FileTweetWriter struct {
	file *os.File
}

func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.OpenFile("tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666)
	writer := new(FileTweetWriter)
	writer.file = file
	return writer
}

func (t *MemoryTweetWriter) WriteTweet(tweet domain.Tweet) {
	t.Tweets = append(t.Tweets, tweet)

}

func (write *FileTweetWriter) WriteTweet(tweet domain.Tweet) {
	if write.file != nil {
		byteSlice := []byte(tweet.PrintableTweet() + "\n")
		write.file.Write(byteSlice)
	}
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	memoryTweetWriter := new(MemoryTweetWriter)
	memoryTweetWriter.Tweets = make([]domain.Tweet, 0)
	return memoryTweetWriter
}
