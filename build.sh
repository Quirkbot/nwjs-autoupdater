#!/bin/sh
rm -rf dist
mkdir -p dist/darwin/x64
mkdir -p dist/win32/ia32
mkdir -p dist/win32/x64
mkdir -p dist/linux/x64
cd src/nwjs-autoupdater/
rsrc -manifest nwjs-autoupdater.exe.manifest -o nwjs-autoupdater.syso
glide install
cd ../../
GOPATH=$(pwd) GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" nwjs-autoupdater
mv nwjs-autoupdater dist/linux/x64/autoupdater
GOPATH=$(pwd) GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" nwjs-autoupdater
mv nwjs-autoupdater dist/darwin/x64/autoupdater
GOPATH=$(pwd) GOOS=windows GOARCH=386 go build -ldflags "-s -w -H=windowsgui" nwjs-autoupdater
mv nwjs-autoupdater.exe dist/win32/ia32/autoupdater.exe
GOPATH=$(pwd) GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -H=windowsgui" nwjs-autoupdater
mv nwjs-autoupdater.exe dist/win32/x64/autoupdater.exe
