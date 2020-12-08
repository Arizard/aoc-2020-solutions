package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arizard/aoc-2020-solutions/pkg/arie"
)

type command struct {
	name string
	value int
	executions int
}

func main() {
	lines := arie.ReadSTDINLines()
	commands := []command{}

	for _, line := range lines {
		splitter := " "
		parts := strings.Split(line, splitter)
		value, _ := strconv.Atoi(parts[1])
		cmd := command{
			name: parts[0],
			value: value,
		}

		commands = append(commands, cmd);
	}

	fmt.Println(executeCommands(commands))
}


func executeCommands(commands []command) int {
	acc := 0
	index := 0
	cmd := &commands[index]

	for cmd.executions < 10 {
		switch cmd.name {
		case "nop":
			index++
		case "acc":
			acc += cmd.value
			index++
		case "jmp":
			index += cmd.value
		}
		cmd.executions++

		if index > len(commands)-1 {
			break
		}

		cmd = &commands[index]
	}

	return acc
}