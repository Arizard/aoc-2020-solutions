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
		occurs := 0

		for i, char := range record.password {
			index := i + 1

			if (index == record.p1 || index == record.p2) && char == record.char {
				occurs++
			}
		}

		if occurs != 1 {
			fails++
			continue
		}
		passes++
	}

	fmt.Printf("Passed: %d, Failed: %d\n", passes, fails)
}

type passwordRecord struct {
	p1       int
	p2       int
	char     rune
	password []rune
}

func newPasswordRecordFromString(line string) *passwordRecord {
	pattern := regexp.MustCompile(`(\d+)-(\d+)\s([a-z]):\s(\w+)`)
	groups := pattern.FindStringSubmatch(line)
	if len(groups) > 0 {
		p1, _ := strconv.Atoi(groups[1])
		p2, _ := strconv.Atoi(groups[2])
		return &passwordRecord{
			p1:       p1,
			p2:       p2,
			char:     rune([]byte(groups[3])[0]),
			password: []rune(groups[4]),
		}
	}

	return nil
}