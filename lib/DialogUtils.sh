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

# Maybe placeholder for other install methods
_dialog_menu() {
    dialog --title "$WTF_TITLE" --checklist "Choose install methods" 15 60 1 \
        1 "WTF installation" on  2>$TEMP_PATH
}

# Print wait value
_dialog_wait() {
    if [ "$1" != "" ];then
        { for i in $(seq 1 100) ; do
            echo $i
            sleep 0.001
        done

        echo 100; } | dialog --backtitle "$WTF_TITLE" \
                         --gauge "$1" 6 60 0
    fi
}

_dialog_agree() {
    res=$(dialog --title "$WTF_TITLE"  --yesno "$1" 6 60 2>/tmp/ans)
}

_dialog_title() {
    dialog --title "$WTF_TITLE"  --msgbox "$1" 6 60
}

_dialog_notice() {
    dialog --title "$1"  --msgbox "$1" 6 60
}




