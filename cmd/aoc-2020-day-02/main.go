package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/arizard/aoc-2020-solutions/pkg/arie"
)

func main() {
	lines := arie.ReadSTDINLines()

	passes := 0
	fails := 0

	for _, line := range lines {
		record := newPasswordRecordFromString(line)
		characterFrequency := map[rune]int{}

		for _, char := range record.password {
			if _, ok := characterFrequency[char]; !ok {
				characterFrequency[char] = 0
			}
			characterFrequency[char]++
		}

		if characterFrequency[record.char] < record.min || characterFrequency[record.char] > record.max {
			fails++
			continue
		}
		passes++
	}

	fmt.Printf("Passed: %d, Failed: %d\n", passes, fails)
}

type passwordRecord struct {
	min int
	max int
	char rune
	password []rune
}

func newPasswordRecordFromString(line string) *passwordRecord {
	pattern := regexp.MustCompile(`(\d+)-(\d+)\s([a-z]):\s(\w+)`)
	groups := pattern.FindStringSubmatch(line)
	if len(groups) > 0 {
		min, _ := strconv.Atoi(groups[1])
		max, _ := strconv.Atoi(groups[2])
		return &passwordRecord{
			min:      min,
			max:      max,
			char:     rune([]byte(groups[3])[0]),
			password: []rune(groups[4]),
		}
	}

	return nil
}