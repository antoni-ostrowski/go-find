package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	var targetPath string
	if len(args) < 1 {
		targetPath = "."
	} else {
		targetPath = args[0]
	}

	var query string
	if len(args) < 2 {
		query = ""
	} else {
		query = args[1]
	}

	err := filepath.WalkDir(targetPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if strings.Contains(path, query) {
			os.Stdout.Write([]byte(path + "\n"))
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
