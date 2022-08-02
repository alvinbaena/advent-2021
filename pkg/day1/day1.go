package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func readInput() []int {
	file, err := os.Open("pkg/day1/input.txt")
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
	var res []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		res = append(res, val)
	}

	return res
}

func Part1() {
	start := time.Now()
	input := readInput()

	var increases, decreases int
	for i, value := range input {
		// The first iteration is ignored
		if i > 0 {
			prevValue := input[i-1]
			if prevValue < value {
				increases++
			} else {
				// Assume equal as decreases
				decreases++
			}
		}
	}

	fmt.Println("Increases:", increases, "Decreases:", decreases)
	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Part2() {
	start := time.Now()
	input := readInput()

	var windows [][]int

	for i, value := range input {
		if len(input)-1 > i+1 {
			row := []int{value, input[i+1], input[i+2]}
			windows = append(windows, row)
		}
	}

	fmt.Println(windows)

	// Now calculate
	var increases, decreases, same int
	for i, window := range windows {
		if i > 0 {
			prevValue := sum(windows[i-1])
			currValue := sum(window)
			if prevValue < currValue {
				increases++
			} else if prevValue > currValue {
				decreases++
			} else {
				same++
			}
		}
	}

	fmt.Println("Increases:", increases, "Decreases:", decreases, "Same:", same)
	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}
