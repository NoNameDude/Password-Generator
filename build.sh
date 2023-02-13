#!/bin/bash
# defining a variable
echo "Wheres the file you want to compile: (/Path/to/directory)"
# reading input 
read directory

compile()
{
    cd "$directory"
    go build "Main.go" 
}

compile
echo "Finished compiling"
