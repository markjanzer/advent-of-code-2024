package lib

type Point struct {
	X int
	Y int
}

func (p Point) Direction(newPoint Point) (x, y int) {
	return newPoint.X - p.X, newPoint.Y - p.Y
}

func (p Point) MoveDirection(x, y int) Point {
	return Point{p.X + x, p.Y + y}
}

func (p Point) IsInGrid(grid Grid) bool {
	return p.X >= 0 && p.X < len(grid[0]) && p.Y >= 0 && p.Y < len(grid)
}

func (p Point) Char(grid Grid) string {
	return string(grid[p.Y][p.X])
}
