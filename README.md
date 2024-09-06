# Goober

Goober is a fast and easy-to-use command line web scraper that effortlessly extract images from web pages.

<img alt="Test passing" src="https://github.com/pacna/goober/workflows/Test/badge.svg" />

## Prerequisites

Before using Goober, make sure you have the following tools and components installed:

-   [Golang](https://golang.org/dl/)

## Installation

```bash
$ go install github.com/pacna/goober@latest
```

> make sure `~/go/bin` is in your path

## Usage

Run Goober and specify the web page you want to scrape:

```bash
$ goober --input https://www.google.com/ --zipdest /path/to/folder
```

### Flags

1. `--input`: Specifies the URL of the web page you want to scrape.
2. `--zipdest`: Defines the destination path where all scraped images will be zipped. (Optional)
