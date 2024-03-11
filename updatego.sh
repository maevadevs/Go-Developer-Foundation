#!/bin/bash

######################################################################################
# SCRIPT:       Allow to quickly update go version in go.work and go.mod to system   #
# CALL SIGN:    bash updatego.sh                                                     #
# CALL EXAMPLE: bash updatego.sh                                                     #
######################################################################################

# Get Current Directory
CURRENT_DIR=$PWD;

# Get system go version
go_version=`go version | { read _ _ v _; echo ${v#go}; }`
echo "System Go version: $go_version"

echo "Replace Go version in go.work..."
# Check if a go.work file exist
if test -f ./go.work; then
    # Replace go version with current
    sed -i -e "s/go [0-9]\+\.[0-9]\+\.[0-9]\+/go ${go_version}/g" ./go.work
fi

echo "Replace Go version in go.mod..."
# Check if a go.work file exist
for obj in *; do
    if [ -d "$obj" ]; then
        cd $obj
        # Check if a go.mod file exist
        if test -f ./go.mod; then
            # Replace go version with current
            sed -i -e "s/go [0-9]\+\.[0-9]\+\.[0-9]\+/go ${go_version}/g" ./go.mod
        fi
        cd ..
    fi
done

echo "All done!"
