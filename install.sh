#!/bin/bash

sudo apt-get install -y tree
sudo apt-get install -y libgtk-3-dev
sudo apt-get install -y libusb-dev
sudo apt-get install -y libusb-1.0-0-dev
sudo apt-get install g++ cmake

# go build.
wget https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz
tar -C /usr/local -zxvf go1.8.1.linux-amd64.tar.gz

mkdir -p ~/go/src
echo "export GOPATH=\$HOME/go" >> ~/.barhrc
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
echo "export PATH=\$PATH:\$GOROOT/bin " >> ~/.bashrc
echo "export CGO_ENABLED=1" >> ~/.bashrc
echo "export CXX=g++" >> ~/.bashrc
echo "export CC=gcc" >> ~/.bashrc
echo "export PATH=\$PATH:~/tools/liteide/bin" >> ~/.bashrc
source  ~/.bashrc

#liteide
#wget https://svwh.dl.sourceforge.net/project/liteide/X25.2/liteidex25.2.linux-64.tar.bz2 -q -O ~/tools/liteidex25.2.linux-64.tar.bz2
#tar -C ~/tools/ -xvf ~/tools/liteidex25.2.linux-64.tar.bz2
#wget http://mirrors.kernel.org/ubuntu/pool/main/libp/libpng/libpng12-0_1.2.54-1ubuntu1_amd64.deb -q -O ~/tools/libpng12.deb
#dpkg -i ~/tools/libpng12.deb

