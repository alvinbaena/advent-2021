package day6

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const input = "2,5,2,3,5,3,5,5,4,2,1,5,5,5,5,1,2,5,1,1,1,1,1,5,5,1,5,4,3,3,1,2,4,2,4,5,4,5,5,5,4,4,1,3,5,1,2,2,4,2,1,1,2,1,1,4,2,1,2,1,2,1,3,3,3,5,1,1,1,3,4,4,1,3,1,5,5,1,5,3,1,5,2,2,2,2,1,1,1,1,3,3,3,1,4,3,5,3,5,5,1,4,4,2,5,1,5,5,4,5,5,1,5,4,4,1,3,4,1,2,3,2,5,1,3,1,5,5,2,2,2,1,3,3,1,1,1,4,2,5,1,2,4,4,2,5,1,1,3,5,4,2,1,2,5,4,1,5,5,2,4,3,5,2,4,1,4,3,5,5,3,1,5,1,3,5,1,1,1,4,2,4,4,1,1,1,1,1,3,4,5,2,3,4,5,1,4,1,2,3,4,2,1,4,4,2,1,5,3,4,1,1,2,2,1,5,5,2,5,1,4,4,2,1,3,1,5,5,1,4,2,2,1,1,1,5,1,3,4,1,3,3,5,3,5,5,3,1,4,4,1,1,1,3,3,2,3,1,1,1,5,4,2,5,3,5,4,4,5,2,3,2,5,2,1,1,1,2,1,5,3,5,1,4,1,2,1,5,3,5,2,1,3,1,2,4,5,3,4,3"

func copulation(school []string, days int) uint64 {
	sea := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	// Fill the sea
	for _, f := range school {
		timer, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		sea[timer] = sea[timer] + 1
	}

	// Start the counter
	for i := 0; i < days; i++ {
		// After the day all 1's become 0's and reset.
		// Move positions
		newFishes := 0
		for timer, amount := range sea {
			if timer == 0 {
				// Become a 6 later
				newFishes = amount
			} else {
				sea[timer-1] = amount
			}
		}

		// Add more 6's and 8's
		sea[6] = sea[6] + newFishes
		sea[8] = newFishes
	}

	fmt.Println("Fishes", sea)

	count := uint64(0)
	for _, fishes := range sea {
		count += uint64(fishes)
	}

	return count
}

func Part1() {
	start := time.Now()
	school := strings.Split(input, ",")
	fmt.Println("Count", copulation(school, 80))

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func Part2() {
	start := time.Now()
	elapsed := time.Since(start)

	school := strings.Split(input, ",")
	fmt.Println("Count", copulation(school, 256))

	fmt.Println("Function call took ", elapsed)
}
