package scanner

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/thommil/roms-cleaner/core"
)

type manager struct {
	scanners map[string]Scanner
}

var instance = manager{
	scanners: make(map[string]Scanner),
}

func registerScanner(name string, scanner Scanner) {
	instance.scanners[name] = scanner
}

// Scanner defines scanners API
type Scanner interface {
	// Scan implements scanning by each scanner returnin completion and error
	Scan(options core.Options, games []core.GameStatus, dat *core.DAT) (bool, error)
}

// Scan is the main entry point for scanner package
func Scan(options core.Options, games []core.GameStatus) error {
	glog.V(1).Infof("Scan(%#v)", options)

	if options.System.Scanners == nil || len(options.System.Scanners) == 0 {
		glog.Errorf("system %s not yet supported", options.System.ID)
		return fmt.Errorf("system %s not yet supported", options.System.ID)
	}

	err := filepath.WalkDir(options.RomsDir, func(path string, d fs.DirEntry, err error) error {
		if d.Type().IsRegular() && strings.Contains(options.System.Exts, filepath.Ext(d.Name())) {
			games = append(games, core.GameStatus{
				Source:   path,
				Errors:   make([]error, 0),
				Warnings: make([]error, 0),
			})
		}
		return nil
	})
	if err != nil {
		glog.Errorf("failed to list roms: %s", err)
		return fmt.Errorf("failed to list roms: %s", err)
	}

	var dat core.DAT
	if err := dat.FromMemory(options.System.ID); err != nil {
		return err
	}

	for _, scanner := range options.System.Scanners {
		complete, err := instance.scanners[scanner].Scan(options, games, &dat)

		if err != nil {
			return err
		}

		if complete {
			break
		}
	}

	return nil
}
