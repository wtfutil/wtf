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

# Window ratio
readonly NAME="Wtf"
readonly RATIO=4
readonly TEMP_PATH="/tmp/chooses"

# WTF variables
WTF_GO_DIR="$GOPATH/src/github.com/senorprogrammer/wtf"
WTF_INSTALL_DIR="/bin/wtf"

# Dialog
readonly DIALOG="$HOME/.dialogrc"
