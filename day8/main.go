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

	antenna := makeAntenna(grid)

	antinodes := make(map[lib.Point]bool)
	for _, antennaList := range antenna {
		combinations := allCombinations(antennaList)
		for _, points := range combinations {
			potentialAntinodes := antinodesFrom(points[0], points[1])
			for _, node := range potentialAntinodes {
				if node.IsInGrid(grid) {
					antinodes[node] = true
				}
			}
		}
	}

	return len(antinodes)
}

func makeAntenna(grid lib.Grid) map[string][]lib.Point {
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
	return antenna
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

func antinodesFrom(p1, p2 lib.Point) []lib.Point {
	xDirection, yDirection := p1.Direction(p2)
	a1 := lib.Point{X: p2.X + xDirection, Y: p2.Y + yDirection}
	a2 := lib.Point{X: p1.X - xDirection, Y: p1.Y - yDirection}
	return []lib.Point{a1, a2}
}

/*
	Part 2 Notes

	Okay this is going to be very similar to part 1, but instead of finding the points
	one frequency away from either antinode, we're going to keep finding points of that same
	distance and direction away from each other until we go off the grid.

	I think that this will just require a change to the antinodesFrom function.
	Instead of returning one antinode for each direction, each will be a for loop that adds to the result
	until it reaches a point out of the grid.

	Ah also didn't read instructions well enough, antenna are also antinodes!
*/

func solvePart2(input string) int {
	grid := lib.Grid{}
	grid.Create(input)

	antenna := makeAntenna(grid)

	antinodes := make(map[lib.Point]bool)
	for _, antennaList := range antenna {
		combinations := allCombinations(antennaList)
		for _, points := range combinations {
			// Points are also now antinodes
			for _, point := range points {
				antinodes[point] = true
			}
			a := harmonicAntinodes(points[0], points[1], grid)
			for _, node := range a {
				antinodes[node] = true
			}
		}
	}

	return len(antinodes)
}

func harmonicAntinodes(p1, p2 lib.Point, grid lib.Grid) []lib.Point {
	xDirection, yDirection := p1.Direction(p2)
	results := []lib.Point{}

	nextPoint := lib.Point{X: p2.X + xDirection, Y: p2.Y + yDirection}
	for nextPoint.IsInGrid(grid) {
		results = append(results, nextPoint)
		nextPoint = lib.Point{X: nextPoint.X + xDirection, Y: nextPoint.Y + yDirection}
	}

	nextPoint = lib.Point{X: p1.X - xDirection, Y: p1.Y - yDirection}
	for nextPoint.IsInGrid(grid) {
		results = append(results, nextPoint)
		nextPoint = lib.Point{X: nextPoint.X - xDirection, Y: nextPoint.Y - yDirection}
	}

	return results
}

func main() {
	lib.AssertEqual(14, solvePart1(TestString))
	lib.AssertEqual(34, solvePart2(TestString))

	dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	result2 := solvePart2(dataString)
	fmt.Println(result2)
}
