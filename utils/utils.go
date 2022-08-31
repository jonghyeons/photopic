package utils

import (
	"os"
	"path/filepath"
)

func MakeDir(path string) error {
	jpgPath := filepath.Join(path, "/jpg")
	if !Exists(jpgPath) {
		err := os.MkdirAll(jpgPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	rawPath := filepath.Join(path, "/raw")
	if !Exists(rawPath) {
		err := os.MkdirAll(rawPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
