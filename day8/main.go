package main

import (
	"advent-of-code-2024/lib"
	"fmt"
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
	grid := lib.Grid{}
	grid.Create(input)

	antenna := make(map[string][]lib.Point)
	for y, row := range grid {
		for x, b := range row {
			char := string(b)
			if char == "." {
				continue
			}
			if antenna[char] == nil {
				antenna[char] = []lib.Point{}
			}
			antenna[char] = append(antenna[char], lib.Point{X: x, Y: y})
		}
	}

	antinodes := make(map[lib.Point]bool)
	for _, antennaList := range antenna {
		combinations := allCombinations(antennaList)
		for _, points := range combinations {
			a1, a2 := antinodesFrom(points[0], points[1])

			if a1.IsInGrid(grid) {
				antinodes[a1] = true
			}
			if a2.IsInGrid(grid) {
				antinodes[a2] = true
			}
		}
	}

	return len(antinodes)
}

func allCombinations[T any](list []T) [][]T {
	result := [][]T{}
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			result = append(result, []T{list[i], list[j]})
		}
	}
	return result
}

func antinodesFrom(p1, p2 lib.Point) (a1, a2 lib.Point) {
	xDirection, yDirection := p1.Direction(p2)
	a1 = lib.Point{X: p2.X + xDirection, Y: p2.Y + yDirection}
	a2 = lib.Point{X: p1.X - xDirection, Y: p1.Y - yDirection}
	return a1, a2
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

	dataString := lib.GetDataString()

	result1 := solvePart1(dataString)
	fmt.Println(result1)

	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
