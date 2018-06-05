#!/usr/bin/env bash

# ============================================================ #
# wtf install script
#
# Copyright (C) 2018 Cyberfee aka deltaxflux
#
# Maintainer: cyberfee / 2018-06-05 11:42
#
# License: MIT
# ============================================================ #

wtf_installer() {
    local -r BRANCH=$(git rev-parse --abbrev-ref HEAD)

    # set go path
    if [ -z ${GOPATH+x} ];then
        export GOPATH=$HOME/go
        export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    fi

    if ! ping -q -w 1 -c 1 8.8.8.8  &> /dev/null; then
        _dialog_notice "No internet connection found"
    fi

    # Get wtf
    go get -u github.com/senorprogrammer/wtf
    cd $GOPATH/src/github.com/senorprogrammer/wtf

    # get dependencies
    go get -v ./...

    # Install wtf
    go install -ldflags="-X main.version=$(git describe --always --abbrev=6)_$BRANCH -X main.date=$(date +%FT%T%z)"

    if [ ! -f "/bin/wtf" ];then
        sudo ln -s $GOPATH/bin/wtf /bin/wtf
    fi
}
