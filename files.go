package main

import (
	"os"
	"path"
	"strings"
)

func readFile(filePath string) ([]byte, string) {
	fullPath := filePath

	// If file path is relative, get the full path
	if strings.HasPrefix(filePath, "./") {
		execDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		fullPath = path.Join(execDir, filePath)
	}

	file, err := os.ReadFile(fullPath)
	if err != nil {
		panic(err)
	}

	return file, fullPath
}
