package day5

import (
	"fmt"
)

type coordinate struct {
	x, y int
}

func newCoordinate(x int, y int) coordinate {
	return coordinate{x: x, y: y}
}

func (c coordinate) ToString() string {
	return fmt.Sprintf("%v,%v", c.x, c.y)
}

type vent struct {
	origin      coordinate
	destination coordinate
}

func newVent(x1 int, y1 int, x2 int, y2 int) *vent {
	return &vent{origin: newCoordinate(x1, y1), destination: newCoordinate(x2, y2)}
}

func (v vent) OrthogonalLine() []coordinate {
	var line []coordinate
	if v.origin.x == v.destination.x || v.origin.y == v.destination.y {
		if v.origin.x == v.destination.x {
			for i := min(v.origin.y, v.destination.y); i <= max(v.origin.y, v.destination.y); i++ {
				line = append(line, newCoordinate(v.origin.x, i))
			}
		} else {
			for i := min(v.origin.x, v.destination.x); i <= max(v.origin.x, v.destination.x); i++ {
				line = append(line, newCoordinate(i, v.origin.y))
			}
		}
	}

	return line
}

func (v vent) Line() []coordinate {
	current := v.origin
	// Initialize with origin
	line := []coordinate{current}

	for current != v.destination {
		var x, y int
		if current.x == v.destination.x {
			x = current.x
		} else if current.x < v.destination.x {
			x = current.x + 1
		} else {
			x = current.x - 1
		}
		if current.y == v.destination.y {
			y = current.y
		} else if current.y < v.destination.y {
			y = current.y + 1
		} else {
			y = current.y - 1
		}

		current = newCoordinate(x, y)
		line = append(line, current)
	}

	return line
}

func (v vent) PrintVent() {
	fmt.Println(v.origin.x, v.origin.y, "->", v.destination.x, v.destination.y)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
