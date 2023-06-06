package main

import (
	"flag"
	"log"

	"github.com/KasonBraley/repo-scripts/pkg/goutils"
)

func main() {
	dirFlag := flag.String("dir", ".", "directory to search for go.mod files. Defaults to current directory")
	pkg := flag.String("pkg", "", "package to update")

	flag.Parse()

	// find all subdirectories that contain a go.mod file
	dirs, err := goutils.FindDirectoriesWithGoMod(*dirFlag)
	if err != nil {
		log.Fatalf("Error finding directories with go.mod: %v", err)
	}

	// loop through each directory and run go get -u and go mod tidy
	for _, dir := range dirs {
		if !goutils.ModuleContainsPackage(dir, *pkg) {
			log.Printf("Skipping. Module %s does not contain package %s", dir, *pkg)
			continue
		}

		if err := goutils.UpdateGoPackage(dir, *pkg); err != nil {
			log.Fatalf("Error updating package in directory %s: %v", dir, err)
		}

		if err := goutils.GoModTidy(dir); err != nil {
			log.Fatalf("Error running go mod tidy in directory %s: %v", dir, err)
		}
	}
}
