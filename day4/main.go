package main

import (
	"advent-of-code-2024/lib"
	"fmt"
)

const TestString string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

/*
	Part 1 Notes

	See how many times XMAS appears in the input. All directions NESW and diagonal.

	I think there are two ways I could go about it. I could re-orient the grid all eight directions and then search for forward XMAS
	in each of the grids with regex. I think I could do this by rotating the grid 45 degrees

	The other way I could do this is find all of the X characters and then parse neighbor characters for an M
	If I find an M then I hold onto the direction [y,x] of the character, then proceed to search for the
	rest of the characters.

	I think the second approach is a little easier. Let's try it

*/

// const searchString string = "XMAS"

func solvePart1(input string) int {
	grid := lib.Grid{}
	grid.Create(input)
	grid.Print()

	total := 0

	for y := range grid {
		for x := range grid[y] {
			point := Point{X: x, Y: y, Grid: grid}
			total += XMASStringsFromPoint(point)
		}
	}

	return total
}

func XMASStringsFromPoint(point Point) int {
	total := 0
	searchString := "XMAS"

	if point.Char() == "X" {
		neighbors := point.ValidNeighbors()
		for _, neighbor := range neighbors {
			searchStringIndex := 1
			nextLetter := string(searchString[searchStringIndex])
			if neighbor.Char() == nextLetter {
				nextPoint := neighbor
				directionX, directionY := point.Direction(nextPoint)
				searchStringIndex++
				for searchStringIndex < len(searchString) {
					nextLetter = string(searchString[searchStringIndex])
					nextPoint = nextPoint.MoveDirection(directionX, directionY)
					if !nextPoint.IsInGrid() {
						break
					}
					if nextPoint.Char() != nextLetter {
						break
					}
					searchStringIndex++
				}
				// If we made it through the search string then we have found an XMAS
				if searchStringIndex == len(searchString) {
					total++
				}
			}
		}
	}
	return total
}

type Point struct {
	X    int
	Y    int
	Grid lib.Grid
}

func (p Point) NeighborPoints() []Point {
	return []Point{
		{p.X - 1, p.Y - 1, p.Grid},
		{p.X - 1, p.Y, p.Grid},
		{p.X - 1, p.Y + 1, p.Grid},
		{p.X, p.Y + 1, p.Grid},
		{p.X + 1, p.Y - 1, p.Grid},
		{p.X + 1, p.Y, p.Grid},
		{p.X + 1, p.Y + 1, p.Grid},
		{p.X, p.Y - 1, p.Grid},
	}
}

func (p Point) ValidNeighbors() []Point {
	neighbors := p.NeighborPoints()
	var validNeighbors []Point
	for _, neighbor := range neighbors {
		if neighbor.IsInGrid() {
			validNeighbors = append(validNeighbors, neighbor)
		}
	}
	return validNeighbors
}

func (p Point) Direction(newPoint Point) (x, y int) {
	return newPoint.X - p.X, newPoint.Y - p.Y
}

func (p Point) MoveDirection(x, y int) Point {
	return Point{p.X + x, p.Y + y, p.Grid}
}

func (p Point) IsInGrid() bool {
	return p.X >= 0 && p.X < len(p.Grid[0]) && p.Y >= 0 && p.Y < len(p.Grid)
}

func (p Point) Char() string {
	return string(p.Grid[p.Y][p.X])
}

/*
	Part 2 Notes

*/

func solvePart2(input string) int {
	return 0
}

func main() {
	lib.AssertEqual(18, solvePart1(TestString))
	// lib.AssertEqual(1, solvePart2(TestString))

	dataString := lib.GetDataString()

	result1 := solvePart1(dataString)
	fmt.Println(result1)

	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
