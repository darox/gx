package gx

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func hasFileName(path string) bool {
	base := filepath.Base(path)
	fmt.Println(base)
	return base != "." && base != "/"
}

// write extraction to host filesystem
func (e Extraction) Write(target string) error {

	// check if file already exists
	//_, err := os.Stat()

	var p string
	if hasFileName(target) {
		p = fmt.Sprintf("%s", target)
		fmt.Println(p)
	} else {
		p = fmt.Sprintf("%s%s", target, e.File.Name())
		fmt.Println(p)
	}

	// create file at the target path
	f, err := os.Create(p)

	if err != nil {
		return err
	}

	// copy file from extraction to target path
	io.Copy(f, e.File)

	return nil
}

// extract the specified source and returns it as part of Extraction struct
func (r *Repository) Extract() (Extraction, error) {
	e := Extraction{}

	if !strings.Contains(r.Source, "/") {
		f, err := r.FileSystem.Open(r.Source)
		if err != nil {
			return e, err
		}
		e.File = f
	}
	return e, nil
}
