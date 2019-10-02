#!/usr/bin/env bash

go build -o $GOPATH/bin/pkgviewer main.go

p=./plantuml.jar
if [ ! -f $p ];then
wget -O $p https://nchc.dl.sourceforge.net/project/plantuml/plantuml.jar
fi


p=$GOPATH/third/plantuml.jar

if [ ! -f $p ];then
cp ./plantuml.jar $p
fi