package day10

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func loadInput() []string {
	file, err := os.Open("pkg/day10/input.txt")
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
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func corruptionScore(char string) int {
	switch char {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	default:
		return 0
	}
}

func lineCorruption(line string) []string {
	st := StringStack{}
	var corruptedChars []string
	for _, val := range strings.Split(line, "") {
		// Add opening characters to the stack
		if val == "(" || val == "[" || val == "{" || val == "<" {
			st.Push(val)
		} else if val == ")" || val == "]" || val == "}" || val == ">" {
			// Remove closing characters from stack
			stored := *st.Pop()
			// If closing characters do not match, then the line is corrupted
			if stored == "(" && val != ")" {
				corruptedChars = append(corruptedChars, val)
				fmt.Println(line, "-", "Expected ), but found", val, "instead")

			} else if stored == "[" && val != "]" {
				corruptedChars = append(corruptedChars, val)
				fmt.Println(line, "-", "Expected ], but found", val, "instead")

			} else if stored == "{" && val != "}" {
				corruptedChars = append(corruptedChars, val)
				fmt.Println(line, "-", "Expected }, but found", val, "instead")

			} else if stored == "<" && val != ">" {
				corruptedChars = append(corruptedChars, val)
				fmt.Println(line, "-", "Expected >, but found", val, "instead")
			}
		}
	}

	return corruptedChars
}

func Part1() {
	start := time.Now()
	lines := loadInput()

	corruptions := make(map[string]int)
	for _, line := range lines {
		corruption := lineCorruption(line)
		if len(corruption) > 0 {
			// Corrupted line
			if val, ok := corruptions[corruption[0]]; ok {
				corruptions[corruption[0]] = val + 1
			} else {
				corruptions[corruption[0]] = 1
			}
		}
	}

	fmt.Println("---------------------------")
	fmt.Println("Corrupted chars:", corruptions)

	sum := 0
	for key, val := range corruptions {
		sum = sum + (corruptionScore(key) * val)
	}

	fmt.Println("---------------------------")
	fmt.Println("Score:", sum)

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}

func completionScore(char string) int {
	switch char {
	case ")":
		return 1
	case "]":
		return 2
	case "}":
		return 3
	case ">":
		return 4
	default:
		return 0
	}
}

func completeLine(line string) []string {
	st := StringStack{}
	for _, val := range strings.Split(line, "") {
		// Add opening characters to the stack
		if val == "(" || val == "[" || val == "{" || val == "<" {
			st.Push(val)
		} else if val == ")" || val == "]" || val == "}" || val == ">" {
			// Remove closing characters from stack
			_ = *st.Pop()
		}
	}

	var missing []string
	for !st.IsEmpty() {
		stored := *st.Pop()
		if stored == "(" {
			missing = append(missing, ")")
		} else if stored == "[" {
			missing = append(missing, "]")
		} else if stored == "{" {
			missing = append(missing, "}")
		} else if stored == "<" {
			missing = append(missing, ">")
		}
	}

	fmt.Println(line, "- Complete by adding", strings.Join(missing, ""))
	return missing
}

func Part2() {
	start := time.Now()
	lines := loadInput()

	var scores []int
	for _, line := range lines {
		// Line is not corrupt
		if len(lineCorruption(line)) == 0 {
			score := 0
			for _, c := range completeLine(line) {
				score = (score * 5) + completionScore(c)
			}

			scores = append(scores, score)
		}
	}

	// Sort to get the middle one
	sort.Ints(scores)

	fmt.Println("---------------------------")
	fmt.Println("Scores:", scores)
	fmt.Println("---------------------------")
	fmt.Println("Middle Score:", scores[len(scores)/2])

	elapsed := time.Since(start)
	fmt.Println("Function call took ", elapsed)
}
