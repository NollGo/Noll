package main

import (
	"io.github.nollgo/assets"
	"io/fs"
	"os"
	"path/filepath"
)

func newSite(path string) error {
	if path == "" {
		path = "."
	}

	dir, err := assets.Dir.ReadDir("theme")
	if err != nil {
		return err
	}

	// create dir
	_ = os.Mkdir(path, 0755)

	for i := range dir {
		if dir[i].IsDir() {
			continue
		}

		if dir[i].Type() == fs.ModeSymlink {
			continue
		}

		filename := filepath.Join(path, dir[i].Name())
		if err := os.WriteFile(filename, []byte{}, 0644); err != nil {
			return err
		}

	}
	return nil
}
