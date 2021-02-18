package scanner

import (
	"embed"
	"fmt"

	"github.com/thommil/roms-cleaner/core"
)

// go:embed embed
var embeddedFS embed.FS

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
	scanner, found := instance.scanners[options.System.ID]

	if !found {
		return fmt.Errorf("scanner for system %s not found", options.System.ID)
	}

	return scanner.Scan(options, games)
}
