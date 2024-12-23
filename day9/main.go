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

	I didn't read instructions carefully enough, and I overengineered the solution XD
	I'm doing this all backwards. Right now I am iterating forwards and upon finding a
	space searching each file from the back.
	I need to be searching each file from the back and for each of those,
	iterating from the front to try to find a spot.

	Okay the only error happening now is that when I'm moving a file I'm not replacing
	the space with an empty file. This is fine early on, but later it creates issues.

	I think this might have been signficantly easier without the file abstraction.
*/

func solvePart2(input string) int {
	diskBlocks := createDiskBlocks2(input)
	compactedDisk := collapseDisk(diskBlocks)
	compactedDisk = compactedDisk.removeEmptyBlocks()
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
	maxId := disk.maxId()

	for id := maxId; id >= 0; id-- {
		disk = disk.moveFileToAvailableSpace(id)
		disk = disk.combineContiguousSpaceBlocks()
	}
	disk = disk.removeEmptyBlocks()

	return disk
}

func (disk Disk) maxId() int {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i].Empty() {
			continue
		}
		return disk[i].ID
	}
	panic(fmt.Sprintf("Disk has no files with an ID: %v", disk))
}

func (disk Disk) fileIndex(id int) int {
	for i, file := range disk {
		if file.ID == id {
			return i
		}
	}
	panic(fmt.Sprintf("Disk does not have file with id: %d", id))
}

func (disk Disk) moveFileToAvailableSpace(id int) Disk {
	idIndex := disk.fileIndex(id)
	fileToMove := disk[idIndex]

	for i := 0; i < idIndex; i++ {
		if disk[i].Empty() && disk[i].Size >= fileToMove.Size {
			disk[i].Size = disk[i].Size - fileToMove.Size
			disk = copyElement(disk, idIndex, i)
			// Ensure that the old space is empty
			disk[idIndex+1].ID = -1
			break
		}
	}
	return disk
}

func copyElement(disk Disk, from int, to int) Disk {
	element := disk[from]
	disk = addElement(disk, to, element)
	return disk
}

func (d Disk) combineContiguousSpaceBlocks() Disk {
	for i := 1; i < len(d); i++ {
		if d[i].Empty() && d[i-1].Empty() {
			d[i-1].Size += d[i].Size
			d = removeElement(d, i)
			i--
		}
	}
	return d
}

func (d Disk) removeEmptyBlocks() Disk {
	for i := 0; i < len(d); i++ {
		if d[i].Size == 0 {
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

func main() {
	lib.AssertEqual(1928, solvePart1(TestString))
	lib.AssertEqual(2858, solvePart2(TestString))

	dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	result2 := solvePart2(dataString)
	fmt.Println(result2)
}
