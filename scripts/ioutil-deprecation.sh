#!/bin/bash

# This script is used to replace all instances of `ioutil` with `io` and `os` equivalents.
# Example of how to run this script:
# multi-gitter run ./scripts/ioutil-deprecation.sh --ssh-auth --concurrent 20 --log-level=debug --git-type=cmd -R KasonBraley/kit -m 'refactor: Remove deprecated `io/ioutil`' -B remove-ioutil --assignees KasonBraley

gofmt -w -r 'ioutil.Discard -> io.Discard' .
gofmt -w -r 'ioutil.NopCloser -> io.NopCloser' .
gofmt -w -r 'ioutil.ReadAll -> io.ReadAll' .
gofmt -w -r 'ioutil.ReadFile -> os.ReadFile' .
gofmt -w -r 'ioutil.TempDir -> os.MkdirTemp' .
gofmt -w -r 'ioutil.TempFile -> os.CreateTemp' .
gofmt -w -r 'ioutil.WriteFile -> os.WriteFile' .
gofmt -w -r 'ioutil.ReadDir -> os.ReadDir ' . # (note: returns a slice of os.DirEntry rather than a slice of fs.FileInfo)

goimports -w .
