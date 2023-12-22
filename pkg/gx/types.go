package gx

import "github.com/go-git/go-billy/v5"

// holds values related to a git repository the extraction is being performed on.
type Repository struct {
	// Url is the url of the git repository.
	Url string
	// Branch is the branch name.
	Branch string
	// FileSystem is the filesystem that the repository is being extracted to.
	FileSystem billy.Filesystem
	// Source is the path to extract from in the git repository.
	Source string
}

// holds values related to the extraction of a git repository.
type Extraction struct {
	// Holds a file if source path is a file
	File billy.File
	// Holds a folder if source path is a folder
	Folder billy.Dir
}
