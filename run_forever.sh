#!/bin/bash

GO_APP_PATH=$1
APP_DIR=$(dirname "$GO_APP_PATH")

# Change to the directory of the Go application
cd "$APP_DIR"

while true
do
    go run $(basename "$GO_APP_PATH")
    if [ $? -ne 0 ]; then
        echo "Go application crashed with exit code $?. Restarting..." >&2
        sleep 1
    fi
done