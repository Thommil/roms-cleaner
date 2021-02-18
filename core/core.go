package core

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// COPY_DIR is the folder in roms to cleaned roms if copy mode is enabled
const COPY_DIR string = "cleaned"

// IMAGE_DIR is the folder in roms to store images
const IMAGE_DIR string = "images"

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

// Region defines a region code
type Region string

const (
	// EUROPE region
	EUROPE Region = "EUR"
	// USA region
	USA Region = "USA"
	// JAPAN region
	JAPAN Region = "JPN"
)

// System defines a system
type System struct {
	ID string
}

// Systems defines supported System list
var Systems = []System{
	{"3do"},
	{"amiga"},
	{"amstradcpc"},
	{"apple2"},
	{"arcade"},
	{"atari800"},
	{"atari2600"},
	{"atari5200"},
	{"atari7800"},
	{"atarilynx"},
	{"atarist"},
	{"atarijaguar"},
	{"atarijaguarcd"},
	{"atarixe"},
	{"colecovision"},
	{"c64"},
	{"intellivision"},
	{"macintosh"},
	{"xbox"},
	{"xbox360"},
	{"msx"},
	{"neogeo"},
	{"ngp"},
	{"ngpc"},
	{"n3ds"},
	{"n64"},
	{"nds"},
	{"nes"},
	{"gb"},
	{"gba"},
	{"gbc"},
	{"gc"},
	{"wii"},
	{"wiiu"},
	{"pc"},
	{"sega32x"},
	{"segacd"},
	{"dreamcast"},
	{"gamegear"},
	{"genesis"},
	{"mastersystem"},
	{"megadrive"},
	{"saturn"},
	{"psx"},
	{"ps2"},
	{"ps3"},
	{"ps4"},
	{"psvita"},
	{"psp"},
	{"snes"},
	{"pcengine"},
	{"wonderswan"},
	{"wonderswancolor"},
	{"zxspectrum"},
}

// Game keeps current state of a rom treatement
type Game struct {
	Title       string
	Source      string
	Destination string
	Image       string
	Warnings    []error
	Errors      []error
}

func init() {
	sort.Slice(Systems, func(i, j int) bool {
		return strings.Compare(Systems[i].ID, Systems[j].ID) < 0
	})
}

// GetRegion returns a region key baed on a key
func GetRegion(key string) (Region, error) {
	upperKey := strings.ToUpper(key)
	if upperKey == "US" || upperKey == "USA" || upperKey == "U" {
		return USA, nil
	}
	if upperKey == "JAPAN" || upperKey == "JPN" || upperKey == "J" {
		return JAPAN, nil
	}
	if upperKey == "EUR" || upperKey == "EU" || upperKey == "EUROPE" || upperKey == "WORLD" || upperKey == "E" {
		return EUROPE, nil
	}
	return EUROPE, fmt.Errorf("invalid region %s", key)
}

// GetSystem returns a system key based on a key or path
func GetSystem(keyOrPath string) (System, error) {
	key := strings.ToLower(keyOrPath)

	for _, system := range Systems {
		if system.ID == key {
			return system, nil
		}
	}

	key = strings.ToLower(filepath.Base(keyOrPath))

	for _, system := range Systems {
		if system.ID == key {
			return system, nil
		}
	}

	return System{}, fmt.Errorf("system %s not found", key)
}
