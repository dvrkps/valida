#!/usr/bin/env bash

set -e
echo "" > coverage.txt

if ls *.go &> /dev/null; then
    go test -race -coverprofile=coverage.txt -covermode=atomic
fi

for d in $(find ./* -maxdepth 10 -type d); do
    if ls $d/*.go &> /dev/null; then
        go test -race -coverprofile=profile.out -covermode=atomic $d
        if [ -f profile.out ]; then
            cat profile.out >> coverage.txt
            rm profile.out
        fi
    fi
done
