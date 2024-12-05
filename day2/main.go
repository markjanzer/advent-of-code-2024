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

type Report []int

func (r Report) direction() int {
	if r[0] == r[1] {
		return 0
	}

	if r[0] < r[1] {
		return 1
	}
	return -1
}

func (r Report) IsSafe() bool {
	direction := r.direction()

	if direction == 0 {
		return false
	}

	for i := 0; i < len(r)-1; i++ {
		difference := 0
		if direction == 1 {
			difference = r[i+1] - r[i]
		} else if direction == -1 {
			difference = r[i] - r[i+1]
		} else {
			panic("unexpected direction")
		}

		if difference < 1 || difference > 3 {
			return false
		}
	}

	return true
}

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")

	safeReports := 0

	for _, line := range lines {
		fields := strings.Fields(line)
		report := Report{}
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			report = append(report, num)
		}
		if report.IsSafe() {
			safeReports++
		}
	}
	return safeReports
}

/*
	Part 2 Notes

	Similar to above, but if there is an error, see if removing one of the numbers makes it safe.
*/

// Returns all reports with one number removed
func (r Report) reportVariants() []Report {
	reports := []Report{}

	for i := 0; i < len(r); i++ {
		report := Report{}
		report = append(report, r[:i]...)
		report = append(report, r[i+1:]...)
		reports = append(reports, report)
	}
	return reports
}

func (r Report) HasOneSafeVariant() bool {
	variants := r.reportVariants()
	for _, variant := range variants {
		if variant.IsSafe() {
			return true
		}
	}
	return false
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\n")

	safeReports := 0

	for _, line := range lines {
		fields := strings.Fields(line)
		report := Report{}
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			report = append(report, num)
		}
		if report.IsSafe() {
			safeReports++
		} else if report.HasOneSafeVariant() {
			safeReports++
		}
	}

	return safeReports
}

func main() {
	lib.AssertEqual(2, solvePart1(SmallTestString))
	lib.AssertEqual(4, solvePart2(SmallTestString))

	dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	result2 := solvePart2(dataString)
	fmt.Println(result2)
}
