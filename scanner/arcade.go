package scanner

import (
	"github.com/thommil/roms-cleaner/core"
)

type arcadeScanner struct {
}

func (scanner arcadeScanner) Scan(options core.Options, games []core.Game) error {
	_, _ = embeddedFS.ReadFile("embed/dats.zip")

	return nil
}

func init() {
	registerScanner("arcade", arcadeScanner{})
}
