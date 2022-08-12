package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func loadInput() []*display {
	file, err := os.Open("pkg/day8/input.txt")
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var readings []*display
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputOutput := strings.Split(line, " | ")
		// Only 2 on the split array
		input := strings.Split(inputOutput[0], " ")
		output := strings.Split(inputOutput[1], " ")

		readings = append(readings, newDisplay(input, output))
	}

	return readings
}

func Part1() {
	start := time.Now()
	displays := loadInput()

	sum := 0
	for _, display := range displays {
		for _, segment := range display.output {
			// Must be either 1, 4, 7, or 8
			if display.digits[1] == segment || display.digits[4] == segment || display.digits[7] == segment || display.digits[8] == segment {
				sum++
			}
		}
	}

	fmt.Println("Output", sum)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func Part2() {
	start := time.Now()
	displays := loadInput()

	sum := 0
	for _, display := range displays {
		display.Print()
		num, err := strconv.Atoi(display.GetOutputNumber())
		if err != nil {
			panic(err)
		}

		sum += num
	}

	fmt.Println("Sum", sum)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}
