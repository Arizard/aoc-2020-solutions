package arie

import (
	"bufio"
	"os"
	"strconv"
)

func ReadSTDINIntegerLines() []int {
	var lines []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			if line, err := strconv.Atoi(text); err == nil {
				lines = append(lines, line)
			}
		}
	}
	return lines
}

func ReadSTDINLines() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}