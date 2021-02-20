package scanner

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/thommil/roms-cleaner/core"
)

type manager struct {
	scanners map[string]Scanner
}

var instance = manager{
	scanners: make(map[string]Scanner),
}

func registerScanner(system string, scanner Scanner) {
	instance.scanners[system] = scanner
}

// Scanner defines scanners API
type Scanner interface {
	Scan(options core.Options, games []core.Game) error
}

// Scan is the main entry point for scanner package
func Scan(options core.Options, games []core.Game) error {
	glog.V(2).Infof("Scan(%#v)", options)

	if options.System.Scanners == nil || len(options.System.Scanners) == 0 {
		glog.Errorf("system %s not yet supported", options.System.ID)
		return fmt.Errorf("system %s not yet supported", options.System.ID)
	}

	// scanner, found := instance.scanners[options.System.ID]

	// if !found {
	// 	return fmt.Errorf("scanner for system %s not found", options.System.ID)
	// }

	// return scanner.Scan(options, games)
	return nil
}
