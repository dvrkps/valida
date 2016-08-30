#!/bin/bash

file="coverage.txt"
mode="atomic"

echo "mode: $mode" > $file

for dir in $(go list ./...); do

  out=$(go test -race -coverprofile=profile.out -covermode=$mode $dir)
  echo $out
  if [[ $out == *FAIL* ]]; then
    rm -f profile.out
    rm -f $file
    exit 1
  fi

  if [ -f profile.out ]; then
    cat profile.out | grep -v "mode: $mode" >> $file 
    rm -f profile.out
  fi

done
