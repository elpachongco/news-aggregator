# News Aggregator

Scans rss feed from multiple sources for news and delivers it via a
notification.

## Usage

Make sure Go is installed and present on PATH environment variable.
[Go Installation instruction](https://go.dev/doc/install)

In the terminal:

	``` 
	$ cd $HOME
	$ git clone https://github.com/elpachongco/news-aggregator
	$ cd news-aggregator
	```

Then add a new file, named `sources.txt`. Put all rss feed urls to be scanned in the file.

Example:

	```
	$ echo "https://reddit.com/r/memes/new/.rss" > sources.txt
	```

In the same directory, run the program with:

	```
	$ go run .
	```

## Architecture

## TODO

- Add content parser for html.
- Make program accept command line arguments.
- Add more tests.
