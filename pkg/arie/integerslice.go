package arie

func MaxInteger(input []int) int {
	if len(input) == 0 {
		panic("input to MaxInteger was empty")
	}
	max := input[0]
	for _, num := range input {
		if num > max {
			max = num
		}
	}

	return max
}

func RemoveFromIntegerSlice(input []int, rem int) []int {
	removals := 0
	maxRemovals := 1
	result := []int{}
	for _, integer := range input {
		if integer != rem && removals < maxRemovals {
			result = append(result, integer)
		} else {
			removals++
		}
	}
	return result
}