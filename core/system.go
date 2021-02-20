package core

import (
	"fmt"
	"path/filepath"

	"github.com/golang/glog"
)

// System defines a system
type System struct {
	ID       string
	Terms    []string
	Scanners []string
	Cleaners []string
}

var systems = map[string]System{
	"3do":             {"3do", nil, nil, nil},
	"amiga":           {"amiga", nil, nil, nil},
	"amstradcpc":      {"amstradcpc", nil, nil, nil},
	"apple2":          {"apple2", nil, nil, nil},
	"arcade":          {"arcade", []string{"arcade", "mame", "fba", "fbn", "burn", "raine", "capcom", "m.a.m.e", "cps", "cps1", "cps2", "cps3"}, nil, nil},
	"atari800":        {"atari800", nil, nil, nil},
	"atari2600":       {"atari2600", nil, nil, nil},
	"atari5200":       {"atari5200", nil, nil, nil},
	"atari7800":       {"atari7800", nil, nil, nil},
	"atarilynx":       {"atarilynx", nil, nil, nil},
	"atarist":         {"atarist", nil, nil, nil},
	"atarijaguar":     {"atarijaguar", nil, nil, nil},
	"atarijaguarcd":   {"atarijaguarcd", nil, nil, nil},
	"atarixe":         {"atarixe", nil, nil, nil},
	"colecovision":    {"colecovision", nil, nil, nil},
	"c64":             {"c64", nil, nil, nil},
	"intellivision":   {"intellivision", nil, nil, nil},
	"macintosh":       {"macintosh", nil, nil, nil},
	"xbox":            {"xbox", nil, nil, nil},
	"xbox360":         {"xbox360", nil, nil, nil},
	"msx":             {"msx", nil, nil, nil},
	"neogeo":          {"neogeo", nil, nil, nil},
	"ngp":             {"ngp", nil, nil, nil},
	"ngpc":            {"ngpc", nil, nil, nil},
	"n3ds":            {"n3ds", nil, nil, nil},
	"n64":             {"n64", nil, nil, nil},
	"nds":             {"nds", nil, nil, nil},
	"nes":             {"nes", nil, nil, nil},
	"gb":              {"gb", nil, nil, nil},
	"gba":             {"gba", nil, nil, nil},
	"gbc":             {"gbc", nil, nil, nil},
	"gc":              {"gc", nil, nil, nil},
	"wii":             {"wii", nil, nil, nil},
	"wiiu":            {"wiiu", nil, nil, nil},
	"pc":              {"pc", nil, nil, nil},
	"sega32x":         {"sega32x", nil, nil, nil},
	"segacd":          {"segacd", nil, nil, nil},
	"dreamcast":       {"dreamcast", nil, nil, nil},
	"gamegear":        {"gamegear", nil, nil, nil},
	"genesis":         {"genesis", nil, nil, nil},
	"mastersystem":    {"mastersystem", nil, nil, nil},
	"megadrive":       {"megadrive", nil, nil, nil},
	"saturn":          {"saturn", nil, nil, nil},
	"psx":             {"psx", nil, nil, nil},
	"ps2":             {"ps2", nil, nil, nil},
	"ps3":             {"ps3", nil, nil, nil},
	"ps4":             {"ps4", nil, nil, nil},
	"psvita":          {"psvita", nil, nil, nil},
	"psp":             {"psp", nil, nil, nil},
	"snes":            {"snes", nil, nil, nil},
	"pcengine":        {"pcengine", nil, nil, nil},
	"wonderswan":      {"wonderswan", nil, nil, nil},
	"wonderswancolor": {"wonderswancolor", nil, nil, nil},
	"zxspectrum":      {"zxspectrum", nil, nil, nil},
}

var systemIndex Index

// GetSystem returns a system key based on a key or path
func GetSystem(keyOrPath string) (System, error) {
	glog.V(2).Infof("GetSystem(%s)", keyOrPath)

	var err error
	if systemIndex == nil {
		systemIndex, err = CreateIndex([]string{"Scanners", "Cleaners"})

		if err != nil {
			glog.Error(err)
			return System{}, err
		}

		for _, supportedSystem := range systems {
			systemIndex.Add(supportedSystem.ID, supportedSystem)
			if err != nil {
				glog.Error(err)
				return System{}, err
			}
		}
	}

	key := filepath.Base(keyOrPath)

	result, err := systemIndex.Search(key)

	if err != nil {
		glog.Error(err)
		return System{}, err
	}

	if len(result) == 0 || result[0].Score < 0.9 {
		glog.Errorf("system %s not found", key)
		return System{}, fmt.Errorf("system %s not found", key)
	}

	glog.V(2).Infof("%s -> %s", keyOrPath, result[0].Key)

	return systems[result[0].Key], nil
}
