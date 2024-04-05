#!/bin/bash

######################################################################################
# SCRIPT:       Allow to quickly generate a new Go project as a Go module            #
# CALL SIGN:    bash init.sh -p <Project_Dir> -o <Out_File> -m <Module_Path_Dir>     #
# CALL EXAMPLE: bash init.sh -p hello_world -o hello -m example.com/proj/hello_world #
######################################################################################

# Get Current Directory
CURRENT_DIR=$PWD;

# Get the project name and module root arguments from the user
while getopts p:o:m: flag
do
    case "${flag}" in
        p) PROJECT_NAME=${OPTARG};;
        o) OUT_FILE=${OPTARG};;
        m) MODULE_ROOT_PATH=${OPTARG};;
    esac
done

# The new project's directory
PROJECT_DIR="$CURRENT_DIR/$PROJECT_NAME"

echo "";
echo "Project Name: $PROJECT_NAME";
echo "Module Root Path: $MODULE_ROOT_PATH";

echo "";
echo "Creating a new Go project: $PROJECT_DIR ..."

# Create directory for the project if it does not exist yet
if test -d $PROJECT_DIR; then
    echo "The project directory already exist. Skipping creating a new directory.";
    echo "";
else
    echo "Creating a new project directory...";
    echo "$PROJECT_DIR";
    mkdir -p "$PROJECT_DIR";
    echo "Done."
    echo "";
fi

# Move into the project directory
cd $PROJECT_DIR;

# Create a new Go module if it does not exist yet
if test -f "$PROJECT_DIR/go.mod"; then
    echo "The go.mod file already exist. Skipping creating a new go module.";
    echo "";
else
    echo "Creating a new Go module...";
    go mod init "$MODULE_ROOT_PATH/$PROJECT_NAME";
    echo "Done."
    echo "";
fi

# Create a new "src" folder if does not exist yet
if test -d "$PROJECT_DIR/src"; then
    echo "The src directory exist. Skipping creating a new src directory.";
    echo "";
else
    echo "Creating a new src directory...";
    mkdir -p "src";
    echo "Done."
    echo "";
fi

# Create a new "main.go" in "src" if does not exist yet
if test -f "$PROJECT_DIR/src/main.go"; then
    echo "The main.go file exist. Skipping creating a main.go file.";
    echo "";
else
    echo "Creating a new main.go file...";
    touch "src/main.go";
    # Add default placeholder contents
    echo "\
// Package
// *******
package main

// Imports
// *******
import \"fmt\"

// Functions
// *********

// This is the main entry of the application.
func main() {
    fmt.Println(\"Hello world!\")
}

// FOR WINDOWS:
//  To run:                 go run $PROJECT_NAME\src\main.go
//  To compile:             go build -o $PROJECT_NAME\bin\\$OUT_FILE.exe $PROJECT_NAME\src\main.go
//  To run after compile:   .\\$PROJECT_NAME\bin\\$OUT_FILE.exe
//  Compile + Run:          go build -o $PROJECT_NAME\bin\\$OUT_FILE.exe $PROJECT_NAME\src\main.go && .\\$PROJECT_NAME\bin\\$OUT_FILE.exe

// FOR LINUX:
//  To run:                 go run $PROJECT_NAME/src/main.go
//  To compile:             go build -o $PROJECT_NAME/bin/$OUT_FILE $PROJECT_NAME/src/main.go
//  To run after compile:   ./$PROJECT_NAME/bin/$OUT_FILE
//  Compile + Run:          go build -o $PROJECT_NAME/bin/$OUT_FILE $PROJECT_NAME/src/main.go && ./$PROJECT_NAME/bin/$OUT_FILE

" >> src/main.go;

    echo "Done."
    echo "";
fi


echo "All processes complete. Exiting."
exit
