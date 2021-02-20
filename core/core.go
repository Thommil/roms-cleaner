package core

// Options defines global cleaner options
type Options struct {
	Region      Region
	System      System
	ImagesDir   string
	RomsDir     string
	CleanedDir  string
	KeepClones  bool
	CopyMode    bool
	FailOnError bool
}

// GameStatus keeps current state of a rom treatement
type GameStatus struct {
	Title       string
	Source      string
	Destination string
	Image       string
	Warnings    []error
	Errors      []error
}
