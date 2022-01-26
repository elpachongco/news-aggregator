package main

import (
	// "fmt"
	"fmt"
	// "log"
	"testing"

	"github.com/mmcdole/gofeed"
)

func TestCompare(t *testing.T) {
	// Make two feeds of same length with different but overlapping items.
	// prev := [5, 4, 3, 2, 1, 0]
	// new := [8, 7, 6, 5, 4, 3]
	// Compare(prev, new) == [8, 7, 6]
	t.Run("Basic Comparison", func(t *testing.T) {
		maxNum := 8
		sliceLen := 6
		items := []*gofeed.Item{}

		for k := maxNum; k >= 0; k-- {
			var item gofeed.Item
			item.Title = fmt.Sprintf("Test Item #%d", k)
			item.Content = fmt.Sprintf("Test Item content #%d", k)
			items = append(items, &item)
		}

		var prevFeed gofeed.Feed
		prevFeed.Title = "Prev Feed"

		var newFeed gofeed.Feed
		newFeed.Title = "New Feed"

		// slice `items` now contains the values:
		// [8,7,6,5,4,3,2,1,0]
		
		prevFeed.Items = items[maxNum-sliceLen+1:] // == [5,4,3,2,1,0]
		newFeed.Items = items[0:sliceLen] // == [8,7,6,5,4,3]

		want := items[0: maxNum-sliceLen+1]
		a := Compare(prevFeed, newFeed) // == [8,7,6]
		for k, v := range a {
			if v.Title != want[k].Title {
				t.Errorf("Basic comparison test fail")
			}
		}



	})
}
