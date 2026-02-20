package main

import (
	"cligo/commands"
	"fmt"
	"os"
	"strconv"
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
		commands.AddTask(args[2:])
	case "done":
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid task id")
			return
		}
		commands.TaskDone(id)
	}
}
