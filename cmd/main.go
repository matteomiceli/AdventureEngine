package main

import (
	"adventureengine/pkg/command"
)

func main() {
	gameOver := false
	for !gameOver {
		input := command.Input()
		command.CommandController(input)
	}
}
