package gx

import (
	"net/url"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

// checks if a passed url is valid
func isValidUrl(urlStr string) error {
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return err
	}
	return err
}

// clone the repository to the memory filesystem.
func (r *Repository) Clone() error {
	// create new memory filesystem
	fs := memfs.New()

	storer := memory.NewStorage()

	// clone the repository to the filesystem
	_, err := git.Clone(storer, fs, &git.CloneOptions{
		Depth: 1,
		URL:   r.Url,
	})

	if err != nil {
		return err
	}

	r.FileSystem = fs

	return nil
}

// create struct to hold data related to the repository to extract from
func NewRepository(url string, branch string, source string) (Repository, error) {
	err := isValidUrl(url)
	if err != nil {
		return Repository{}, err
	}

	return Repository{
		Url:        url,
		Branch:     branch,
		Source:     source,
		FileSystem: nil,
	}, nil
}
