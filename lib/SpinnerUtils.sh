#!/usr/bin/env bash
function _spinner() {
    local on_success="OK"
    local on_FAIL="FAILED"
    local white="\033[0;37m"
    local green="\033[01;32m"
    local red="\033[01;31m"
    local nc="\033[0m"

    case $1 in
        start)
            i=1
            sp='\|/-'
            delay=${SPINNER_DELAY:-0.25}
            counter=0
            timer=0

            while :
            do

                # Display seconds or minutes it takes
                if [[ ${SPINNER_TIMER} ]]; then
                    counter=$(echo "(${counter} + ${delay})" | bc)
                    timer=$(echo $counter/1 | bc)
                    if [[ "${timer}" -lt 60 ]]; then
                        timer="${timer} sec"
                    else
                        timer="$(echo ${timer} / 60 | bc) min "
                    fi
                    echo -en "\r[  ${sp:i++%${#sp}:1} ]  ${timer}\t${2}"
                else
                    echo -en "\r[  ${sp:i++%${#sp}:1} ]\t${2}"
                fi

                sleep ${delay}
            done

            ;;
        stop)
            if [[ -z ${3} ]]; then
                echo -en "[${red}${on_FAIL}${nc}]  Spinner is not running."
                exit 3
            fi

            kill $3 > /dev/null 2>&1

            echo -en "\r["
            if [[ $2 -eq 0 ]]; then
                echo -en "${green}${on_success}${nc}"
            else
                echo -en "${red}${on_FAIL}${nc}]"
                exit 3
            fi
            echo -e "]"
            ;;
    esac
}

function spinner_start {
    _spinner "start" "${1}" &
    _sp_pid=$!
    disown
}

function spinner_stop {
    _spinner "stop" $1 $_sp_pid
    unset _sp_pid
}
