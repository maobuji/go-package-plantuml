#!/usr/bin/env bash

go build -o $GOPATH/bin/pkgviewer main.go

p=$GOPATH/third/plantuml.jar

if [ ! -f $p ];then
cp ./plantuml.jar $p
fi