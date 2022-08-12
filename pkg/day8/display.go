package day8

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	TopSegment         int = 0
	TopLeftSegment         = 1
	TopRightSegment        = 2
	MiddleSegment          = 3
	BottomLeftSegment      = 4
	BottomRightSegment     = 5
	BottomSegment          = 6
)

func sortedString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func initDisplayTable(digits []string, segments []string, input []string) {
	// find the easy digits by number of segments
	// 0 uses 6 segments
	// 1 uses 2 segments
	// 2 uses 5 segments
	// 3 uses 5 segments
	// 4 uses 4 segments
	// 5 uses 5 segments
	// 6 uses 6 segments
	// 7 uses 3 segments
	// 8 uses 7 segments
	// 9 uses 6 segments
	// The easy ones are 1, 4, 7, and 8, as they have a unique number of segments lit up
	// Assume input and output segments have the same lit up wires
	for _, segment := range input {
		// Store sorted for easy lookup later
		sorted := sortedString(segment)

		switch len(segment) {
		case 2:
			digits[1] = sorted
		case 3:
			digits[7] = sorted
		case 4:
			digits[4] = sorted
		case 7:
			digits[8] = sorted
		}
	}

	// Find the segments by removing known segments from known numbers
	// 7 - 1 = top segment
	top := difference(strings.Split(digits[7], ""), strings.Split(digits[1], ""))
	segments[TopSegment] = strings.Join(top, "")
	// 4 + top = almost a 9 shape
	semi9 := append(strings.Split(digits[4], ""), segments[TopSegment])
	// (3, 5 or 9) - almost 9  = bottom segment
	for _, segment := range input {
		diff := difference(strings.Split(segment, ""), semi9)
		if len(diff) == 1 {
			segments[BottomSegment] = strings.Join(diff, "")
			break
		}
	}
	// almost 9 + bottom segment = 9
	digits[9] = sortedString(strings.Join(append(semi9, segments[BottomSegment]), ""))
	// 8 - 9 = bottom left segment
	segments[BottomLeftSegment] = strings.Join(difference(strings.Split(digits[8], ""), strings.Split(digits[9], "")), "")
	// Either 2, 3 or 5 have 5 segments, but only 2 has the bottom left segment
	for _, segment := range input {
		if len(segment) == 5 {
			diff := difference(strings.Split(segment, ""), strings.Split(digits[9], ""))
			// 2 - 9 = bottom left segment
			if strings.Join(diff, "") == segments[BottomLeftSegment] {
				digits[2] = segment
				break
			}
		}
	}
	// 2 - (7 + bottom left + bottom) = middle segment
	semi0 := append(strings.Split(digits[7], ""), segments[BottomLeftSegment], segments[BottomSegment])
	segments[MiddleSegment] = strings.Join(difference(strings.Split(digits[2], ""), semi0), "")
	// 8 - middle segment = 0
	digits[0] = strings.Join(difference(strings.Split(digits[8], ""), []string{segments[MiddleSegment]}), "")
	// 1 + top segment + middle segment + bottom segment = 3
	digits[3] = sortedString(strings.Join(append(strings.Split(digits[1], ""), segments[TopSegment], segments[MiddleSegment], segments[BottomSegment]), ""))
	// 8 - (7 + bottom + bottom left + middle) = top left segment
	mirror6 := append(semi0, segments[MiddleSegment])
	segments[TopLeftSegment] = strings.Join(difference(strings.Split(digits[8], ""), mirror6), "")
	// 2 - (top + middle + bottom left + bottom) = top right segment
	segments[TopRightSegment] = strings.Join(difference(strings.Split(digits[2], ""), []string{segments[TopSegment], segments[MiddleSegment], segments[BottomLeftSegment], segments[BottomSegment]}), "")
	// 1 - top right = bottom right segment
	segments[BottomRightSegment] = strings.Join(difference(strings.Split(digits[1], ""), []string{segments[TopRightSegment]}), "")
	// 5 = top + top left + middle + bottom right + bottom
	five := []string{segments[TopSegment], segments[TopLeftSegment], segments[MiddleSegment], segments[BottomRightSegment], segments[BottomSegment]}
	digits[5] = sortedString(strings.Join(five, ""))
	// 6 = five + bottom left
	digits[6] = sortedString(strings.Join(append(five, segments[BottomLeftSegment]), ""))
}

type display struct {
	input  []string
	output []string
	// Lookup table that maps segments to a digit. The digit equals the index
	digits   []string
	segments []string
}

func newDisplay(input []string, output []string) *display {
	digits := make([]string, 10)
	segments := make([]string, 7)

	initDisplayTable(digits, segments, input)
	initDisplayTable(digits, segments, output)

	// Sort input / output arrays' strings
	var sortedInput []string
	for _, s := range input {
		sortedInput = append(sortedInput, sortedString(s))
	}

	var sortedOutput []string
	for _, s := range output {
		sortedOutput = append(sortedOutput, sortedString(s))
	}

	d := &display{input: sortedInput, output: sortedOutput, digits: digits, segments: segments}
	return d
}

func (d display) GetOutputNumber() string {
	var out []string
	for _, segments := range d.output {
		for i, digit := range d.digits {
			if digit == segments {
				out = append(out, strconv.Itoa(i))
				break
			}
		}
	}

	return strings.Join(out, "")
}

func (d display) GetInputNumber() string {
	var out []string
	for _, segments := range d.input {
		for i, digit := range d.digits {
			if digit == segments {
				out = append(out, strconv.Itoa(i))
				break
			}
		}
	}

	return strings.Join(out, "")
}

func (d display) Print() {
	fmt.Println("segments", d.segments)
	fmt.Println("digits", d.digits)
	fmt.Println("input", d.input, "output", d.output)
	fmt.Println("input", d.GetInputNumber(), "output", d.GetOutputNumber())
	fmt.Println("----------------------------------------")
}
