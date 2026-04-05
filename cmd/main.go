package main

import (
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func main() {
	query := flag.String("q", "", "query to search for")
	targetPath := flag.String("p", ".", "path you want to search in")

	flag.Parse()

	output := make([]string, 0)

	err := filepath.WalkDir(*targetPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.Contains(path, *query) {
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
