package main

import (
	"flag"
	"log"

	"github.com/KasonBraley/repo-scripts/pkg/goutils"
)

func main() {
	dirFlag := flag.String("dir", ".", "directory to search for go.mod files. Defaults to current directory")
	version := flag.String("version", "", "version of Go to update to")

	flag.Parse()

	// find all subdirectories that contain a go.mod file
	dirs, err := goutils.FindDirectoriesWithGoMod(*dirFlag)
	if err != nil {
		log.Fatalf("Error finding directories with go.mod: %v", err)
	}

	// loop through each directory and run go get -u and go mod tidy
	for _, dir := range dirs {
		if err := goutils.UpdateGoVersion(dir, *version); err != nil {
			log.Fatalf("Error updating to go version %s in directory %s: %v", *version, dir, err)
		}

		if err := goutils.GoModTidy(dir); err != nil {
			log.Fatalf("Error running go mod tidy in directory %s: %v", dir, err)
		}
	}
}
