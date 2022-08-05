package day5

import (
	"bufio"
	"os"
)

func loadInput() {
	file, err := os.Open("pkg/day4/input.txt")
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
}

func Part1() {

}

func Part2() {

}
