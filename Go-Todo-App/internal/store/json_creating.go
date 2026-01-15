package store

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Aparneesh-Patil/GO-Todo-CLI/internal/model"
)

// creates a json file containing the task details.
func CreateJson(task string) {
	taskName := task[9:]
	var tasks []model.Task
	tasks = append(tasks, model.Task{Name: taskName, Done: false})
	b, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("tasks.json", b, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nTask Added!")
}
