package main

import (
	"cligo/commands"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage : todo <command> [arguments]")
		return
	}
	command := args[1]
	for _, cmd := range commands.GetCommands() {
		if cmd.Name == command {
			cmd.Action(args[2:])
			return
		}
	}
	fmt.Println("Invalid command!")
}
