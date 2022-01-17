package main

import (
	"bufio"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
)

// Notifier continuously scans the given channel, and notifies when updates are
// found.
func Notifier(c <-chan string, f func(string, ...interface{}) (int, error)) {
	for {
		if len(c) > 0 {
			x := <-c
			f(x + "\n")
		}
	}
}

// GetFeed returns a feed from a given rss url.
func GetFeed(url string) gofeed.Feed {

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	HandleErr(err, url)
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

func HandleErr(err error, s ...interface{}) {
	if err != nil {
		for _, v := range s {
			log.Println(v)
		}
		log.Fatal(err)
	}
}

// Compare compares the states of prev & new Feeds and returns a slice of
// gofeed.Item that is comprised of new items not found in prev.
func Compare(prev, new gofeed.Feed) []gofeed.Item {
	var newItems []gofeed.Item
	for _, newVal := range new.Items {
		maxTries := len(prev.Items) - 1 
		for tries, prevVal := range prev.Items {
			if tries >= maxTries {
				newItems = append(newItems, *newVal)
			}
			if newVal.Content == prevVal.Content &&
				newVal.Updated == prevVal.Updated {
				break
			}
		}
	}
	return newItems
}

// SendNotifs loops through n and sends titles to c.
func SendNotifs(n []gofeed.Item, c chan<- string) {
	for _, v := range n {
		c <- FormatItem(v)
	}
}

// FormatItem contains the final formatting of the output notifications
func FormatItem(i gofeed.Item) string {
	publishTime := i.PublishedParsed.String()
	return  publishTime + "\n" + i.Title + "\n" + i.Link + "\n"
}
