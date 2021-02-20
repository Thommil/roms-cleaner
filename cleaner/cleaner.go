package cleaner

import (
	"github.com/thommil/roms-cleaner/core"
)

// COPY_DIR is the folder in roms to cleaned roms if copy mode is enabled
const COPY_DIR string = "cleaned"

// IMAGE_DIR is the folder in roms to store images
const IMAGE_DIR string = "images"

type manager struct {
	cleaners map[string]Cleaner
}

var instance = manager{
	cleaners: make(map[string]Cleaner),
}

func registerCleaner(name string, cleaner Cleaner) {
	instance.cleaners[name] = cleaner
}

// Cleaner defines cleaners API
type Cleaner interface {
	Clean(options core.Options, games []core.GameStatus) (bool, error)
}

// Clean is the main entry point for cleaner package
func Clean(options core.Options, games []core.GameStatus) error {
	// cleaner, found := instance.cleaners[options.System.ID]

	// if !found {
	// 	glog.Errorf(cleaner for system %s not found", options.System.ID)
	// 	return fmt.Errorf("cleaner for system %s not found", options.System.ID)
	// }

	// //First if mode copy is active create dst folder
	// if options.CopyMode {
	// 	options.CleanedDir = filepath.Join(options.RomsDir, COPY_DIR)
	// 	if err := os.RemoveAll(options.CleanedDir); err != nil {
	// 		return err
	// 	}
	// 	if err := os.Mkdir(options.CleanedDir, fs.FileMode(0777)); err != nil {
	// 		return err
	// 	}
	// } else {
	// 	options.CleanedDir = options.RomsDir
	// }

	// //Create image folder if not found
	// if _, err := os.Stat(filepath.Join(options.CleanedDir, IMAGE_DIR)); os.IsNotExist(err) {
	// 	if err = os.Mkdir(filepath.Join(options.CleanedDir, IMAGE_DIR), fs.FileMode(0777)); err != nil {
	// 		return err
	// 	}
	// }

	// return cleaner.Clean(options, games)
	return nil
}
