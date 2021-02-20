package core

import (
	"fmt"

	"strings"
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

// GetRegion returns a region key baed on a key
func GetRegion(key string) (Region, error) {
	//log.Debug("GetRegion", key)
	lkey := strings.TrimSpace(strings.ToLower(key))
	if strings.Contains("usa slus scus slus [u] [us] [usa]", lkey) {
		return USA, nil
	}
	if strings.Contains("japan jpn slps slpm [japan] [jpn] [j]", lkey) {
		return JAPAN, nil
	}
	if strings.Contains("europe eu e world sles sces [europe] [eur] [eu] [e] [world]", lkey) {
		return EUROPE, nil
	}
	return EUROPE, fmt.Errorf("invalid region %s", key)
}
