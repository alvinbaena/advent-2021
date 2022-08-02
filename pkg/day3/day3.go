package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput() [][]string {
	file, err := os.Open("pkg/day3/input.txt")
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
		split := strings.Split(scanner.Text(), "")
		res = append(res, split)
	}

	return res
}

// Get the most common bit for the passed vertical index
func mostCommon(slices [][]string, i int) string {
	one := 0
	zero := 0

	for _, slice := range slices {
		switch slice[i] {
		case "0":
			zero++
		case "1":
			one++
		}
	}

	if zero > one {
		return "0"
	} else {
		return "1"
	}
}

func Part1() {
	start := time.Now()
	input := readInput()

	var gamma, epsilon string

	// Let's cheat a bit, we know each line has 12 bits, so just iterate using the known length
	for i := 0; i < 12; i++ {
		most := mostCommon(input, i)
		gamma += most

		// This could also be achieved by negating (NOT) the gamma binary number, but this is simpler
		if most == "0" {
			epsilon += "1"
		} else {
			epsilon += "0"
		}
	}

	gammaDec, err := strconv.ParseInt(gamma, 2, 32)
	if err != nil {
		panic(err)
	}

	epsilonDec, err := strconv.ParseInt(epsilon, 2, 32)
	if err != nil {
		panic(err)
	}

	fmt.Println("Gamma:", gamma, "Gamma Decimal:", gammaDec)
	fmt.Println("Epsilon:", epsilon, "Epsilon Decimal:", epsilonDec)
	fmt.Println("Multiplication:", gammaDec*epsilonDec)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func findByCommonality(slices [][]string, i int, negate bool) string {
	one := 0
	zero := 0

	for _, slice := range slices {
		switch slice[i] {
		case "0":
			zero++
		case "1":
			one++
		}
	}

	var common string
	if negate {
		// For equal amounts common is 0
		if zero > one {
			common = "1"
		} else {
			common = "0"
		}
	} else {
		// For equal amounts common is 1
		if one >= zero {
			common = "1"
		} else {
			common = "0"
		}
	}

	var commonSlices [][]string
	for _, slice := range slices {
		if slice[i] == common {
			commonSlices = append(commonSlices, slice)
		}
	}

	if len(commonSlices) > 1 {
		return findByCommonality(commonSlices, i+1, negate)
	} else {
		// End of recursion. Should only be one element left
		return strings.Join(commonSlices[0], "")
	}
}

func Part2() {
	start := time.Now()
	input := readInput()

	o2 := findByCommonality(input, 0, false)
	co2 := findByCommonality(input, 0, true)

	o2Dec, err := strconv.ParseInt(o2, 2, 32)
	if err != nil {
		panic(err)
	}

	co2Dec, err := strconv.ParseInt(co2, 2, 32)
	if err != nil {
		panic(err)
	}

	fmt.Println("O2:", o2, "O2 Decimal:", o2Dec)
	fmt.Println("CO2:", co2, "CO2 Decimal:", co2Dec)
	fmt.Println("Multiplication:", o2Dec*co2Dec)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}
