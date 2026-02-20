package storage

import (
	"cligo/tasks"
	"encoding/json"
	"os"
)

func LoadTasks() []tasks.Task {
	_, err := os.Stat("storage.json")
	if err != nil {
		return []tasks.Task{}
	}
	file, err := os.ReadFile("storage.json")
	if err != nil {
		return []tasks.Task{}
	}
	var tasksList []tasks.Task
	err = json.Unmarshal(file, &tasksList)
	if err != nil {
		return []tasks.Task{}
	}
	return tasksList
}

func SaveTasks(w []tasks.Task) error {

	file, err := json.Marshal(w)
	if err != nil {
		return err
	}
	return os.WriteFile("storage.json", file, 0644)
}

