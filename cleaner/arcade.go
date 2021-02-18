package cleaner

import "github.com/thommil/roms-cleaner/core"

type arcadeCleaner struct {
}

func (cleaner arcadeCleaner) Clean(options core.Options, games []core.Game) error {
	return nil
}

func init() {
	registerCleaner("arcade", arcadeCleaner{})
}
