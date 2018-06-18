#!/bin/bash
# https://github.com/guomingyi/go.git

mypath=$PWD

export GOPATH=$mypath
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:/usr/local/go/bin

export PATH=$PATH:$mypath/build
export LD_LIBRARY_PATH=$mypath/c/lib/
