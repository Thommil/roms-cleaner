package core

import (
	"fmt"

	"github.com/golang/glog"
)

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

var regions = map[Region][]string{
	USA:    {"usa", "slus", "scus", "slus", "[u]", "[us]", "ntsc-u"},
	EUROPE: {"europe", "eur", "world", "sles", "sces", "[e]", "world", "eng", "english", "pal"},
	JAPAN:  {"japan", "jpn", "slps", "slpm", "[j]", "en-ja", "ja", "ntsc-j"},
}

var regionIndex Index

// GetRegion returns a region key baed on a key
func GetRegion(key string) ([]Region, error) {
	glog.V(2).Infof("GetRegion(%s)", key)

	var err error
	if regionIndex == nil {
		regionIndex, err = CreateIndex(nil)

		if err != nil {
			glog.Error(err)
			return nil, err
		}

		for k, region := range regions {
			//regionS := struct{ List []string }{List: region}
			regionIndex.Add(string(k), region)
			if err != nil {
				glog.Error(err)
				return nil, err
			}
		}
	}

	result, err := regionIndex.Search(key)

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	if len(result) == 0 {
		glog.Errorf("invalid region %s", key)
		return nil, fmt.Errorf("invalid region %s", key)
	}

	foundRegions := make([]Region, len(result))
	for _, r := range result {
		foundRegions = append(foundRegions, Region(r.Key))
	}

	glog.V(2).Infof("%s -> %v", key, foundRegions)

	return foundRegions, nil
}
