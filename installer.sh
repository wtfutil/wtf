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

source lang/en.sh
source conf/Config.sh
source lib/InitUtils.sh
source lib/DialogUtils.sh
source lib/PrintUtils.sh

source lib/installer/WtfInstaller.sh

# ============================================================ #
# ==================== < WTF Parameters > ==================== #
# ============================================================ #

readonly WtfPath=$(dirname $(readlink -f "$0"))
readonly WtfTempPath="/tmp/wtfspace"

readonly WtfInstallerVersion=1
readonly WtfInstallerRevision=1

Init() {
    _dialog_title "$WTF_HELLO"
    _dialog_agree "$WTF_AGREE"

    _dialog_wait "$WTF_BACKUP"
    _init_files

}

Gui() {
    _dialog_menu
    IN=$(cat $TEMP_PATH | head -n 1)

    for i in $(seq 1 $(echo $IN | wc -w));do
        case $(echo $IN | cut -d" " -f${i}) in
            1) wtf_installer;;
            *) _print_error "$WTF_OPTION";;
        esac

    done
}

main() {
    Init
    Gui
}

main
