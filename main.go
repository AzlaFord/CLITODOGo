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
	switch command {
	case "list":
		commands.GetTasksList()
	case "add":
		commands.AddTask(args[2])
	}
}
