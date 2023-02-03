package main

import (
	"os"
	"path/filepath"
)

// MkdirFileFolderIfNotExists 创建文件所在的目录，如果这个目录不存在
func MkdirFileFolderIfNotExists(path string) error {
	pathDir := filepath.Dir(path)
	if _, err := os.Stat(pathDir); os.IsNotExist(err) {
		return os.MkdirAll(pathDir, os.ModePerm)
	}
	return nil
}
