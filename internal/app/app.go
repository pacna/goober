package app

import (
	"flag"

	"github.com/pacna/goober/internal/scraper"
)


func Run() {
	userInput := flag.String("input", "", "user URL input");
	zipDest := flag.String("zipdest", "", "zip file and export to desire path")

	flag.Parse()

	scraper.New(*zipDest, *userInput).Configure().Run()
	
}