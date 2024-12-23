#!/bin/bash

# Check if a directory name is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <directory-name>"
    exit 1
fi

DIR_NAME=$1
MAIN_GO_CONTENT='package main

import (
	"advent-of-code-2024/lib"
)

const TestString string = ``

/*
	Part 1 Notes

*/

func solvePart1(input string) int {
	return 0
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(1, solvePart1(TestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	// dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
'
# Create the directory
mkdir -p "$DIR_NAME"

# Create an empty data.txt file
touch "$DIR_NAME/data.txt"

# Create main.go with the provided template
echo "$MAIN_GO_CONTENT" > "$DIR_NAME/main.go"
