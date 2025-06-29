package main

import (
	"errors"
	"os"
	"strings"
)

func ScanDir(dirPath string) ([]string, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return []string{}, errors.New("Unable to access directory " + dirPath)
	}

	files, err := f.Readdir(0)
	if err != nil {
		return []string{}, errors.New("unable to access direcotry " + dirPath)
	}

	result := make([]string, len(files))
	for i := range files {
		result[i] = strings.Split(files[i].Name(), ".")[0]
	}

	return result, nil
}
