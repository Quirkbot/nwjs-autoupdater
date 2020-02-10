#!/bin/sh
export GOPATH=$(pwd)
rm -rf dist
mkdir -p dist/darwin/x64
mkdir -p dist/win32/ia32
mkdir -p dist/win32/x64
mkdir -p dist/linux/x64
cd src/nwjs-autoupdater/
$GOPATH/bin/rsrc -manifest nwjs-autoupdater.exe.manifest -o nwjs-autoupdater.syso
glide install
cd ../../
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" nwjs-autoupdater
mv nwjs-autoupdater dist/linux/x64/autoupdater
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" nwjs-autoupdater
mv nwjs-autoupdater dist/darwin/x64/autoupdater
GOOS=windows GOARCH=386 go build -ldflags "-s -w -H=windowsgui" nwjs-autoupdater
mv nwjs-autoupdater.exe dist/win32/ia32/autoupdater.exe
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -H=windowsgui" nwjs-autoupdater
mv nwjs-autoupdater.exe dist/win32/x64/autoupdater.exe
