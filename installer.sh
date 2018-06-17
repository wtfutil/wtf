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
source lib/OSUtils.sh
source lib/SpinnerUtils.sh
source lib/InitUtils.sh
source lib/DialogUtils.sh
source lib/PrintUtils.sh

source lib/installer/WtfInstaller.sh

# ============================================================ #
# ==================== < WTF Parameters > ==================== #
# ============================================================ #

readonly WtfTempPath="/tmp/wtfspace"
readonly WtfInstallerVersion=1
readonly WtfInstallerRevision=1

## need quick fux
Init() {
    clear
    if [ "$OS" == "Linux" ];then
        _dialog_title "$WTF_HELLO"
        _dialog_agree "$WTF_AGREE"

        _dialog_wait "$WTF_BACKUP"

        _init_files
    else
        echo -e "\033[32mWTF gui installer\033[0m \n"
        spinner_start "$WTF_BACKUP"; sleep 0.1
        _init_files
        spinner_stop $?
    fi

}

Gui() {
    if [ "$OS" == "Linux" ];then
        _dialog_menu
        IN=$(cat $TEMP_PATH | head -n 1)

        for i in $(seq 1 $(echo $IN | wc -w));do
            case $(echo $IN | cut -d" " -f${i}) in
                1) wtf_installer; break;;
                *) _print_error "$WTF_OPTION";;
            esac

        done
    else
        wtf_installer
    fi

}

main() {
    Init
    Gui
}

main
