package gx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRepository(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		branch  string
		source  string
		wantErr bool
	}{
		{
			name:    "valid URL",
			url:     "https://github.com/darox/gx.git",
			branch:  "main",
			wantErr: false,
		},
		{
			name:    "invalid URL",
			url:     "://githu.com/daorx/gx.git",
			branch:  "main",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewRepository(tt.url, tt.branch, tt.source)
			if tt.wantErr {
				require.Error(t, err, "Expected error for test case: %s", tt.name)
			} else {
				require.NoError(t, err, "Unexpected error for test case: %s", tt.name)
			}
		})
	}
}

func TestRepository_Clone(t *testing.T) {
	tests := []struct {
		name    string
		repo    *Repository
		wantErr bool
		file    string
	}{
		{
			name: "successful clone and file exists",
			repo: &Repository{
				Url:    "https://github.com/darox/gx.git",
				Branch: "main",
			},
			wantErr: false,
			file:    "README.md",
		},
		{
			name: "successful clone and file does not exists",
			repo: &Repository{
				Url:    "https://github.com/darox/gx.git",
				Branch: "main",
			},
			wantErr: false,
			file:    "inexistent.md",
		},
		{
			name: "unsuccessful clone",
			repo: &Repository{
				Url:    "https://github.com/darox/inexistent.git",
				Branch: "main",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.repo.Clone()
			if tt.wantErr {
				require.Error(t, err, "Expected error for test case: %s", tt.name)
			} else {
				require.NoError(t, err, "Unexpected error for test case: %s", tt.name)

				// Check if a file exists in the cloned repository
				_, err = tt.repo.FileSystem.Stat("README.md")
				require.NoError(t, err, "Failed to access file in cloned repository for test case: %s", tt.name)
			}
		})
	}
}
