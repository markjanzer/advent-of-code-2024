package lib

type Point struct {
	X    int
	Y    int
	Grid Grid
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
