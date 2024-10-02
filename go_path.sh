#!/bin/bash

export ROOT_FOLDER="$(git rev-parse --show-toplevel)"
export GOPATH=${ROOT_FOLDER}/bin

# Folder contains your golang source codes
mkdir -p $GOPATH/src

# Folder contains the binaries when you install an go based executable
mkdir -p $GOPATH/bin

# Folder contains the Go packages you install
mkdir -p $GOPATH/pkg

# Folder contains the Github Source code for the repos you cloned
mkdir -p $GOPATH/src/github.com
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH:$GOBIN
