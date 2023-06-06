package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	dirFlag := flag.String("dir", ".github", "Directory to search for GitHub Workflow files. Defaults to the `.github` directory")
	oldFlag := flag.String("old", "", "The string to replace.")
	newFlag := flag.String("new", "", "The string to replace it with.")

	flag.Parse()

	if *oldFlag == "" || *newFlag == "" {
		log.Fatal("old and new flags must be set")
	}

	// find all .yml files in the directory
	files, err := FindFiles(*dirFlag)
	if err != nil {
		log.Fatal(err)
	}

	// "s|actions/checkout@v2|actions/checkout@v3|"
	substitution := fmt.Sprintf("s|%s|%s|", *oldFlag, *newFlag)

	// loop through each file and run sed to update the checkout action
	for _, file := range files {
		err := ReplaceString(file, substitution)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Done")
}

func ReplaceString(file string, substitution string) error {
	cmd := exec.Command("sed", "-i", "-e", substitution, file)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("sed failed: %s\n%s", err.Error(), stderr.String())
	}

	return nil
}

func FindFiles(root string) ([]string, error) {
	var dirs []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".yml" {
			return nil
		}

		dirs = append(dirs, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dirs, nil
}
