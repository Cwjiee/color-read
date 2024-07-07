#!/bin/sh

bin="clread"
go build -o $bin
rm -f $HOME/go/bin/$bin
mv $bin $HOME/go/bin/
