package scraper

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pacna/goober/internal/utility"
	"github.com/pacna/goober/internal/zipper"
)

type Scraper struct {
	zipDest string
	inputURL string
	reader io.Reader
}

func New(zipDest string, inputURL string) *Scraper {
	return &Scraper{
		zipDest:  zipDest,
		inputURL: inputURL,
	}
}

func (s *Scraper) Configure() *Scraper {

	if  s.validate(s.inputURL) != nil {
		panic("Unable to process due to validation error")
	}

	body := s.getResponseBody(s.inputURL)

	if (body == nil) {
		panic("Unable to get response from URL")
	}

	s.reader = body

	return s
}

func (s *Scraper) validate(inputURL string) error {
	if utility.IsEmpty(inputURL) {
		return errors.New("URL is empty")
	}
	
	if !utility.IsHttpURL(inputURL) {
		return errors.New("invalid URL")
	}

	return nil
}

func (s *Scraper) getResponseBody(inputURL string) io.Reader {
	var buffer bytes.Buffer
	response, err := http.Get(inputURL)

	if err != nil {
		return nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil
	}

	buffer.ReadFrom(response.Body)

	body := io.NopCloser(&buffer)

	return body
}

func (s *Scraper) Run() {
	var uniqueImgURLs map[string]bool = make(map[string]bool)
	var imgURLs []string

	document, err := goquery.NewDocumentFromReader(s.reader)

	if err != nil {
		log.Fatal("no html content")
	}

	document.Find("img").Each(func(index int, imgContent *goquery.Selection) {
		imgSrc, isSrcEmpty := imgContent.Attr("src")
		dataImgSrc, isDataSrcEmpty := imgContent.Attr("data-src")

		if isSrcEmpty {
			uniqueImgURLs[strings.TrimSpace(imgSrc)] = true
		}

		if isDataSrcEmpty {
			uniqueImgURLs[strings.TrimSpace(dataImgSrc)] = true
		}
	})

	var imgURLsSize int64 = 0

	for imgURL := range uniqueImgURLs {
		fmt.Printf("%q\n", imgURL)
		imgURLsSize ++
		imgURLs = append(imgURLs, imgURL)

	}

	if !utility.IsEmpty(s.zipDest) {
		fmt.Println("\n\nZipping to folder")
		zipper.DownloadImages(s.zipDest, s.inputURL, imgURLs, imgURLsSize) 
	}
}

