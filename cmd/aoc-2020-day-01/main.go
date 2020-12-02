package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arizard/aoc-2020-solutions/pkg/arie"
)

func main() {
	if len(os.Args) < 2 {
		panic("missing argument 1: number of components")
	}

	expenses := arie.ReadSTDINIntegerLines()
	target := 2020

	components, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic("could not convert argument 1 to integer")
	}

	fmt.Printf("%d component: %+v\n", components, FindComponents(expenses, target, components))
}

func FindComponents(input []int, target int, components int) []int {
	valueMap := map[int]bool{}

	for _, value := range input {
		valueMap[value] = true
	}

	for _, expense := range input {
		diff := target - expense
		if components == 1 {
			if expense == target {
				return []int{expense}
			}
		} else if components == 2 {
			if _, ok := valueMap[diff]; ok {
				return []int{expense, diff}
			}
		} else {
			results := []int{expense}
			newInput := arie.RemoveFromIntegerSlice(input, expense)
			for _, result := range FindComponents(newInput, diff, components - 1) {

				results = append(results, result)
			}
			if len(results) > 1 {
				return results
			}
		}
	}

	return []int{}
}