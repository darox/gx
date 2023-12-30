package gx

import (
	"io"
	"os"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/require"
)

func TestExtract(t *testing.T) {

	// Setup fs for testing
	fs := memfs.New()
	err := fs.MkdirAll("testdata", 0755)
	if err != nil {
		t.Fatal(err)
	}
	files := []string{
		"testdata/files/main.go",
		"testdata/files/bubbles.jpg",
		"testdata/README.md",
	}
	for _, file := range files {
		f, err := fs.Create(file)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		s, err := os.Open(file)
		if err != nil {
			t.Fatal(err)
		}
		defer s.Close()
		_, err = io.Copy(f, s)
		if err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		name    string
		url     string
		file    string
		wantErr bool
	}{
		{
			name:    "valid extraction",
			file:    "testdata/README.md",
			wantErr: false,
		},
		{
			name:    "invalid extraction",
			file:    "testdata/README.x",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				FileSystem: fs,
				Source:     tt.file,
			}
			_, err := r.Extract()
			if tt.wantErr {
				require.Error(t, err, "Expected error for test case: %s", tt.name)
			} else {
				require.NoError(t, err, "Unexpected error for test case: %s", tt.name)
			}
		})
	}
}
