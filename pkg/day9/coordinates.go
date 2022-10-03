package day9

type Coordinate struct {
	x, y int8
}

func NewCoordinate(x int8, y int8) *Coordinate {
	return &Coordinate{x: x, y: y}
}

// Adjacent Only includes the up, down, left, and right adjacent points
func (c Coordinate) Adjacent(width, height int8) []Coordinate {
	if c.x == 0 && c.y == 0 {
		// Top left corner
		return []Coordinate{
			*NewCoordinate(0, 1),
			*NewCoordinate(1, 0),
		}
	} else if c.x == 0 && c.y == height-1 {
		// Bottom left corner
		return []Coordinate{
			*NewCoordinate(0, height-2),
			*NewCoordinate(1, height-1),
		}
	} else if c.x == width-1 && c.y == height-1 {
		// Bottom right corner
		return []Coordinate{
			*NewCoordinate(width-2, height-1),
			*NewCoordinate(width-1, height-2),
		}
	} else if c.x == width-1 && c.y == 0 {
		// Top right corner
		return []Coordinate{
			*NewCoordinate(width-2, 0),
			*NewCoordinate(width-1, 1),
		}
	} else if c.y == 0 && c.x > 0 && c.x < width-1 {
		// Top row
		return []Coordinate{
			*NewCoordinate(c.x-1, c.y),
			*NewCoordinate(c.x, c.y+1),
			*NewCoordinate(c.x+1, c.y),
		}
	} else if c.x == 0 && c.y > 0 && c.y < height-1 {
		// Leftmost column
		return []Coordinate{
			*NewCoordinate(c.x, c.y-1),
			*NewCoordinate(c.x+1, c.y),
			*NewCoordinate(c.x, c.y+1),
		}
	} else if c.y == height-1 && c.x > 0 && c.x < width-1 {
		// Bottom row
		return []Coordinate{
			*NewCoordinate(c.x-1, c.y),
			*NewCoordinate(c.x, c.y-1),
			*NewCoordinate(c.x+1, c.y),
		}
	} else if c.x == width-1 && c.y > 0 && c.y < height-1 {
		// Rightmost column
		return []Coordinate{
			*NewCoordinate(c.x, c.y-1),
			*NewCoordinate(c.x-1, c.y),
			*NewCoordinate(c.x, c.y+1),
		}
	} else {
		// Middle of matrix
		return []Coordinate{
			*NewCoordinate(c.x-1, c.y),
			*NewCoordinate(c.x, c.y-1),
			*NewCoordinate(c.x+1, c.y),
			*NewCoordinate(c.x, c.y+1),
		}
	}
}

// DiagonalAdjacent Includes all adjacent points, including diagonals
func (c Coordinate) DiagonalAdjacent(width, height int8) []Coordinate {
	if c.x == 0 && c.y == 0 {
		// Top left corner
		return []Coordinate{
			*NewCoordinate(0, 1),
			*NewCoordinate(1, 0),
			*NewCoordinate(1, 1),
		}
	} else if c.x == 0 && c.y == height-1 {
		// Bottom left corner
		return []Coordinate{
			*NewCoordinate(0, height-2),
			*NewCoordinate(1, height-1),
			*NewCoordinate(1, height-2),
		}
	} else if c.x == width-1 && c.y == height-1 {
		// Bottom right corner
		return []Coordinate{
			*NewCoordinate(width-2, height-1),
			*NewCoordinate(width-1, height-2),
			*NewCoordinate(width-2, height-2),
		}
	} else if c.x == width-1 && c.y == 0 {
		// Top right corner
		return []Coordinate{
			*NewCoordinate(width-2, 0),
			*NewCoordinate(width-1, 1),
			*NewCoordinate(width-2, 1),
		}
	} else if c.y == 0 && c.x > 0 && c.x < width-1 {
		// Top row
		return []Coordinate{
			*NewCoordinate(c.x-1, c.y),
			*NewCoordinate(c.x, c.y+1),
			*NewCoordinate(c.x+1, c.y),
			*NewCoordinate(c.x-1, c.y+1),
			*NewCoordinate(c.x+1, c.y+1),
		}
	} else if c.x == 0 && c.y > 0 && c.y < height-1 {
		// Leftmost column
		return []Coordinate{
			*NewCoordinate(c.x, c.y-1),
			*NewCoordinate(c.x+1, c.y),
			*NewCoordinate(c.x, c.y+1),
			*NewCoordinate(c.x+1, c.y-1),
			*NewCoordinate(c.x+1, c.y+1),
		}
	} else if c.y == height-1 && c.x > 0 && c.x < width-1 {
		// Bottom row
		return []Coordinate{
			*NewCoordinate(c.x-1, c.y),
			*NewCoordinate(c.x, c.y-1),
			*NewCoordinate(c.x+1, c.y),
			*NewCoordinate(c.x-1, c.y-1),
			*NewCoordinate(c.x+1, c.y-1),
		}
	} else if c.x == width-1 && c.y > 0 && c.y < height-1 {
		// Rightmost column
		return []Coordinate{
			*NewCoordinate(c.x, c.y-1),
			*NewCoordinate(c.x-1, c.y),
			*NewCoordinate(c.x, c.y+1),
			*NewCoordinate(c.x-1, c.y-1),
			*NewCoordinate(c.x-1, c.y+1),
		}
	} else {
		// Middle of matrix
		return []Coordinate{
			*NewCoordinate(c.x-1, c.y),
			*NewCoordinate(c.x+1, c.y),
			*NewCoordinate(c.x, c.y-1),
			*NewCoordinate(c.x, c.y+1),
			*NewCoordinate(c.x-1, c.y-1),
			*NewCoordinate(c.x-1, c.y+1),
			*NewCoordinate(c.x+1, c.y-1),
			*NewCoordinate(c.x+1, c.y+1),
		}
	}
}
