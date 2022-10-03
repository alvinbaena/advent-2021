package day11

import (
	"advent-2021/pkg/day9"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func loadInput() map[day9.Coordinate]int8 {
	file, err := os.Open("pkg/day11/input.txt")
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	coords := make(map[day9.Coordinate]int8)
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

			coords[*day9.NewCoordinate(int8(x), int8(y))] = int8(val)
		}
		y++
	}

	return coords
}

func flash(c day9.Coordinate, wh int8, octopi map[day9.Coordinate]int8, flashed map[day9.Coordinate]struct{}) {
	octopi[c] = 0
	flashed[c] = struct{}{}

	adjacent := c.DiagonalAdjacent(wh, wh)
	for _, point := range adjacent {
		if _, flashedd := flashed[point]; !flashedd {
			// Only increase the power level if not flashed
			octopi[point] = octopi[point] + 1
		}
	}

	for _, point := range adjacent {
		if octopi[point] > 9 {
			// Cascading flash
			flash(point, wh, octopi, flashed)
		}
	}
}

func printGrid(octopi map[day9.Coordinate]int8, w int, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fmt.Print(octopi[*day9.NewCoordinate(int8(i), int8(j))], " ")
		}
		fmt.Println()
	}
}

func Part1() {
	start := time.Now()
	octopi := loadInput()

	// Square matrix
	wh := math.Sqrt(float64(len(octopi)))

	printGrid(octopi, int(wh), int(wh))
	fmt.Println("***************************")

	count := 0
	for i := 0; i < 100; i++ {
		flashed := make(map[day9.Coordinate]struct{})
		for c, val := range octopi {
			if _, flashedd := flashed[c]; !flashedd {
				// Only increase the power level if not flashed
				octopi[c] = val + 1
			}
		}

		for c, val := range octopi {
			if val > 9 {
				// this is a flash
				flash(c, int8(wh), octopi, flashed)
			}
		}

		count += len(flashed)
		printGrid(octopi, int(wh), int(wh))
		fmt.Println("---------------------")
	}

	fmt.Println("Flashes:", count)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func Part2() {
	start := time.Now()
	octopi := loadInput()

	// Square matrix
	wh := math.Sqrt(float64(len(octopi)))

	printGrid(octopi, int(wh), int(wh))
	fmt.Println("***************************")

	i := 0
	for {
		flashed := make(map[day9.Coordinate]struct{})
		for c, val := range octopi {
			if _, flashedd := flashed[c]; !flashedd {
				// Only increase the power level if not flashed
				octopi[c] = val + 1
			}
		}

		for c, val := range octopi {
			if val > 9 {
				// this is a flash
				flash(c, int8(wh), octopi, flashed)
			}
		}

		printGrid(octopi, int(wh), int(wh))
		fmt.Println("---------------------")
		i++

		if len(octopi) == len(flashed) {
			fmt.Println("Iterations:", i)
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}
