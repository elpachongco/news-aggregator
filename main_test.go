package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/mmcdole/gofeed"
)


func TestGetFeed(t *testing.T) {
	var x gofeed.Feed
	x = GetFeed("https://hnrss.org/newcomments")
	fmt.Println(x)
}

func TestGetSources(t *testing.T) {

	x, err := GetSources("sources.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(x)
}

func TestCompare(t *testing.T) {
	
	previous := `
<?xml version="1.0" encoding="windows-1252"?>
<rss version="2.0">
  <channel>
    <title>FeedForAll Sample Feed</title>
    <description>RSS is a fascinating technology. The uses for RSS are expanding daily. Take a closer look at how various industries are using the benefits of RSS in their businesses.</description>
    <link>http://www.feedforall.com/industry-solutions.htm</link>
    <category domain="www.dmoz.com">Computers/Software/Internet/Site Management/Content Management</category>
    <copyright>Copyright 2004 NotePage, Inc.</copyright>
    <docs>http://blogs.law.harvard.edu/tech/rss</docs>
    <language>en-us</language>
    <lastBuildDate>Tue, 19 Oct 2004 13:39:14 -0400</lastBuildDate>
    <managingEditor>marketing@feedforall.com</managingEditor>
    <pubDate>Tue, 19 Oct 2004 13:38:55 -0400</pubDate>
    <webMaster>webmaster@feedforall.com</webMaster>
    <generator>FeedForAll Beta1 (0.0.1.8)</generator>
    <image>
      <url>http://www.feedforall.com/ffalogo48x48.gif</url>
      <title>FeedForAll Sample Feed</title>
      <link>http://www.feedforall.com/industry-solutions.htm</link>
      <description>FeedForAll Sample Feed</description>
      <width>48</width>
      <height>48</height>
    </image>
    <item>
      <title>Test Item 1</title>
	  <description>
		Test Item 1 Description
	  </description>
      <link>http://www.feedforall.com/law-enforcement.htm</link>
      <category domain="www.dmoz.com">Computers/Software/Internet/Site Management/Content Management</category>
      <comments>http://www.feedforall.com/forum</comments>
      <pubDate>Tue, 19 Oct 2004 11:08:56 -0400</pubDate>
    </item>
  </channel>
</rss>
`
	current :=
`
<?xml version="1.0" encoding="windows-1252"?>
<rss version="2.0">
  <channel>
    <title>FeedForAll Sample Feed</title>
    <description>RSS is a fascinating technology. The uses for RSS are expanding daily. Take a closer look at how various industries are using the benefits of RSS in their businesses.</description>
    <link>http://www.feedforall.com/industry-solutions.htm</link>
    <category domain="www.dmoz.com">Computers/Software/Internet/Site Management/Content Management</category>
    <copyright>Copyright 2004 NotePage, Inc.</copyright>
    <docs>http://blogs.law.harvard.edu/tech/rss</docs>
    <language>en-us</language>
    <lastBuildDate>Tue, 19 Oct 2004 13:39:14 -0400</lastBuildDate>
    <managingEditor>marketing@feedforall.com</managingEditor>
    <pubDate>Tue, 19 Oct 2004 13:38:55 -0400</pubDate>
    <webMaster>webmaster@feedforall.com</webMaster>
    <generator>FeedForAll Beta1 (0.0.1.8)</generator>
    <image>
      <url>http://www.feedforall.com/ffalogo48x48.gif</url>
      <title>FeedForAll Sample Feed</title>
      <link>http://www.feedforall.com/industry-solutions.htm</link>
      <description>FeedForAll Sample Feed</description>
      <width>48</width>
      <height>48</height>
    </image>
     <item>
      <title>Test Item 3</title>
	  <description>
		Test Item 3 Description
	  </description>
      <link>http://www.feedforall.com/law-enforcement.htm</link>
      <category domain="www.dmoz.com">Computers/Software/Internet/Site Management/Content Management</category>
      <comments>http://www.feedforall.com/forum</comments>
      <pubDate>Tue, 19 Oct 2004 11:08:56 -0400</pubDate>
    </item>
   <item>
      <title>Test Item 2</title>
	  <description>
		Test Item 2 Description
	  </description>
      <link>http://www.feedforall.com/law-enforcement.htm</link>
      <category domain="www.dmoz.com">Computers/Software/Internet/Site Management/Content Management</category>
      <comments>http://www.feedforall.com/forum</comments>
      <pubDate>Tue, 19 Oct 2004 11:08:56 -0400</pubDate>
    </item>
    <item>
      <title>Test Item 1</title>
	  <description>
		Test Item 1 Description	
	  </description>
      <link>http://www.feedforall.com/law-enforcement.htm</link>
      <category domain="www.dmoz.com">Computers/Software/Internet/Site Management/Content Management</category>
      <comments>http://www.feedforall.com/forum</comments>
      <pubDate>Tue, 19 Oct 2004 11:08:56 -0400</pubDate>
    </item>
  </channel>
</rss>
`

	pf, _ := gofeed.NewParser().ParseString(previous)
	nf, _ := gofeed.NewParser().ParseString(current)

	a := Compare(*pf, *nf)
	log.Println(a)
}
