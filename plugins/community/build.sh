#!/bin/zsh
rm -rf community.so
go build -buildmode=plugin -gcflags="all=-N -l"