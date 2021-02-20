package scanner

import (
	"archive/zip"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/thommil/roms-cleaner/core"
)

type indexScanner struct {
}

// Name        string
// 	CloneOf     string
// 	RomOf       string
// 	Description string
// 	Year        string
// 	Roms        []string
// 	Merges      []string

func (indexScanner indexScanner) Scan(options core.Options, games []core.GameStatus, dat *core.DAT) (bool, error) {
	glog.V(1).Infof("Scan(options, %d games)", len(games))

	index, err := core.CreateIndex([]string{"CloneOf", "RomOf", "Year", "Merges"})
	if err != nil {
		return false, err
	}
	defer index.Close()

	for _, datGame := range dat.Games {
		err := index.Add(datGame.Name, datGame)
		if err != nil {
			glog.Error(err)
		}
	}

	var found int = 0

	// Try with name
	for i, game := range games {
		if game.Title == "" {
			gameName := strings.Replace(filepath.Base(game.Source), filepath.Ext(game.Source), "", 1)
			results, err := index.Search(gameName)
			if err != nil {
				game.Errors = append(game.Errors, err)
			}
			if len(results) > 0 && results[0].Score > 0.9 {
				for _, datGame := range dat.Games {
					if datGame.Name == results[0].Key {
						glog.V(2).Infof("found %s", datGame.Description)
						games[i].Title = datGame.Description
						found++
						break
					}
				}
			}
		} else {
			found++
		}
	}

	if found == len(games) {
		return true, nil
	}

	// Not complete, try with roms
	for i, game := range games {
		if game.Title == "" {
			r, _ := zip.OpenReader(game.Source)
			if err != nil {
				glog.Error(err)
			} else if r != nil && r.File != nil && len(r.File) > 0 {
				firstRom := r.File[0].Name
				results, err := index.Search(firstRom)
				if err != nil {
					game.Errors = append(game.Errors, err)
				}
				if len(results) > 0 && results[0].Score > 0.9 {
					for _, datGame := range dat.Games {
						if datGame.Name == results[0].Key {
							glog.V(2).Infof("found %s", datGame.Description)
							games[i].Title = datGame.Description
							found++
							break
						}
					}
				}
				r.Close()
			}
		}
	}

	glog.V(1).Infof("found %d/%d", found, len(games))

	return found == len(games), nil
}

func init() {
	registerScanner("index", indexScanner{})
}
