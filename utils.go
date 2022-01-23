// utils.go contains functions that is used by the main functions

package main

import (
	"bufio"
	"log"
	// "fmt"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

// Notifier continuously scans the given channel, and notifies when updates are
// found.
func Notifier(c <-chan string, f func(string, ...interface{}) (int, error)) {
	for {
		if len(c) > 0 {
			x := <-c
			f(x)
			continue
		}
		time.Sleep(time.Millisecond * 200)
	}
}

// SendNotifs loops through n and sends titles to c.
func SendNotifs(n []gofeed.Item, c chan<- string) {
	for _, v := range n {
		c <- FormatItem(v)
	}
}

// GetFeed returns a feed from a given rss url.
func GetFeed(url string) gofeed.Feed {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	HandleErr(err, url)
	// fmt.Println(feed.Title, len(feed.Items), feed.Items[0].Title)
	return *feed
}

// GetSources finds the sourc list given by path and returns a slice of url
// strings.
func GetSources(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	var sources []string
	// Reads file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" || t[0] == '#' {
			continue
		}
		sources = append(sources, t)
	}
	return sources, nil
}

// HandleErr is used for checking function return errors.
func HandleErr(err error, s ...interface{}) {
	if err != nil {
		for _, v := range s {
			log.Println(v)
		}
		log.Fatal(err)
	}
}

// Compare loops through the items of prev & new Feeds and returns a slice of
// gofeed.Item that is comprised of new items not found in prev.
// For each feed item in new, look for a match in prev.
// If none is found, then the item is new.
func Compare(prev, new gofeed.Feed) []gofeed.Item {
	var newItems []gofeed.Item
	for _, newVal := range new.Items {
		maxTries := len(prev.Items) - 1
		for tries, prevVal := range prev.Items {
			if newVal.Title == prevVal.Title {
				break
			}
			if tries == maxTries {
				newItems = append(newItems, *newVal)
			}
		}
	}
	return newItems
}

// GetNew detects new items in a feed based on time t.
// It is intended as a replacment for Compare.
// Main issue is that not all feeds have the publish and update time.
func GetNew(feed gofeed.Feed, t time.Duration) []gofeed.Item {
	var newItems []gofeed.Item
	// Any items published or updated after the epoch is considered new.
	epoch := time.Now().UTC().Add(-t).UnixMilli()
	for _, v := range feed.Items {
		published := v.PublishedParsed.UnixMilli()
		if published > epoch {
			newItems = append(newItems, *v)
		}
	}
	return newItems
}

// FormatItem contains the final formatting of the output notifications
func FormatItem(i gofeed.Item) string {
	publishTime := i.PublishedParsed.Local().String()
	content := i.Content
	maxContentSize := 100
	if len(i.Content) > maxContentSize {
		content = i.Content[:maxContentSize]
	}
	return publishTime + "\n" + i.Title + "\n" + content + "\n" + i.Link + "\n\n"
}
