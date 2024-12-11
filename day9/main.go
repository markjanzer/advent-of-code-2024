package main

import (
	"advent-of-code-2024/lib"
	"strconv"
)

const TestString string = `2333133121414131402`

/*
	Part 1 Notes

	Okay we could go with their approach.
	12345 -> 0..111....22222
	Create a new string
	Have an id iterator at 0
	For each number we either add numbers or space alternating
	When we add the number we append X amount of the id (and increment id)
	When we add the space we append X amount of .

	0..111....22222 -> 022111222......
	For this part, let's create a new string
	While the old string has any characters we
	Pop the first element from the array
		If it is a number we append it to the new array
		If it is a . then we discard it, and we pop the Last element from the array
			If the last element was a . then we try again
			If the last element was not a . then we append that to the new array

	To generate the checksum have total. Multiply the id with the position of the number
	and add it to the total

	Ahh answer is too high. It works with test data but not real data.

	Part of my answer has emojis in it, that can't be good.
	I think what's happening is that numbers are getting to double digits and then
	everything is breaking down.
*/

func solvePart1(input string) int {
	diskBlocks := createDiskBlocks(input)
	compactedDisk := collapseBlocks(diskBlocks)
	checksum := checksum(compactedDisk)

	return checksum
}

// 12345 -> 0..111....22222
func createDiskBlocks(input string) []string {
	var result []string
	id := 0
	for i, char := range input {
		num := int(char - '0')
		appendChar := "."
		if i%2 == 0 {
			appendChar = strconv.Itoa(id)
			id++
		}
		for j := 0; j < num; j++ {
			result = append(result, appendChar)
		}
	}
	return result
}

func collapseBlocks(diskBlocks []string) []string {
	chars := diskBlocks
	result := make([]string, 0, len(chars))

	for len(chars) > 0 {
		first := chars[0]
		chars = chars[1:]

		if first != "." {
			result = append(result, first)
		} else {
			last := "."
			for len(chars) > 0 && last == "." {
				lastIndex := len(chars) - 1
				last = chars[lastIndex]
				chars = chars[:lastIndex]

				if last != "." {
					result = append(result, last)
				}
			}
		}
	}

	return result
}

func checksum(disk []string) int {
	total := 0
	for i, char := range disk {
		num, _ := strconv.Atoi(char)
		total += num * i
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
	lib.AssertEqual(1928, solvePart1(TestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	// dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
