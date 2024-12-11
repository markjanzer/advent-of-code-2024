package main

import (
	"advent-of-code-2024/lib"
	"fmt"
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
		if char == "." {
			continue
		}
		num, _ := strconv.Atoi(char)
		total += num * i
	}
	return total
}

/*
	Part 2 Notes

	We need to write a different function to collapse bocks. Instead of moving individual files
	We need to move blocks of files.

	One thing we need to do is have the checksum be able to take "." and read it as 0

	Another this is we probably want to store the data a little differently.
	We could have structs that have an id and a length, and the disk could be a
	slice of these structs.

*/

func solvePart2(input string) int {
	diskBlocks := createDiskBlocks2(input)
	compactedDisk := collapseDisk(diskBlocks)
	compactedDisk = compactedDisk.RemoveEmptyBlocks()
	compactedDiskAsSlice := compactedDisk.ToSlice()
	checksum := checksum(compactedDiskAsSlice)

	return checksum
}

type File struct {
	ID   int
	Size int
}

type Disk []File

func (f File) Empty() bool {
	return f.ID == -1
}

func (d Disk) Print() {
	slice := d.ToSlice()
	fmt.Println(slice)
}

// 12345 -> 0..111....22222
func createDiskBlocks2(input string) Disk {
	disk := Disk{}

	id := 0
	for i, char := range input {
		size := int(char - '0')
		fileId := -1
		if i%2 == 0 {
			fileId = id
			id++
		}
		file := File{ID: fileId, Size: size}
		disk = append(disk, file)
	}

	return disk
}

func collapseDisk(disk Disk) Disk {
	// Iterate over the current Disk
	// If there is an empty space (id == -1) then we start a reversed search on the disk
	// When we find a file that is <= the space then we move that file in the disk ahead
	// of the current space, and replace the space by the difference.

	// Hopefully at this point we're still working on that space, and we try to reverse search
	// again.
	// If that doesn't work then we keep on iterating.

	// Then we might need a function that takes a disk object and turns it into an array
	// So that the checksum can happen.

	// Ran into some errors with indices.
	// Instead of removing the space and potentially ending up with index errors,
	// we can keep the block and set the space to 0.
	// However we also need to combine contiguous space blocks.

	i := 0
	for i < len(disk) {
		disk = disk.CombineContiguousSpaceBlocks()
		disk.Print()

		empty := disk[i].Empty()
		if !empty {
			i++
			continue
		}
		availableSpace := disk[i].Size

		j := len(disk) - 1
		// Should j > i or j > 0?
		for j > i {
			if disk[j].Empty() {
				j--
				continue
			}
			if disk[j].Size > availableSpace {
				j--
				continue
			}

			// If we're removing the space or inserting an element,
			// lets try to stay in the same index to ensure that we don't
			// miss one of the blocks

			disk[i].Size = availableSpace - disk[j].Size
			fmt.Printf("Moving from %d to %d (len: %d)\n", j, i, len(disk))
			disk = moveElement(disk, j, i)
			j = -1
		}

		i++
	}
	return disk
}

func (d Disk) CombineContiguousSpaceBlocks() Disk {
	for i := 1; i < len(d); i++ {
		if d[i].Empty() && d[i-1].Empty() {
			d[i-1].Size += d[i].Size
			d = removeElement(d, i)
			i--
		}
	}
	return d
}

func (d Disk) RemoveEmptyBlocks() Disk {
	for i := 0; i < len(d); i++ {
		if d[i].Empty() {
			d = removeElement(d, i)
			i--
		}
	}
	return d
}

func (d Disk) ToSlice() []string {
	var result []string
	for _, file := range d {
		char := "."
		if file.ID != -1 {
			char = strconv.Itoa(file.ID)
		}
		for i := 0; i < file.Size; i++ {
			result = append(result, char)
		}
	}
	return result
}

func removeElement[T any](slice []T, i int) []T {
	return append(slice[:i], slice[i+1:]...)
}

func addElement[T any](slice []T, i int, element T) []T {
	// Make room for the new element by appending a zero value
	slice = append(slice, *new(T))
	// Shift elements to the right
	copy(slice[i+1:], slice[i:])
	// Insert the new element
	slice[i] = element
	return slice
}

func moveElement[T any](slice []T, from int, to int) []T {
	if from < 0 || from >= len(slice) || to < 0 || to >= len(slice) {
		panic(fmt.Sprintf("Invalid indices: from=%d, to=%d, len=%d", from, to, len(slice)))
	}

	element := slice[from]
	sliceWithoutElement := removeElement(slice, from)
	return addElement(sliceWithoutElement, to, element)
}

func main() {
	lib.AssertEqual(1928, solvePart1(TestString))
	lib.AssertEqual(2858, solvePart2(TestString))

	// dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
