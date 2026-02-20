package commands

import (
	"cligo/storage"
	"fmt"
)

func GetTasksList() {
	tasks := storage.LoadTasks()
	if len(tasks) == 0 {
		fmt.Println("The list is empty add something!")
	}
	for _, t := range tasks {
		fmt.Printf("%d. %s [Done: %v]\n", t.Id, t.Title, t.Done)
	}
}
