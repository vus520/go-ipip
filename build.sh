#!/bin/sh

file="ipip"_$(uname)
go build -ldflags "-w -s" -o $file

#build for Linux
cd $(brew info go| grep Cellar | awk '{print $1"/libexec/src"}')
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash;

cd -
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ipip_linux
