package main

import (
	// "fmt"
	"fmt"
	// "log"
	"testing"

	"github.com/mmcdole/gofeed"
)

func TestCompare(t *testing.T) {

	// Initialize test variables
	items := []*gofeed.Item{}

	for k := 19; k >= 0; k-- {
		var item gofeed.Item
		item.Title = fmt.Sprintf("Test Item #%d", k)
		item.Content = fmt.Sprintf("Test Item content #%d", k)
		items = append(items, &item)
	}

	// `items` now has length 20 and contains the values:
	// ["Test Item content #19"..."Test Item content #0"]
	// represented in the comments as
	// [19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1,0]

	var prevFeed, newFeed gofeed.Feed
	prevFeed.Title = "Prev Feed"
	newFeed.Title = "New Feed"

	// Basic test creates two feeds of the same length: prevFeed and newFeed.
	// prevFeed.Items == [5,4,3,2,1,0],
	// newFeed.Items == [8,7,6,5,4,3].
	// Basic test ensures that Compare(prevFeed, newFeed) == [8,7,6].
	t.Run("Basic test",
		func(t *testing.T) {

			prevFeed.Items = items[14:19]
			newFeed.Items = items[11:17]
			want := items[11:14]

			got := Compare(prevFeed, newFeed)
			for k, v := range got {
				if v.Title != want[k].Title {
					t.Errorf("Compare basic test fail, Want %s, got %s",
					want[k].Title, v.Title)
				}
			}
		})

	// Tests whether Compare will work properly when newFeed contains items
	// that are new (not present in prevFeed).
	t.Run("Compare 2 feeds of same length with no common items.",
		func(t *testing.T) {
			prevFeed.Items = items[9:17]
			newFeed.Items = items[0:8]

			got := Compare(prevFeed, newFeed)
			want := items[0:8]

			for k, v := range want {
				if v.Title != got[k].Title {
					t.Errorf("Error Compare() 2 feeds of same length" +
					"with no common items. Want %s, Got %s.", v.Title, got[k].Title)
				}
			}
	})

	t.Run("Compare feeds where no new items are present.",
		func(t *testing.T) {
			prevFeed.Items = items[0:8]
			newFeed.Items = items[0:8]
			a := Compare(prevFeed, newFeed)
			got := len(a)
			want := 0
			if got != want {
				t.Errorf("Error Compare() 2 feeds where no new item is present."+
					"Want %d, Got %d", want, got)
			}

		})
}

/* TODO

 */
