package main

import (
	"fmt"
	"strings"

	"github.com/arizard/aoc-2020-solutions/pkg/arie"
)

func main() {
	lines := arie.ReadSTDINLines()
	seats := map[int]bool{}

	for _, line := range lines {
		rowSequence := strings.ReplaceAll(strings.ReplaceAll(line[0:7], "F", "0"), "B", "1")
		colSequence := strings.ReplaceAll(strings.ReplaceAll(line[7:10], "L", "0"), "R", "1")
		seat := bsp(0, 127, rowSequence) * 8 + bsp(0, 7, colSequence)
		seats[seat] = true
	}

	for i := 8; i < 1015; i++ {
		if _, ok := seats[i]; !ok {
			_, prevOk := seats[i-1]
			_, nextOk := seats[i+1]
			if prevOk && nextOk {
				fmt.Println("Missing seat", i)
			}
		}
	}
}

func bsp(min int, max int, sequence string) int {
	for _, bit := range sequence {
		if bit == '0' {
			max = min + (max-min)/2
		}
		if bit == '1' {
			min = max - (max-min)/2
		}
	}
	if min == max {
		return min
	}
	panic("sequence too short")
}
