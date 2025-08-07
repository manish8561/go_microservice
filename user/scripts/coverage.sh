#!/usr/bin/env bash

set -e

# http://stackoverflow.com/a/21142256/2055281

echo "mode: atomic" > coverage.txt
// This script finds all directories containing Go files and runs tests with coverage.
// It appends the coverage results to a file named coverage.txt.
for d in $(find ./* -maxdepth 10 -type d); do
    if ls $d/*.go &> /dev/null; then
        go test  -coverprofile=profile.out -covermode=atomic $d
        if [ -f profile.out ]; then
            echo "$(pwd)"
            cat profile.out | grep -v "mode: " >> coverage.txt
            rm profile.out
        fi
    fi
done

