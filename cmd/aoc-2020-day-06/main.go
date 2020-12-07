package main

import (
	"fmt"
	"strings"

	"github.com/arizard/aoc-2020-solutions/pkg/arie"
)

func main() {
	lines := arie.ReadSTDINLines()

	input := strings.Join(lines, "\n")

	groups := strings.Split(input, "\n\n")

	count := 0

	for _, group := range groups {
		responseMap := map[string]int{}
		groupLines := strings.Split(group, "\n")
		for _, char := range strings.Split(strings.ReplaceAll(group, " \r\n", ""), "") {
			if char != "\n" {
				if _, ok := responseMap[char]; !ok {
					responseMap[char] = 0
				}

				responseMap[char]++
			}
		}

		for _, quantity := range responseMap {
			if quantity == len(groupLines) {
				count++
			}
		}

	}

	fmt.Println(count)
}
