#!/bin/zsh
#
# wrapper for sh/bash
#
# @github.com/ocxco/gd

gd() {
    output=$(_gd $@)
    ret=$?
    if [[  "$output" == "" ]]; then
        exit 0
    elif [[ $ret -eq 0 ]]; then
        cd "$output"
    else
        echo "$output"
    fi

    unset output
    unset ret
}
