package main

import (
	"fmt"
	"strings"

	"github.com/arizard/aoc-2020-solutions/pkg/arie"
)

type tile int
type grid map[int]map[int]tile

const (
	empty tile = iota
	tree
)

var tileMap = map[string]tile{
	".": empty,
	"#": tree,
}

func main() {

	lines := arie.ReadSTDINLines()
	forest := grid{}
	for row, line := range lines {
		forest[row] = map[int]tile{}
		chars := strings.Split(line, "")
		for col, ch := range chars {
			if t, ok := tileMap[ch]; ok {
				forest[row][col] = t
			}
		}
	}

	fmt.Printf("1-1 hit %d trees\n", countTrees(forest, 1, 1))
	fmt.Printf("3-1 hit %d trees\n", countTrees(forest, 3, 1))
	fmt.Printf("5-1 hit %d trees\n", countTrees(forest, 5, 1))
	fmt.Printf("7-1 hit %d trees\n", countTrees(forest, 7, 1))
	fmt.Printf("1-2 hit %d trees\n", countTrees(forest, 1, 2))
}

func countTrees(forest grid, dx int, dy int) (treeCount int) {
	treeCount = 0
	playerRow := 0
	playerColumn := 0
	width := len(forest[0])

	for playerRow <= len(forest) {
		tile := forest[playerRow][playerColumn % width]

		if tile == tree {
			treeCount++
		}

		playerRow += dy
		playerColumn += dx
	}

	return treeCount
}