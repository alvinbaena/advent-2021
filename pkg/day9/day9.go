package day9

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func loadInput() map[coordinate]int8 {
	file, err := os.Open("pkg/day9/input.txt")
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	coords := make(map[coordinate]int8)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		for x, p := range split {
			val, err := strconv.ParseInt(p, 10, 8)
			if err != nil {
				panic(err)
			}

			coords[*newCoordinate(int8(x), int8(y))] = int8(val)
		}
		y++
	}

	return coords
}

type coordinate struct {
	x, y int8
}

func newCoordinate(x int8, y int8) *coordinate {
	return &coordinate{x: x, y: y}
}

func (c coordinate) adjacent(width, height int8) []coordinate {
	if c.x == 0 && c.y == 0 {
		// Top left corner
		return []coordinate{*newCoordinate(0, 1), *newCoordinate(1, 0)}
	} else if c.x == 0 && c.y == height-1 {
		// Bottom left corner
		return []coordinate{*newCoordinate(0, height-2), *newCoordinate(1, height-1)}
	} else if c.x == width-1 && c.y == height-1 {
		// Bottom right corner
		return []coordinate{*newCoordinate(width-2, height-1), *newCoordinate(width-1, height-2)}
	} else if c.x == width-1 && c.y == 0 {
		// Top right corner
		return []coordinate{*newCoordinate(width-2, 0), *newCoordinate(width-1, 1)}
	} else if c.y == 0 && c.x > 0 && c.x < width-1 {
		// Top row
		return []coordinate{*newCoordinate(c.x-1, c.y), *newCoordinate(c.x, c.y+1), *newCoordinate(c.x+1, c.y)}
	} else if c.x == 0 && c.y > 0 && c.y < height-1 {
		// Leftmost column
		return []coordinate{*newCoordinate(c.x, c.y-1), *newCoordinate(c.x+1, c.y), *newCoordinate(c.x, c.y+1)}
	} else if c.y == height-1 && c.x > 0 && c.x < width-1 {
		// Bottom row
		return []coordinate{*newCoordinate(c.x-1, c.y), *newCoordinate(c.x, c.y-1), *newCoordinate(c.x+1, c.y)}
	} else if c.x == width-1 && c.y > 0 && c.y < height-1 {
		// Rightmost column
		return []coordinate{*newCoordinate(c.x, c.y-1), *newCoordinate(c.x-1, c.y), *newCoordinate(c.x, c.y+1)}
	} else {
		// Middle of matrix
		return []coordinate{*newCoordinate(c.x-1, c.y), *newCoordinate(c.x, c.y-1), *newCoordinate(c.x+1, c.y), *newCoordinate(c.x, c.y+1)}
	}
}

func lowestAdjacent(c coordinate, wh int8, traversed map[coordinate]struct{}, points map[coordinate]int8) coordinate {
	adjacent := c.adjacent(wh, wh)
	// Find the lowest point
	lowest := c
	for _, c2 := range adjacent {
		traversed[lowest] = struct{}{}
		if points[lowest] > points[c2] {
			lowest = c2
		}
	}

	// Exit condition
	if lowest == c {
		return lowest
	}

	return lowestAdjacent(lowest, wh, traversed, points)
}

func Part1() {
	start := time.Now()
	points := loadInput()

	// Square matrix
	wh := math.Sqrt(float64(len(points)))

	traversed := make(map[coordinate]struct{})
	lowest := make(map[coordinate]int8)

	for c := range points {
		// Do not check if coordinate was already traversed
		if _, ok := traversed[c]; !ok {
			low := lowestAdjacent(c, int8(wh), traversed, points)
			// Probably a low with a value of 9 is not a low...
			if points[low] < 9 {
				lowest[low] = points[low]
			}
		}
	}

	risk := 0
	for _, val := range lowest {
		risk += int(val) + 1
	}

	fmt.Println(lowest)
	fmt.Println("---------------------------")
	fmt.Println("Risk:", risk)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func findBasin(next []coordinate, wh int8, basin map[coordinate]int8, points map[coordinate]int8) {
	if next == nil || len(next) == 0 {
		return
	}
	// A basin is complete where all edges have a height of 9. This means the recursion should end
	// if all adjacent coordinates have a value of 9 or have already been traversed
	for _, c := range next {
		// not an edge of the basin or not already checked
		if _, ok := basin[c]; !ok && points[c] < 9 {
			basin[c] = points[c]

		}
	}
}

func lowestBasin(c coordinate, wh int8, basin map[coordinate]int8, points map[coordinate]int8) {
	// Add the current if not an edge
	if points[c] < 9 {
		basin[c] = points[c]
	}

	adjacent := c.adjacent(wh, wh)
	edge := true
	for _, c2 := range adjacent {
		// not an edge of the basin or not already checked
		if _, ok := basin[c2]; !ok && points[c2] < 9 {
			basin[c2] = points[c2]
			edge = false
			lowestBasin(c2, wh, basin, points)
		}
	}

	if edge {
		return
	}
}

func Part2() {
	start := time.Now()
	points := loadInput()
	// Square matrix
	wh := math.Sqrt(float64(len(points)))

	traversed := make(map[coordinate]struct{})
	var basins []int

	for c := range points {
		// Skip previous basins coordinates
		if _, ok := traversed[c]; ok {
			continue
		}

		basinMap := make(map[coordinate]int8)
		lowestBasin(c, int8(wh), basinMap, points)
		basins = append(basins, len(basinMap))

		// Append to traversed
		for c2 := range basinMap {
			traversed[c2] = struct{}{}
		}

		fmt.Println("------------------------------")
		fmt.Println(basinMap)
	}

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}
