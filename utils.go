package main

import (
	"bufio"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
)

// Notifier continuously scans the given channel, and notifies when updates are
// found.
func Notifier(c <-chan string, f func(string, ...interface{}) (int, error)) {
	for {
		if len(c) > 0 {
			x := <-c
			f(x)
		}
	}
}

// GetFeed returns a feed from a given rss url. 
func GetFeed(url string) gofeed.Feed {

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	HandleErr(err)
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
		if t == "" {
			continue
		}
		sources = append(sources, t)
	}
	return sources, nil
}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Compare compares the states of prev & new Feeds and returns a slice of
// gofeed.Item.
func Compare(prev, new gofeed.Feed) []gofeed.Item {

	// [5, 4, 3, 2, 1]
	// [7, 6, 5, 4, 3]
	var newItems []gofeed.Item
	for i := len(new.Items); i >= 0; i-- {
		var temp []gofeed.Item
		for j := len(prev.Items); j >= 0; j-- {
			if new.Items[i] != prev.Items[j] {
				
			}
		}
	}
	return newItems
}

// SendNotifs loops through n and sends titles to c.
func SendNotifs(n []gofeed.Item, c chan<- string) {
	for _, v := range n {
		c <- v.Title
	}
}
