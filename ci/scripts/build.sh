#!/bin/bash
# resource-marathon-deploy build.sh

set -e -x

# The code is located in /resource-marathon-deploy
echo "List whats in the current directory"
ls -lat 
echo ""

# Setup the gopath based on current directory.
export GOPATH=$PWD

# Now we must move our code from the current directory ./resource-marathon-deploy to $GOPATH/src/github.com/JeffDeCola/resource-marathon-deploy
mkdir -p src/github.com/JeffDeCola/
cp -R ./resource-marathon-deploy src/github.com/JeffDeCola/.

# All set and everything is in the right place for go
echo "Gopath is: " $GOPATH
echo "pwd is: " $PWD
echo ""
cd src/github.com/JeffDeCola/resource-marathon-deploy

# Put the binary resource-marathon-deploy filename in /dist
go build -o dist/resource-marathon-deploy ./main.go

# cp the Dockerfile into /dist
cp ci/Dockerfile dist/Dockerfile

# Check
echo "List whats in the /dist directory"
ls -lat dist
echo ""

# Move what you need to $GOPATH/dist
# BECAUSE the resource type docker-image works in /dist.
cp -R ./dist $GOPATH/.
cp -R ./assets-go $GOPATH/dist/.
cp -R ./assets-bash $GOPATH/dist/.
cp  ./bin/tree $GOPATH/dist/.

cd $GOPATH
# Check whats here
echo "List whats in top directory"
ls -lat 
echo ""

# Check whats in /dist
echo "List whats in /dist"
ls -lat dist
echo ""