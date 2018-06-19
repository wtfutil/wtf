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

    # set go path
    if [ -z ${GOPATH+x} ];then
        export GOPATH=$HOME/go
        export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    fi

    if [ "$OS" == "Linux" ];then
        _dialog_notice "$WTF_CLEAN"
        go clean

        _dialog_notice "$WTF_RELES"

        # Get wtf
        _dialog_notice "$WTF_DEP"

        go get -u github.com/senorprogrammer/wtf
        cd $GOPATH/src/github.com/senorprogrammer/wtf

        # Install
        go install -ldflags="-X main.version=$(git describe --always --abbrev=6) -X main.date=$(date +%FT%T%z)"

    else
        spinner_start "$WTF_DEP"; sleep 0.1
        if [ ! -d "$GOPATH/src/github.com/senorprogrammer/wtf" ];then
            go get -u github.com/senorprogrammer/wtf
        else
            echo "Directory already exists"
        fi
        spinner_stop $?

        cd $GOPATH/src/github.com/senorprogrammer/wtf

        spinner_start "$WTF_CLEAN"; sleep 0.1k
        go clean
        spinner_stop $?

        spinner_start "Install wtf"; sleep 0.1
        go install -ldflags="-X main.version=$(git describe --always --abbrev=6) -X main.date=$(date +%FT%T%z)"
        spinner_stop $?

        spinner_start "Build wtf"; sleep 0.1
        go build -o bin/wtf
        spinner_stop $?

        if [ ! -f "/bin/wtf" ];then
            sudo ln -s $GOPATH/bin/wtf /usr/bin/wtf
        fi
    fi

    if hash wtf ;then echo "Installed!";fi
}

