#!/bin/bash

function getShell() {
    shell=$(env | grep SHELL)
    if [[ $shell =~ "fish" ]]; then
        echo "fish"
    elif [[ $shell =~ "zsh" ]]; then
        echo "zsh"
    else
        echo "bash"
    fi
}

function getSystem() {
    system=$(uname -a)
    if [[ $system =~ "Darwin" ]]; then
        echo "mac"
    elif [[ $system =~ "Linux" ]]; then
        echo "linux"
    else
        echo "windows"
    fi
}

function installZsh() {
    if [[ -d ~/.oh-my-zsh/lib ]]; then
        echo "You are using oh-my-zsh, it will be install at ~/.oh-my-zsh/lib/"
        # use oh-my-zsh will auto load from lib
        cp ../wrapper/gd.zsh ~/.oh-my-zsh/lib/
        return 0
    elif [[ ! -d ~/.zsh/ ]]; then
        mkdir ~/.zsh/
    fi
    echo "It will be install at ~/.zsh/"
    cp ../wrapper/gd.zsh ~/.zsh/
    # add to shell config
    echo "source ~/.zsh/gd.zsh" >> ~/.zshrc
}

function installFish() {
    echo "You are using fish shell, it will be install at ~/.config/fish/functions/"
    if [[ ! -d ~/.config/fish/functions/ ]]; then
        mkdir -p ~/.config/fish/functions/
    fi
    cp ../wrapper/gd.fish ~/.config/fish/functions/
}

function installBash() {
    echo "You are using normal bash/sh, it will be install at ~/.bashrc"
    cat ../wrapper/gd.sh >> ~/.bashrc
}

function installWrapper() {
    echo "Install the shell wrapper"
    shell=$(getShell)
    case "$shell" in
        "fish")
            installFish
            ;;
        "zsh")
            installZsh
            ;;
        "bash")
            installBash
            ;;
    esac
}

# chdir，in case can not fount file
cd $(dirname $0)

# install start
system=$(getSystem)
bin="$(pwd)/$system/gd"
if [[ "$system" == "windows" ]]; then
    if [ -f ~/.config/gd/_gd ]; then
        echo "gd already installed"
        exit 1
    fi
    if [ ! -d ~/.config/gd ]; then
        mkdir -p ~/.config/gd/
    fi
    echo "Copy binary file to ~/.config/gd/_gd"
    cp "$bin" ~/.config/gd/_gd
    echo "Install the shell wrapper"
    cp ../wrapper/gd.sh ~/.config/gd/
    echo 'export PATH=$PATH:~/.config/gd/' >> ~/.bashrc
    echo "source ~/.config/gd/gd.sh" >> ~/.bashrc
else
    if [ -f /usr/local/bin/_gd ]; then
        echo "gd already installed"
        exit 1
    fi
    echo "Copy binary file to /usr/local/bin/_gd"
    cp "$bin" /usr/local/bin/_gd
    installWrapper
fi

echo "Install finished, you can use it when you start terminal next"
