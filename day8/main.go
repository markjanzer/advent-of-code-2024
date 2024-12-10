package main

import (
	"advent-of-code-2024/lib"
)

const TestString string = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

/*
	Part 1 Notes

	Alright, we make lists of each of the frequency antenna locations,
	A separate list for each frequency.
	We create a map of antinodes (points) with booleans as values (unique list)
	For each list of points, we iterate through all of the unique combinations of two
	antenna, and pass a function that gets the two antinodes for the given points.
	For those two points, if they are inside the grid, add them to the map of antinodes
	Return the length of the map of antinodes

	Function to return antinodes: takes 2 points (pointA, pointB), returns 2 points
	Get the distance from pointA to pointB
	return (pointB + distance, pointA - distance)
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
	lib.AssertEqual(14, solvePart1(TestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	// dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
