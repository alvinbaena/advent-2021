package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput() [][]string {
	file, err := os.Open("pkg/day2/input.txt")
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
	var res [][]string
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		res = append(res, split)
	}

	return res
}

func Part1() {
	start := time.Now()
	input := readInput()

	depth := 0
	position := 0

	for _, value := range input {
		num, err := strconv.Atoi(value[1])
		if err != nil {
			panic(err)
		}

		switch value[0] {
		case "forward":
			position += num
		case "down":
			depth += num
		case "up":
			depth -= num
		}
	}

	fmt.Println("Position:", position, "Depth:", depth, "Multiplication:", position*depth)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func Part2() {
	start := time.Now()
	input := readInput()

	depth := 0
	position := 0
	aim := 0

	for _, value := range input {
		num, err := strconv.Atoi(value[1])
		if err != nil {
			panic(err)
		}

		switch value[0] {
		case "forward":
			position += num
			depth += aim * num
		case "down":
			aim += num
		case "up":
			aim -= num
		}
	}

	fmt.Println("Position:", position, "Depth:", depth, "Aim:", aim, "Multiplication:", position*depth)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}
