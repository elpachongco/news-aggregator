package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"time"
)

const (
	// Time to wait before re-scanning feeds.
	scanInterval = time.Millisecond * 1000
	// Path where the sources file is found.
	sourceListPath = "sources.txt"
	// Buffer size for new notifications.
	notifierBufferSize = 20
)

func main() {
	var sources []string
	sources, err := GetSources(sourceListPath)
	HandleErr(err)

	feeds := make([]gofeed.Feed, len(sources))
	prevFeeds := make([]gofeed.Feed, len(sources))

	notifBuffer := make(chan string, notifierBufferSize)
	go Notifier(notifBuffer, fmt.Printf)

	for {
		for k, url := range sources {
			feeds[k] = GetFeed(url)
			newPosts := Compare(prevFeeds[k], feeds[k])

			if k > 0 {
				go SendNotifs(newPosts, notifBuffer)
			}
		}
		prevFeeds = feeds
		time.Sleep(scanInterval)
	}
}
