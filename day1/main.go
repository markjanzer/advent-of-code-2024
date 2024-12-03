package main

import (
	"advent-of-code-2024/lib"
	"sort"
	"strconv"
	"strings"
)

// import (
// 	"advent-of-code-2024/lib"
// )

const SmallTestString string = `3   4
4   3
2   5
1   3
3   9
3   3`

const TestString string = ``

const DataFile string = "data.txt"

/*
	Part 1 Notes

	Read each line, put the first number in one list, and the second in another
	Order both of the lists from smallest to largest

	For each index, get the difference between the value of each list at that index,
	and add it to a total sum

	Return the total sum
*/

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")

	list1 := []int{}
	list2 := []int{}

	for _, line := range lines {
		fields := strings.Fields(line)
		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	result := 0

	for i, _ := range list1 {
		result += list2[i] - list1[i]
	}

	return result
}

/*
	Part 2 Notes

*/

// func solvePart2(input string) int {
// 	return 0
// }

func main() {
	lib.AssertEqual(11, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	// lib.AssertEqual(1, solvePart1(SmallTestString))
	// lib.AssertEqual(1, solvePart2(SmallTestString))

	// dataString := lib.GetDataString(DataFile)
	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// dataString := lib.GetDataString(DataFile)
	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
