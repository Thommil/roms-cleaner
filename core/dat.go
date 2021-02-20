package core

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"encoding/gob"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/golang/glog"
)

//go:embed embed/dats.zip
var datsArchive []byte

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
	Merges      []string
}

// FromXML builds a DAT instance from XML dat version
func (dat *DAT) FromXML(data []byte) error {
	var xmlDat xmlDAT
	if err := xml.Unmarshal(data, &xmlDat); err != nil {
		glog.Error(err)
		return err
	}

	dat.Games = make([]Game, 0, len(xmlDat.Games))
	for _, game := range xmlDat.Games {
		roms := make([]string, 0, 0)
		merges := make([]string, 0, 0)
		for _, rom := range game.Roms {
			if rom.Merge == "" {
				roms = append(roms, rom.Name)
			} else {
				merges = append(merges, rom.Name)
			}
		}

		dat.Games = append(dat.Games, Game{
			Name:        game.Name,
			CloneOf:     game.CloneOf,
			RomOf:       game.RomOf,
			Description: game.Description,
			Year:        game.Year,
			Roms:        roms,
			Merges:      merges,
		})
	}

	return nil
}

// FromMemory load current instance with embedded data based on system
func (dat *DAT) FromMemory(system string) error {
	reader, err := zip.NewReader(bytes.NewReader(datsArchive), int64(len(datsArchive)))
	if err != nil {
		glog.Error(err)
		return err
	}

	for _, file := range reader.File {
		if strings.Replace(file.Name, ".bin", "", 1) == system {
			datReader, err := file.Open()
			if err != nil {
				glog.Error(err)
				return err
			}
			defer datReader.Close()

			data := make([]byte, int(file.UncompressedSize))

			_, err = datReader.Read(data)
			if err != nil {
				glog.Error(err)
				return err
			}

			err = dat.Deserialize(data)
			if err != nil {
				glog.Error(err)
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("no entry found for %s", system)
}

// Serialize DAT
func (dat *DAT) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(dat); err != nil {
		glog.Error(err)
		return nil, err
	}

	return buf.Bytes(), nil
}

// Deserialize DAT
func (dat *DAT) Deserialize(data []byte) error {
	if data == nil {
		glog.Error("nil data")
		return fmt.Errorf("nil data")
	}

	dec := gob.NewDecoder(bytes.NewBuffer(data))
	if err := dec.Decode(dat); err != nil {
		glog.Error(err)
		return err
	}

	if len(dat.Games) == 0 {
		glog.Error("no game found")
		return fmt.Errorf("no game found")
	}

	return nil
}