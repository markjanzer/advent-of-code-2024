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

	Okay, I'm not sure what's wrong with that solution but there is a simpler approach to the problem.
	Right now we are iterating through the whole path, and on each step simulating what the path would be
	if there were an obstruction in front.
	Instead, we could run the whole path once keeping track of traveled points (as we did in part 1),
	and then we could iterate over those points, seeing if the path would be a loop if we added an
	obstruction on the point. We would also ignore the starting point. Let's attempt this.

	Ohhh wow the other answer didn't work because it wasn't properly simulating the scenarios, it was looking
	for a loop from the current position, not from the guard's starting position. What it failed to recognize
	is that if you put the new obstruction in fron the of the current path, that same obstruction could
	block an earlier path, meaning the guard would have never gotten to that position, and any loop found
	would be a false positive.
*/

func solvePart2(input string) int {
	guard, obstructions, xMax, yMax := scanGrid(input)
	originalGuardPoint := guard
	direction := north

	// Use visited from part one to get an array of the visited squares
	visited, loop := travel(guard, direction, xMax, yMax, obstructions, Point{})
	if loop {
		panic("Original path was loop")
	}

	visitedPoints := []Point{}
	for point := range visited {
		if point == originalGuardPoint {
			continue
		}
		visitedPoints = append(visitedPoints, point)
	}

	// Iterate over the visited point, adding to a total when loop is true.
	total := 0
	for _, newObstruction := range visitedPoints {
		_, loop := travel(guard, direction, xMax, yMax, obstructions, newObstruction)
		if loop {
			total += 1
		}
	}

	return total
}

func travel(guard Point, direction Direction, xMax, yMax int, obstructions map[Point]bool, newObstruction Point) (visited map[Point]map[Direction]bool, loop bool) {
	visited = make(map[Point]map[Direction]bool)

	for guard.IsInGrid(xMax, yMax) {
		if visited[guard] == nil {
			visited[guard] = make(map[Direction]bool)
		}
		if visited[guard][direction] {
			return visited, true
		}
		visited[guard][direction] = true

		nextPosition := Point{guard.X + direction.X, guard.Y + direction.Y}
		if obstructions[nextPosition] || nextPosition == newObstruction {
			direction = direction.TurnRight()
		} else {
			guard = nextPosition
		}
	}
	return visited, false
}

// func solvePart2(input string) int {
// 	guard, obstructions, xMax, yMax := scanGrid(input)
// 	originalGuardLocation := guard
// 	visited := make(map[Point]map[string]bool)
// 	possibleObstructions := make(map[Point]bool)
// 	direction := north

// 	for guard.IsInGrid(xMax, yMax) {
// 		if visited[guard] == nil {
// 			visited[guard] = make(map[string]bool)
// 		}
// 		visited[guard][direction.Name] = true

// 		nextPosition := Point{guard.X + direction.X, guard.Y + direction.Y}
// 		if obstructions[nextPosition] {
// 			direction = direction.TurnRight()
// 			continue
// 		}

// 		if nextPosition.IsInGrid(xMax, yMax) && pathHasLoop(guard, direction, obstructions, xMax, yMax, nextPosition) {
// 			possibleObstructions[nextPosition] = true
// 		}

// 		guard = nextPosition
// 	}

// 	result := len(possibleObstructions)
// 	if possibleObstructions[originalGuardLocation] {
// 		result -= 1
// 	}
// 	return result
// }

// func pathHasLoop(guard Point, direction Direction, obstructions map[Point]bool, xMax, yMax int, newObstruction Point) bool {
// 	visited := make(map[Point]map[string]bool)
// 	visited[guard] = make(map[string]bool)
// 	visited[guard][direction.Name] = true

// 	for guard.IsInGrid(xMax, yMax) {
// 		nextPosition := Point{guard.X + direction.X, guard.Y + direction.Y}
// 		if obstructions[nextPosition] || nextPosition == newObstruction {
// 			direction = direction.TurnRight()
// 			continue
// 		}

// 		guard = nextPosition
// 		if visited[guard] == nil {
// 			visited[guard] = make(map[string]bool)
// 		}
// 		// If visited before then we're in a loop
// 		if visited[guard][direction.Name] {
// 			return true
// 		}
// 		visited[guard][direction.Name] = true
// 	}

// 	return false
// }

func main() {
	lib.AssertEqual(41, solvePart1(TestString))
	lib.AssertEqual(6, solvePart2(TestString))

	dataString := lib.GetDataString()

	result1 := solvePart1(dataString)
	fmt.Println(result1)

	result2 := solvePart2(dataString)
	fmt.Println(result2)
}
