#!/bin/zsh
rm -rf enterprise.so
go build -buildmode=plugin -gcflags="all=-N -l"