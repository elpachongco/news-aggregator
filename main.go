// main.go is the main program that orchestrates the operation of the program by
// making use of the utility functions in utils.go

package main

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

const (
	// Time to wait before re-scanning feeds.
	scanInterval = time.Millisecond * 3000
	// Path where the sources file is found.
	sourceListPath = "sources.txt"
	// Buffer size for new notifications.
	notifierBufferSize = 50
)

func main() {
	var sources []string
	sources, err := GetSources(sourceListPath)
	HandleErr(err)

	notifBuffer := make(chan string, notifierBufferSize)
	go Notifier(notifBuffer, fmt.Printf)

	feeds := make([]gofeed.Feed, len(sources))
	prevFeeds := make([]gofeed.Feed, len(sources))

	for {
		for feedIndex, url := range sources {
			feeds[feedIndex] = GetFeed(url)
			newPosts := Compare(prevFeeds[feedIndex], feeds[feedIndex])
			go SendNotifs(newPosts, notifBuffer)
			prevFeeds[feedIndex] = feeds[feedIndex]
		}
		time.Sleep(scanInterval)
	}
}
