package util

import (
	"fmt"
	"log"
	"os"
)

func VerificaDirectory(dir string) error {
	_, err := directory(dir)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return err
	}

	return nil
}

func directory(pathName string) (string, error) {
	if directoryExists(pathName) {
		os.RemoveAll(pathName)
	}
	err := os.Mkdir(pathName, 0700)
	if err != nil {
		return "", fmt.Errorf("%c", err)
	}
	return pathName, nil
}

func directoryExists(path string) bool {
	dir, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return dir.IsDir()
}

func IsDirectory(dir string) bool {
	return directoryExists(dir)
}

func Files(dir string) (string, error) {
	var files string
	var nextFiles string

	path, err := os.ReadDir(dir)

	if err != nil {
		return "", err
	}

	for _, f := range path {
		if len(files) > 0 {
			nextFiles = ",."
		} else {
			nextFiles = "."
		}

		files += nextFiles + SeparatorPath() + dir + SeparatorPath() + f.Name()
	}

	return files, nil
}

func SeparatorPath() string {
	return string(os.PathSeparator)
}
