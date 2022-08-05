package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func loadInput() []*vent {
	file, err := os.Open("pkg/day5/input.txt")
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var vents []*vent

	for scanner.Scan() {
		line := scanner.Text()
		splitDest := strings.Split(line, " -> ")
		// Should only have 2 items
		originStr := strings.Split(splitDest[0], ",")
		x1, err := strconv.Atoi(originStr[0])
		if err != nil {
			panic(err)
		}
		y1, err := strconv.Atoi(originStr[1])
		if err != nil {
			panic(err)
		}

		// Should only have 2 items
		destStr := strings.Split(splitDest[1], ",")
		x2, err := strconv.Atoi(destStr[0])
		if err != nil {
			panic(err)
		}
		y2, err := strconv.Atoi(destStr[1])
		if err != nil {
			panic(err)
		}

		vents = append(vents, newVent(x1, y1, x2, y2))
	}

	return vents
}

func Part1() {
	start := time.Now()
	vents := loadInput()

	ventLines := make(map[coordinate]int)
	for _, vent := range vents {
		for _, line := range vent.OrthogonalLine() {
			value, exists := ventLines[line]
			if exists {
				ventLines[line] = value + 1
			} else {
				ventLines[line] = 1
			}
		}
	}

	count := 0
	for _, value := range ventLines {
		if value >= 2 {
			count++
		}
	}

	fmt.Println("Intersections", count)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func Part2() {
	start := time.Now()
	vents := loadInput()

	ventLines := make(map[coordinate]int)
	for _, vent := range vents {
		for _, line := range vent.Line() {
			value, exists := ventLines[line]
			if exists {
				ventLines[line] = value + 1
			} else {
				ventLines[line] = 1
			}
		}
	}

	count := 0
	for _, value := range ventLines {
		if value >= 2 {
			count++
		}
	}

	fmt.Println("Intersections", count)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}
