#!bin/bash

goenv install 1.11.4
goenv global 1.11.4
goenv rehash

echo "export GOPATH=$HOME/go" >> ~/.bashrc
echo "PATH=$PATH:$GOPATH/bin" >> ~/.bashrc
