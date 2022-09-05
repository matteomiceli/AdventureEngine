package main

import (
	"saran/pkg/command"
)

func main() {
	gameEnd := false
	for !gameEnd {
		input := command.Input()
		command.CommandController(input)
	}
}
