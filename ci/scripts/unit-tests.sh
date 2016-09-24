#!/bin/bash
# resource-marathon-deploy unit-test.sh

set -e -x

# The code is located in /resource-marathon-deploy
echo "List whats in the current directory"
ls -lat 

# Setup the gopath based on current directory.
export GOPATH=$PWD

# Now we must move our code from the current directory ./resource-marathon-deploy to $GOPATH/src/github.com/JeffDeCola/resource-marathon-deploy
mkdir -p src/github.com/JeffDeCola/
cp -R ./resource-marathon-deploy src/github.com/JeffDeCola/.

# All set and everything is in the right place for go
echo "Gopath is: " $GOPATH
echo "pwd is: " $PWD
cd src/github.com/JeffDeCola/resource-marathon-deploy

# RUN unit_tests
# go test -v -cover ./...