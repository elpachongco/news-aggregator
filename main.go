package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"time"
)

const (
	scanInterval   = time.Millisecond * 1000
	notifBuffSize  = 20
	sourceListPath = "sources.txt"
)

type Post struct {
	Title   string
	Content string
}

// A feed is comprised of posts assumed to be in chronological order.
type Feed []Post

func main() {
	var sources []string
	sources, err := GetSources(sourceListPath)
	HandleErr(err)

	feeds := make([]gofeed.Feed, len(sources))
	prevFeeds := make([]gofeed.Feed, len(sources))

	newPosts := make(chan string, notifBuffSize)
	go Notifier(newPosts, fmt.Printf)

	for {
		for k, url := range sources {
			feeds[k] = GetFeed(url)
			newPostKey := Compare(prevFeeds[k], feeds[k])

			if newPostKey > 0 && k != 0 {
				for _, v := range feeds {
					newPosts <- v.Title
				}
			}
		}
		prevFeeds = feeds
		time.Sleep(scanInterval)
	}
}

