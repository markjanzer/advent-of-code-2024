package lib

import (
	"fmt"
	"strings"
)

type Grid [][]byte

func (g *Grid) Create(input string) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		*g = append(*g, []byte(line))
	}
}

func (g Grid) ToString() string {
	var lines []string
	for _, line := range g {
		lines = append(lines, string(line))
	}

	return strings.Join(lines, "\n")
}

func (g Grid) HasPoint(x, y int) bool {
	return IndexInSlice(y, g) && IndexInSlice(x, g[y])
}

func (g Grid) Print() {
	for y := range g {
		fmt.Println(string(g[y]))
	}
}
