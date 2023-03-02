package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"testing"
)

func TestQwe(t *testing.T) {
	filepath.Walk("main.go", func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
