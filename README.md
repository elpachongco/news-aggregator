# News Aggregator

Scans rss feed from multiple sources for news and delivers it via a
notification.

## Usage

Make sure golang is installed.

In the terminal:
$ cd $HOME
$ git clone https://github.com/elpachongco/news-aggregator
$ cd news-aggregator

Then add a new file, named `sources.txt`. Put all rss feed urls to be scanned in the file.

To run the program, in the same directory, run
$ go run .

## Architecture

## TODO

- Add content parser for html
