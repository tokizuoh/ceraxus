#!/bin/zsh

source ./setup/.env

readonly MOD_FILE="go.mod"

if [ -e $MOD_FILE ]; then
    printf "%s already exists.\n" $MOD_FILE
    exit 0
fi

command=$(printf "go mod init github.com/%s/%s" $GITHUB_ID $MODULE)
eval $command

printf "Finished!: go mod init github.com/%s/%s \n" $GITHUB_ID $MODULE