#!/bin/bash

sudo apt install -y golang-go

git clone https://github.com/syndbg/goenv.git ~/.goenv

echo 'export GOENV_ROOT=$HOME/.goenv' >> ~/.bashrc
echo 'export PATH=$GOENV_ROOT/bin:$PATH' >> ~/.bashrc
echo 'eval "$(goenv init -)"' >> ~/.bashrc
