package main

import (
	"advent-of-code-2024/lib"
	"fmt"
	"strings"
)

const TestString string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

/*
	Part 1 Notes

	Get the coordinates of the guard
	Get the coordinates of the obstructions, and store them in a hash
	Get the dimensions of the grid
	Get the four directions in order (turning right), NESW
	Make a map to store visited coordinates
	Add the original guard location to visited

	Iterate over movement, starting with North, and moving to the
	next direction whenever the next step would be obstructed.
	Add each location to visited.
	Break when the guard leaves the grid
	Return the length of visited
*/

type Point struct {
	X int
	Y int
}

type Direction struct {
	X    int
	Y    int
	Name string
}

var north = Direction{X: 0, Y: -1, Name: "N"}
var east = Direction{X: 1, Y: 0, Name: "E"}
var south = Direction{X: 0, Y: 1, Name: "S"}
var west = Direction{X: -1, Y: 0, Name: "W"}

func (d Direction) TurnRight() Direction {
	switch d.Name {
	case "N":
		return east
	case "E":
		return south
	case "S":
		return west
	case "W":
		return north
	default:
		return d
	}
}

func (p Point) IsInGrid(xMax, yMax int) bool {
	return p.X >= 0 && p.X < xMax && p.Y >= 0 && p.Y < yMax
}

func scanGrid(input string) (guard Point, obstructions map[Point]bool, xMax, yMax int) {
	obstructions = make(map[Point]bool)
	rows := strings.Split(input, "\n")
	yMax = len(rows)
	xMax = len(rows[0])

	for y, row := range rows {
		for x, char := range row {
			if char == '^' {
				guard = Point{x, y}
			} else if char == '#' {
				ob := Point{x, y}
				obstructions[ob] = true
			}
		}
	}

	return guard, obstructions, xMax, yMax
}

func solvePart1(input string) int {
	guard, obstructions, xMax, yMax := scanGrid(input)
	visited := make(map[Point]bool)
	visited[guard] = true
	direction := north

	for guard.IsInGrid(xMax, yMax) {
		visited[guard] = true
		nextPosition := Point{guard.X + direction.X, guard.Y + direction.Y}
		if obstructions[nextPosition] {
			direction = direction.TurnRight()
		} else {
			guard = nextPosition
		}
	}

	return len(visited)
}

/*
	Part 2 Notes

	Create a newObstruction map
	So instead of just putting true as visited, we'll put the direction name
	If turning right at any point would be a duplicate, then we save that
	space as a potential newObstruction

	Oh the problem is that there are times when a loop would form but it still
	is getting a unique point/direction. That unique point would eventually meet
	a non-unique point-direction.

	We don't need a recursive solution, because we aren't going more than 2 levels deep.
	We need one navigation of the current path, then we need to turn right wherever possible
	and generate those paths to see if they ever loop.

	Answer is too high, meaning I've included one that can't be included. Oh I realized that I can't put the
	barrier where to guard is currently stationed!
	Ah that wasn't the issue.
	Theoretically could it be putting the obstructions outside of the grid?
*/

func solvePart2(input string) int {
	guard, obstructions, xMax, yMax := scanGrid(input)
	originalGuardLocation := guard
	newObstructions := make(map[Point]bool)
	visited := make(map[Point]map[string]bool)
	direction := north
	visited[guard] = make(map[string]bool)
	visited[guard][direction.Name] = true

	for guard.IsInGrid(xMax, yMax) {
		if visited[guard] == nil {
			visited[guard] = make(map[string]bool)
		}
		visited[guard][direction.Name] = true

		nextPosition := Point{guard.X + direction.X, guard.Y + direction.Y}
		if obstructions[nextPosition] {
			direction = direction.TurnRight()
			continue
		}

		// Check if adding an obstruction would cause a loop
		modifiedObstructions := copyMap(obstructions)
		modifiedObstructions[nextPosition] = true
		if pathHasLoop(guard, modifiedObstructions, xMax, yMax, direction) {
			newObstructions[nextPosition] = true
		}

		guard = nextPosition
	}

	result := len(newObstructions)
	if newObstructions[originalGuardLocation] {
		result -= 1
	}
	return result
}

func pathHasLoop(guard Point, obstructions map[Point]bool, xMax, yMax int, direction Direction) bool {
	visited := make(map[Point]map[string]bool)
	visited[guard] = make(map[string]bool)
	visited[guard][direction.Name] = true

	for guard.IsInGrid(xMax, yMax) {
		nextPosition := Point{guard.X + direction.X, guard.Y + direction.Y}
		if obstructions[nextPosition] {
			direction = direction.TurnRight()
			continue
		}

		guard = nextPosition
		if visited[guard] == nil {
			visited[guard] = make(map[string]bool)
		}
		// If visited before then we're in a loop
		if visited[guard][direction.Name] {
			return true
		}
		visited[guard][direction.Name] = true
	}

	return false
}

func copyMap(original map[Point]bool) map[Point]bool {
	newMap := make(map[Point]bool)
	for key, value := range original {
		newMap[key] = value
	}
	return newMap
}

func main() {
	lib.AssertEqual(41, solvePart1(TestString))
	lib.AssertEqual(6, solvePart2(TestString))

	dataString := lib.GetDataString()

	// result1 := solvePart1(dataString)
	// fmt.Println(result1)

	result2 := solvePart2(dataString)
	fmt.Println(result2)
}
