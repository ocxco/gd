#!/usr/local/bin/fish
#
# wrapper for fish
#
# place this file at ~/.config/fish/functions/gd.fish
#

function gd
    if set output (_gd $argv)
        cd $output
    else
        for line in $output
            echo $line
        end
    end
end