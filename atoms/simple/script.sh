#!/bin/bash

# Check if exactly two arguments are provided
if [ $# -ne 2 ]; then
    echo "Usage: ./script.sh <arg1> <arg2>"
    exit 1
fi

# Extract the arguments
arg1=$1
arg2=$2

# Print the Hello World message with the arguments
echo "Hello, World! $arg1 $arg2"
