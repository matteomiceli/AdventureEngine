package main

import (
	"adventureengine/helpers"
	"adventureengine/pkg/command"
)

func main() {
	helpers.ClearScreen()
	gameOver := false
	helpers.PrintIntro()
	for !gameOver {
		input := command.Input()
		command.CommandController(input)
	}
}
