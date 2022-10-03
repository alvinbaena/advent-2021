package day9

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func loadInput() map[Coordinate]int8 {
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

	coords := make(map[Coordinate]int8)
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		for x, p := range split {
			// Convert to int8
			val, err := strconv.ParseInt(p, 10, 8)
			if err != nil {
				panic(err)
			}

			coords[*NewCoordinate(int8(x), int8(y))] = int8(val)
		}
		y++
	}

	return coords
}

func lowestAdjacent(c Coordinate, wh int8, traversed map[Coordinate]struct{}, points map[Coordinate]int8) Coordinate {
	adjacent := c.Adjacent(wh, wh)
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

func lowestPoints(wh int8, points map[Coordinate]int8) map[Coordinate]int8 {
	traversed := make(map[Coordinate]struct{})
	lowest := make(map[Coordinate]int8)

	for c := range points {
		// Do not check if Coordinate was already traversed
		if _, ok := traversed[c]; !ok {
			low := lowestAdjacent(c, wh, traversed, points)
			// Probably a low with a value of 9 is not a low...
			if points[low] < 9 {
				lowest[low] = points[low]
			}
		}
	}

	return lowest
}

func Part1() {
	start := time.Now()
	points := loadInput()

	// Square matrix
	wh := math.Sqrt(float64(len(points)))
	lowest := lowestPoints(int8(wh), points)

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

func buildBasin(c Coordinate, wh int8, basin map[Coordinate]struct{}, points map[Coordinate]int8) {
	adjacent := c.Adjacent(wh, wh)

	// If all Adjacent points have already been traversed or are edges then exit
	exit := true
	for _, c2 := range adjacent {
		if _, ok := basin[c2]; !ok && points[c2] < 9 {
			exit = false
		}
	}

	if exit {
		return
	}

	basin[c] = struct{}{}
	for _, c2 := range adjacent {
		// Do not check if Coordinate was already traversed
		if _, ok := basin[c2]; !ok && points[c2] < 9 {
			// Not an edge
			basin[c2] = struct{}{}
			buildBasin(c2, wh, basin, points)
		}
	}
}

func Part2() {
	start := time.Now()
	points := loadInput()
	// Square matrix
	wh := math.Sqrt(float64(len(points)))
	// Find the lowest points first, then build basins from there
	lowest := lowestPoints(int8(wh), points)

	// Save basin sizes
	var basins []int
	for c := range lowest {
		basin := make(map[Coordinate]struct{})
		buildBasin(c, int8(wh), basin, points)
		basins = append(basins, len(basin))
	}

	// Sort largest to the smallest size
	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	mult := 1
	for _, val := range basins[:3] {
		mult *= val
	}

	fmt.Println("Risk:", mult)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}
