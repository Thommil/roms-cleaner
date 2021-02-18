package scanner

import (
	"bytes"
	"encoding/gob"
	"encoding/xml"
	"fmt"
)

type xmlDAT struct {
	Games []struct {
		Name         string `xml:"name,attr"`
		CloneOf      string `xml:"cloneof,attr"`
		RomOf        string `xml:"romof,attr"`
		Description  string `xml:"description"`
		Year         string `xml:"year"`
		Manufacturer string `xml:"manufacturer"`
		Roms         []struct {
			Name   string `xml:"name,attr"`
			Merge  string `xml:"merge,attr"`
			Size   int    `xml:"size,attr"`
			CRC    string `xml:"crc,attr"`
			Status string `xml:"status,attr"`
		} `xml:"rom"`
		Driver struct {
			Status string `xml:"status,attr"`
		} `xml:"driver"`
	} `xml:"game"`
}

type DAT struct {
	Games []Game
}

type Game struct {
	Name        string
	CloneOf     string
	RomOf       string
	Description string
	Year        string
	Roms        []string
}

// FromXML builds a DAT instance from XML dat version
func (dat *DAT) FromXML(data []byte) error {
	var xmlDat xmlDAT
	if err := xml.Unmarshal(data, &xmlDat); err != nil {
		return err
	}

	dat.Games = make([]Game, len(xmlDat.Games))
	for _, game := range xmlDat.Games {
		roms := make([]string, len(game.Roms))
		for _, rom := range game.Roms {
			roms = append(roms, rom.Name)
		}

		dat.Games = append(dat.Games, Game{
			Name:        game.Name,
			CloneOf:     game.CloneOf,
			RomOf:       game.RomOf,
			Description: game.Description,
			Year:        game.Year,
			Roms:        roms,
		})
	}

	return nil
}

// Serialize DAT
func (dat *DAT) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(dat); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Deserialize DAT
func (dat *DAT) Deserialize(data []byte) error {
	if data == nil {
		return fmt.Errorf("nil data")
	}

	dec := gob.NewDecoder(bytes.NewBuffer(data))
	if err := dec.Decode(dat); err != nil {
		return err
	}

	if len(dat.Games) == 0 {
		return fmt.Errorf("no game found")
	}

	return nil
}
