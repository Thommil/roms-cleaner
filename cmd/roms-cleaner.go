package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thommil/roms-cleaner/cleaner"
	"github.com/thommil/roms-cleaner/core"
	"github.com/thommil/roms-cleaner/scanner"
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

func main() {
	var err error
	var region core.Region
	var system core.System
	var imagesDir, romsDir string

	// Folders
	flag.StringVar(&imagesDir, "img", "", "images directory matching roms filenames")

	// System
	flag.Func("system", "if folder is not recognized, force system (n64, nes...)", func(sys string) error {
		system, err = core.GetSystem(sys)
		return err
	})

	// Region
	flag.Func("region", "privileged region: EUR | USA | JPN (default EUR)", func(reg string) error {
		region, err = core.GetRegion(reg)
		return err
	})

	// Switches
	keepClones := flag.Bool("clones", false, "force to keep clones (default false)")
	copyMode := flag.Bool("copy", false, "copy mode, do not delete source content (default false)")
	failOnError := flag.Bool("fail", false, "fail on error (default false)")

	// Parse
	flag.Usage = usage
	flag.Parse()
	romsDir = flag.Arg(0)

	// Roms folder
	if romsDir == "" {
		fmt.Fprintln(os.Stderr, "ERROR: missing roms directory")
		usage()
		os.Exit(1)
	}

	romsDir, err = filepath.Abs(romsDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if _, err = os.Stat(romsDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "ERROR: roms directory not found: %s\n", romsDir)
		os.Exit(1)
	}
	if system.ID == "" {
		if system, err = core.GetSystem(romsDir); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
			os.Exit(1)
		}
	}

	// Images folder
	imagesDir, err = filepath.Abs(imagesDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
	if _, err = os.Stat(imagesDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", fmt.Errorf("image directory not found: %s", imagesDir))
		os.Exit(1)
	}

	// Init options & game list
	options := core.Options{
		Region:      region,
		System:      system,
		ImagesDir:   imagesDir,
		RomsDir:     romsDir,
		KeepClones:  *keepClones,
		CopyMode:    *copyMode,
		FailOnError: *failOnError,
	}
	games := make([]core.Game, 1000)

	// Scan
	if err = scanner.Scan(options, games); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(2)
	}

	// Clean
	if err = cleaner.Clean(options, games); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(3)
	}
}
