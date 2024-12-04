package main

import (
	"advent-of-code-2024/lib"
	"fmt"
	"strconv"
	"strings"
)

const SmallTestString string = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

/*
	Part 1 Notes

*/

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")

	safeReports := 0

	for _, line := range lines {
		fields := strings.Fields(line)
		report := []int{}
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			report = append(report, num)
		}
		if isSafeReport(report) {
			fmt.Println(report)
			safeReports++
		}
	}
	return safeReports
}

func isSafeReport(report []int) bool {
	direction := levelsDirection(report)
	if direction == 0 {
		return false
	}

	for i := 0; i < len(report)-1; i++ {
		difference := 0
		if direction == 1 {
			difference = report[i+1] - report[i]
		} else if direction == -1 {
			difference = report[i] - report[i+1]
		} else {
			panic("unexpected direction")
		}

		if difference < 1 || difference > 3 {
			return false
		}
	}

	return true
}

func levelsDirection(report []int) int {
	if report[0] == report[1] {
		return 0
	}

	if report[0] < report[1] {
		return 1
	}
	return -1
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(2, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	dataString := lib.GetDataString()

	result1 := solvePart1(dataString)
	fmt.Println(result1)

	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
