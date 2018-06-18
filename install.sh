#!/bin/bash

sudo apt-get install -y tree
sudo apt-get install -y libgtk-3-dev

wget https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz
tar -C /usr/local -zxvf go1.8.1.linux-amd64.tar.gz

mkdir -p ~/go/src
echo "export GOPATH=$HOME/go" >> ~/.barhrc
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
echo "export PATH=\$PATH:\$GOROOT/bin " >> ~/.bashrc

echo "export CGO_ENABLED=1" >> ~/.bashrc
echo "export CXX=g++" >> ~/.bashrc
echo "export CC=gcc" >> ~/.bashrc


source  ~/.bashrc

#test
/usr/local/go/bin/go version

mkdir -p src/github.com/andlabs
wget https://github.com/andlabs/ui/archive/master.zip
unzip master.zip -d src/github.com/andlabs
mv src/github.com/andlabs/ui-master/ src/github.com/andlabs/ui


exit 0
