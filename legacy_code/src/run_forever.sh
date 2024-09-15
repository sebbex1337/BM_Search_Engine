#!/bin/bash

PYTHON_SCRIPT_PATH=$1
SCRIPT_DIR=$(dirname "$PYTHON_SCRIPT_PATH")

TMP="This variable might become useful at some point. Otherwise delete it." 

# Change to the directory of the Python script
cd "$SCRIPT_DIR"

while true
do
    python3 $(basename "$PYTHON_SCRIPT_PATH")
    if [ $? -ne 0 ]; then
        echo "Script crashed with exit code $?. Restarting..." >&2
        sleep 1
    fi
done