package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/thommil/roms-cleaner/scanner"
)

func zipFiles(filename string, files []string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(filename)
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}

func generateDats(folder string) error {
	var binFileList []string
	fmt.Printf("-> parsing DAT files in %s\n", folder)
	err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if d.Type().IsRegular() && filepath.Ext(d.Name()) == ".dat" {
			fmt.Printf("  -> DAT File found: %s\n", path)
			var dat scanner.DAT
			if data, err := os.ReadFile(path); err != nil {
				return err
			} else {
				if err = dat.FromXML(data); err != nil {
					return err
				}
				fmt.Printf("    -> DAT loaded, %d games found\n", len(dat.Games))

				if data, err = dat.Serialize(); err != nil {
					return err
				} else {
					binPath := filepath.Join(os.TempDir(), strings.Replace(filepath.Base(path), ".dat", ".bin", 1))
					if err = os.WriteFile(binPath, data, os.FileMode(0666)); err != nil {
						return err
					}
					binFileList = append(binFileList, binPath)
				}

			}

		}
		return nil
	})

	if err == nil {
		wd, _ := os.Getwd()
		zipPath := filepath.Join(wd, "dats.zip")
		if err := zipFiles(zipPath, binFileList); err != nil {
			return err
		}
		fmt.Printf("-> DAT dumped and zipped: %s\n", zipPath)
		return nil
	}

	return err
}

func main() {
	mode := flag.String("type", "dats", "indicate the type of ressources to generate: dats|")

	flag.Parse()

	switch *mode {
	case "dats":
		fmt.Printf("Mode dats:\n")
		folder := flag.Arg(0)
		if folder != "" {
			if folder, err := filepath.Abs(folder); err != nil {
				fmt.Fprintln(os.Stderr, "ERROR: ", err)
				os.Exit(1)
			} else {
				err = generateDats(folder)
				if err != nil {
					fmt.Fprintln(os.Stderr, "ERROR: ", err)
					os.Exit(2)
				}
			}
		} else {
			fmt.Fprintln(os.Stderr, "ERROR: missing xml dats directory")
			os.Exit(1)
		}
	}

}
