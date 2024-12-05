package main

import (
	"advent-of-code-2024/lib"
	"fmt"
	"regexp"
	"strconv"
)

const TestString1 string = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const TestString2 string = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

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

	Similar to part one, but now do() enables all next mul() operations and don't() disables all next mul() operations

	Regex should parse these in order, so if I capture all mul() opersations as well as do() and don't() then I should be able to iterate through the list
*/

func solvePart2(input string) int {
	total := 0

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else {
			if enabled {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				total += x * y
			}
		}
	}

	return total
}

func main() {
	lib.AssertEqual(161, solvePart1(TestString1))
	lib.AssertEqual(48, solvePart2(TestString2))

	dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	result2 := solvePart2(dataString)
	fmt.Println(result2)
}
