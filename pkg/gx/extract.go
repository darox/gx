package gx

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// write extraction to host filesystem
func (e Extraction) Write(target string) error {

	// get absolute path
	absolute, err := filepath.Abs(target)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", absolute, e.File.Name())

	// check if file already exists
	fileInfo, _ := os.Stat(path)
	if fileInfo != nil {
		log.Fatalf("File %s already exists", path)
	}

	// create file
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	// copy file from extraction to target path
	io.Copy(f, e.File)

	defer f.Close()

	return nil
}

// extract the specified source and returns it as part of Extraction struct
func (r *Repository) Extract() (Extraction, error) {
	e := Extraction{}
	f, err := r.FileSystem.Open(r.Source)
	if err != nil {
		return e, err
	}
	e.File = f
	return e, nil
}
