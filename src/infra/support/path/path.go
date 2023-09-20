package pathutil

import (
	"os"
	"path/filepath"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func CreateDirIfNotExists(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func CreateFileIfNotExists(path string) error {
	dir := filepath.Dir(path)
	err := CreateDirIfNotExists(dir)
	if err != nil {
		return err
	}

	if PathExists(path) {
		return nil
	}
	_, err = os.Create(path)
	return err
}
