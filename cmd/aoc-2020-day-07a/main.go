package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/arizard/aoc-2020-solutions/pkg/arie"
)

func main() {
	rules := []bagRule{}
	lines := arie.ReadSTDINLines()
	bagPattern := regexp.MustCompile(`\sbag(s*)`)
	digitPattern := regexp.MustCompile(`\d+`)

	for _, line := range lines {
		splitPattern := regexp.MustCompile(`(contain|,|\.)`)
		parts := splitPattern.Split(line, -1)
		parent := parts[0]
		rule := bagRule{
			color:    strings.Trim(bagPattern.ReplaceAllString(parent, ""), " "),
			children: map[string]int{},
		}

		for _, child := range parts[1:len(parts)-1] {
			if strings.Contains(child, "no other") {
				continue
			}
			childColor := strings.Trim(bagPattern.ReplaceAllString(
				digitPattern.ReplaceAllString(child, ""),
				"",
			), " ")
			childQuantity, _ := strconv.Atoi(digitPattern.FindString(child))

			rule.children[childColor] = childQuantity
		}

		rules = append(rules, rule)
	}

	// A struct-ized ruleset
	fmt.Println("Ruleset", rules)

	colors := []string{}
	for _, rule := range rules {
		colors = append(colors, rule.color)
	}

	fmt.Println("Colors", colors)

	fmt.Println("contained", findBagsWithin("shiny gold", rules))

	total := len(findBagsWithin("shiny gold", rules))

	fmt.Println(total)
}

type bagRule struct {
	color string
	children map[string]int
}

func findBagsWithin(color string, rules []bagRule) []string {
	within := []string{}
	var currentRule bagRule

	for _, rule := range rules {
		if color == rule.color {
			currentRule = rule
			break
		}
	}

	for color, quantity := range currentRule.children {
		bagsWithin := findBagsWithin(color, rules)
		for i := 0; i < quantity; i++ {
			within = append(within, color)
			within = append(within, bagsWithin...)
		}
	}

	return within
}

