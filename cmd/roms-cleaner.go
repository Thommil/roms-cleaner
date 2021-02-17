package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	helpMessage :=
		`Usage: roms-cleaner [options] roms-dir

Desc

Options:
`
	fmt.Println(helpMessage)
	flag.PrintDefaults()
}

func parseCommandline() {
	// Folders
	dstDir := flag.String("dst", "./out", "destination folder")
	imgDir := flag.String("img", "", "images directory matching roms filenames")

	// Filters
	flag.Func("region", "privileged region: EUR | USA | JPN (default EUR)", func(r string) error {

		return nil
	})

	// Scraper
	flag.Func("scraper", "Scraping service : screenscraper | thegamesdb (default screenscraper)", func(s string) error {

		return nil
	})
	user := flag.String("user", "", "user on scrapping service")
	password := flag.String("pwd", "", "password scrapping service")

	// Switches
	keepClones := flag.Bool("clones", false, "force to keep clones (default false)")
	failOnError := flag.Bool("fail", false, "fail on error (default false)")

	// Parse
	flag.Usage = usage
	flag.Parse()
	romsDir := flag.Arg(0)

	// Check
	if romsDir == "" {
		fmt.Fprintln(os.Stderr, "ERROR: missing roms directory")
		usage()
		os.Exit(1)
	}

	fmt.Println(*keepClones, *dstDir, *imgDir, *user, *password, *failOnError, romsDir)
}

func main() {
	fmt.Println("Parsing commend line...")
	parseCommandline()
}
