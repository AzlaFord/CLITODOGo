package commands

import (
	"cligo/storage"
	"cligo/tasks"
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Action      func(args []string)
}

var CommandStruct = []Command{
	{"list", "List all Tasks ", GetTasksList},
	{"add", "Add tasks to the list ", AddTask},
	{"done", "Mark task as done ", MarkDone},
}

func GetTasksList(args []string) {
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
func MarkDone(args []string) {
	tasksList := storage.LoadTasks()
	if len(tasksList) == 0 {
		fmt.Println("There are no tasks undone")
		return
	}
	values := strings.Join(args, " ")
	num, err := strconv.Atoi(values)
	if err != nil {
		return
	}
	for i := range tasksList {

		if tasksList[i].Id == num && !tasksList[i].Done {
			tasksList[i].Done = true
			fmt.Println("Task found and Done ,use command list")
		}
	}
	storage.SaveTasks(tasksList)
}
