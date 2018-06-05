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

_init_files () {
    readonly DATE=$(date +"%m_%d_%Y_%T")
    dependencies=( "$WTF_GO_DIR" "$WTF_INSTALL_DIR" )

    if [ ! -d "$HOME/BackupWtf/${DATE}" ];then
        mkdir -p "$HOME/BackupWtf/${DATE}"
        exist=false
    else
        exist=true
    fi

    for i in "${dependencies[@]}";do
        if [ -f "$i" ] || [ -d "$i" ];then
            if [ "$exist" == "false" ];then
                cp -r $i "$HOME/BackupWtf/${DATE}"
            fi
            rm -rf $i
        fi
    done

    cp misc/.dialogrc $HOME
}
