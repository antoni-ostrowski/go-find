package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
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

	output := make([]string, 0)

	err := filepath.WalkDir(targetPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.Contains(path, query) {
			output = append(output, path+"\n")
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	slices.Sort(output)
	os.Stdout.Write([]byte(strings.Join(output, "")))
}
