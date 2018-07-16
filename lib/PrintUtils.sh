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

_print_string() {
    if [ "$1" != "" ];then
        echo -e "[\033[32m+\033[0m] $1" | tee -a $TMP_PATH
    fi
}

_print_error() {
    if [ "$1" != "" ];then
        echo -e "[\033[31m!\033[0m] \033[31m$1 \033[0m" | tee -a $TMP_PATH
    fi
}

_print_information() {
    if [ "$1" != "" ];then
        echo -e "[\033[31mI\033[0m] $1" | tee -a $TMP_PATH
    fi
}


