#!/bin/bash

set -eou pipefail

printf "Update the system...\n";
sudo apt update

printf "Installing dependencies...\n";

if dpkg --get-selections | grep golang > /dev/null; then
    printf "Golang already installed!\n";
else
    printf "Installing golang...\n";
    wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz

    mkdir -p $HOME/go
    printf 'export GOPATH=$HOME/go' >> $HOME/.bashrc;
    printf 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> $HOME/.bashrc;
fi

# To execute Golang without sudo
if test -f "/etc/sudoers.d/env_keep"; then
    printf "Sudoers file exists!\n";
else
    printf 'Defaults    env_keep += "GOPATH"' | sudo tee --append /etc/sudoers.d/env_keep;
    printf "Sudoers file included with sucess!\n";
fi

printf "Completed dependencies installation!\n";

printf "Running golang application...\n";
go run main.go
