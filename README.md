# News Aggregator

Scans feed from multiple sources (rss, html) for news and delivers it via a
notification.

## Goals & Non Goals
### Be responsive, not fast
The program should aim to quickly report new findings but should not hog the
system resources because of too frequent and unnecessarily fast update times.

### Be simple, not universal
The program is meant to only accomodate rss feeds. If the webpage has none, use
other programs to make one don't attempt to create an integrated webpage parser.

## Architecture

Each feed has its own state. The state of the feed is determined by the titles
of posts the feed has. If a new post appears in the feed, then there's a new
title and its state has changed. 

The program scans the feed every X seconds, and checks if the feed's state is
the same as the previous. If it is, wait X seconds and scan again. If the state
is not the same, i.e. a new post has appeared, the program finds the part
that is the same as previous. Parts that aren't are sent to the buffered
notifier. The notifier is extensible: any notification method can be easily
implemented.

Each feed has this process and each process runs concurrently hence the reason
for a buffered notifier so that the feed scanners will just do their job of
scanning without waiting for the notifier to finish its job which may be limited
by the implementation of the notifier.

