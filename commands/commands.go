package commands

import (
	"cligo/storage"
	"cligo/tasks"
	"fmt"
	"strings"
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
func AddTask(args []string) {
	tasksList := storage.LoadTasks()
	title := strings.Join(args, " ")
	id := 1
	if len(tasksList) > 0 {
		id = tasksList[len(tasksList)-1].Id + 1
	}
	newTask := tasks.Task{
		Id:    id,
		Title: title,
		Done:  false,
	}
	tasksList = append(tasksList, newTask)
	storage.SaveTasks(tasksList)
}
func TaskDone(args int) {
	tasksList := storage.LoadTasks()
	if len(tasksList) == 0 {
		fmt.Println("There are no tasks undone")
		return
	}

	for i := range tasksList {
		if tasksList[i].Id == args && !tasksList[i].Done {
			tasksList[i].Done = true
			fmt.Println("Task found and Done ,use command list")
		}
	}
	storage.SaveTasks(tasksList)
}
