#!/bin/bash

mypath=$PWD
mkdir -p $mypath/src/github.com/andlabs
cd $mypath/src/github.com/andlabs

git clone https://github.com/andlabs/ui
git clone https://github.com/andlabs/libui.git
cd libui && git checkout alpha3.5
rm -rf build && mkdir build && cd build

cmake -DBUILD_SHARED_LIBS=OFF .. && make
cp out/libui.a $mypath/src/github.com/andlabs/ui/libui_linux_amd64.a

cd $mypath/src/github.com/andlabs/ui
go build .
