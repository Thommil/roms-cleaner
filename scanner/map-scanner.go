package scanner

import (
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/thommil/roms-cleaner/core"
)

type arcadeScanner struct {
}

func (scanner arcadeScanner) Scan(options core.Options, games []core.GameStatus, dat *core.DAT) (bool, error) {
	glog.V(1).Infof("Scan(options, %d games)", len(games))

	var found int = 0
	for i, game := range games {
		name := strings.Replace(filepath.Base(game.Source), filepath.Ext(game.Source), "", 1)

		for _, datGame := range dat.Games {
			if datGame.Name == name {
				glog.V(2).Infof("found %s", datGame.Description)
				games[i].Title = datGame.Description
				found++
				break
			}
		}
	}

	glog.V(1).Infof("found %d/%d", found, len(games))

	return found == len(games), nil
}

func init() {
	registerScanner("map", arcadeScanner{})
}
