package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/pacna/goober/internal/scraper"
	"github.com/pacna/goober/internal/utility"
)


func Run() {
	userInput := flag.String("input", "", "Specifies the URL of the web page you want to scrape");
	zipDest := flag.String("outdir", "", "Defines the destination path where all scraped images will be zipped")

	flag.Parse()

	if utility.IsEmpty(*userInput) {
		fmt.Println("Missing required flag: --input")
		flag.Usage()
		os.Exit(1)
	}

	scraper.New(*zipDest, *userInput).Configure().Run()
	
}