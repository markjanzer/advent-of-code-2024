package main

import (
	"advent-of-code-2024/lib"
	"fmt"
	"regexp"
	"strconv"
)

const SmallTestString string = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

/*
	Part 1 Notes

	Parse the string for sequences that match mul(x,y)
	Multiply x and y and add to the total
*/

func solvePart1(input string) int {
	total := 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		total += x * y
	}

	return total
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(161, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	dataString := lib.GetDataString()
	result1 := solvePart1(dataString)
	fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
