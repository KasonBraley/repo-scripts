package goutils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// findDirectoriesWithGoMod searches for all directories in the given root directory
// that contain a go.mod file.
func FindDirectoriesWithGoMod(root string) ([]string, error) {
	var dirs []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && fileExists(filepath.Join(path, "go.mod")) {
			dirs = append(dirs, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dirs, nil
}

// GoModTidy runs 'go mod tidy' in the given directory.
func GoModTidy(dir string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go mod tidy failed: %s\n%s", err.Error(), stderr.String())
	}

	return nil
}

// UpdateGoPackage runs 'go get -u' in the given directory.
// Leaving the pkg argument empty will update all packages.
// Otherwise, only the given package will be updated.
func UpdateGoPackage(dir, pkg string) error {
	cmd := exec.Command("go", "get", "-u")
	if pkg != "" {
		cmd.Args = append(cmd.Args, pkg)
	}
	cmd.Dir = dir

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go get -u failed: %s\n%s", err.Error(), stderr.String())
	}

	return nil
}

func UpdateGoVersion(dir, version string) error {
	if version == "" {
		return fmt.Errorf("cannot proceed: version is empty")
	}

    // check the current go version first
	cmd := exec.Command("go", "list", "-m", "-f", "{{.GoVersion}}")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	goVersion := strings.TrimSpace(stdout.String())

	if goVersion == version {
		return fmt.Errorf("go version is already %s", version)
	}

	cmd = exec.Command("go", "mod", "edit", "-go", version)
	cmd.Dir = dir

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go mod edit -go failed: %s\n%s", err.Error(), stderr.String())
	}

	return nil
}

// ModuleContainsPackage returns true if the go.mod file in the given directory
// contains the given package.
func ModuleContainsPackage(dir, pkg string) bool {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Path}}", pkg)
	cmd.Dir = dir

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return false
	}

	// trim off the newline from 'go list'
	return strings.TrimSpace(stdout.String()) == pkg
}

// fileExists returns true if the given file exists and is a regular file.
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
