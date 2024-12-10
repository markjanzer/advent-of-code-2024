package main

import (
	"advent-of-code-2024/lib"
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

	total := 0

	for y := range grid {
		for x := range grid[y] {
			point := lib.Point{X: x, Y: y}
			total += XMASStringsFromPoint(point, grid)
		}
	}

	return total
}

func XMASStringsFromPoint(point lib.Point, grid lib.Grid) int {
	total := 0
	searchString := "XMAS"

	if point.Char(grid) == "X" {
		neighbors := ValidNeighbors(point, grid)
		for _, neighbor := range neighbors {
			searchStringIndex := 1
			nextLetter := string(searchString[searchStringIndex])
			if neighbor.Char(grid) == nextLetter {
				nextPoint := neighbor
				directionX, directionY := point.Direction(nextPoint)
				searchStringIndex++
				for searchStringIndex < len(searchString) {
					nextLetter = string(searchString[searchStringIndex])
					nextPoint = nextPoint.MoveDirection(directionX, directionY)
					if !nextPoint.IsInGrid(grid) {
						break
					}
					if nextPoint.Char(grid) != nextLetter {
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

func NeighborPoints(p lib.Point) []lib.Point {
	return []lib.Point{
		{X: p.X - 1, Y: p.Y - 1},
		{X: p.X - 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y + 1},
		{X: p.X, Y: p.Y + 1},
		{X: p.X + 1, Y: p.Y - 1},
		{X: p.X + 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
	}
}

func ValidNeighbors(p lib.Point, grid lib.Grid) []lib.Point {
	neighbors := NeighborPoints(p)
	var validNeighbors []lib.Point
	for _, neighbor := range neighbors {
		if neighbor.IsInGrid(grid) {
			validNeighbors = append(validNeighbors, neighbor)
		}
	}
	return validNeighbors
}

/*
	Part 2 Notes

	Now we're looking for MAS strings in the shape of Xs.

	So I think we look for any A character, then we get pairs of neighbors:
	top left + bottom right
	top right + bottom left

	If both those pairs of characters are the collection of [M,S] then we add to the total.
*/

func solvePart2(input string) int {
	grid := lib.Grid{}
	grid.Create(input)

	total := 0

	for y := range grid {
		for x := range grid[y] {
			point := lib.Point{X: x, Y: y}
			if XShapedMASFromPoint(point, grid) {
				total++
			}
		}
	}

	return total
}

func XShapedMASFromPoint(point lib.Point, grid lib.Grid) bool {
	if point.Char(grid) != "A" {
		return false
	}

	topLeft := point.MoveDirection(-1, -1)
	bottomRight := point.MoveDirection(1, 1)

	topRight := point.MoveDirection(1, -1)
	bottomLeft := point.MoveDirection(-1, 1)

	allCrossPoints := []lib.Point{topLeft, topRight, bottomLeft, bottomRight}

	for _, crossPoint := range allCrossPoints {
		if !crossPoint.IsInGrid(grid) {
			return false
		}
	}

	cross1 := []lib.Point{topLeft, bottomRight}
	cross2 := []lib.Point{topRight, bottomLeft}

	return validCross(cross1, grid) && validCross(cross2, grid)
}

func validCross(cross []lib.Point, grid lib.Grid) bool {
	return cross[0].Char(grid) == "M" && cross[1].Char(grid) == "S" ||
		cross[0].Char(grid) == "S" && cross[1].Char(grid) == "M"
}

func main() {
	lib.AssertEqual(18, solvePart1(TestString))
	lib.AssertEqual(9, solvePart2(TestString))

	// dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	// result2 := solvePart2(dataString)
	// fmt.Println(result2)
}
