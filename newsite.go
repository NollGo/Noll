package main

import (
	"io"
	"io.github.nollgo/assets"
	"io/fs"
	"os"
	"path/filepath"
)

const themePath = "theme"

func newSite(path string) error {
	if path == "" {
		path = "."
	}

	dir, err := assets.Dir.ReadDir(themePath)
	if err != nil {
		return err
	}

	// create dir
	_ = os.MkdirAll(path, 0755)

	if err = write(path, themePath, dir); err != nil {
		return err
	}
	return nil
}

func write(path string, parent string, entries []fs.DirEntry) error {
	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			if err := os.Mkdir(entryPath, 0755); err != nil {
				return err
			}
			dir, err := assets.Dir.ReadDir(filepath.Join(parent, entry.Name()))
			if err != nil {
				return err
			}
			if err = write(entryPath, filepath.Join(parent, entry.Name()), dir); err != nil {
				return err
			}
		} else {
			f, err := assets.Dir.Open(filepath.Join(parent, entry.Name()))
			if err != nil {
				return err
			}
			bytes, err := io.ReadAll(f)
			if err != nil {
				return err
			}

			if err = os.WriteFile(entryPath, bytes, 0644); err != nil {
				return err
			}
		}
	}
	return nil
}
