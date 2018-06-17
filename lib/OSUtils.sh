#!/usr/bin/env bash

unameOut="$(uname -s)"
case "${unameOut}" in
        Linux*)     : "Linux";;
		Darwin*)    : "Mac";;
		CYGWIN*)    : "Cygwin";;
		*)          : "Unknown";;
esac

declare -r OS="$_"
