package app

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Aparneesh-Patil/GO-Todo-CLI/internal/model"
)

// function used for add tasks
func AddTask(task string) {
	taskName := task[9:]
	// slice of tasks
	var tasks []model.Task

	// read the exisiting json
	fileRead, err := os.ReadFile("tasks.json")
	if err != nil {
		panic(err)
	}

	// decode the read json file
	err = json.Unmarshal(fileRead, &tasks)
	if err != nil {
		panic(err)
	}

	// add the current task to the already existing json
	tasks = append(tasks, model.Task{Name: taskName, Done: false})
	b, erro := json.Marshal(tasks)
	if erro != nil {
		panic(erro)
	}

	// write the updated slice of Tasks and turn it into a json file
	err = os.WriteFile("tasks.json", b, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nTask Added!")
}

// loads tasks from the json file
func LoadTasks() {
	var tasks []model.Task
	fileRead, err := os.ReadFile("tasks.json")
	if err != nil {
		panic(err)
	}

	// decode the read json file
	err = json.Unmarshal(fileRead, &tasks)
	if err != nil {
		panic(err)
	}

	// prints out tasks using [X] task_name order if completed and [ ] task_name if not completed
	fmt.Println("\nList of all tasks: ")
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Done {
			fmt.Println("[X] " + tasks[i].Name)
		} else {
			fmt.Println("[ ] " + tasks[i].Name)
		}
	}
}

func CompleteTasks(task string) {
	taskName := task[9:]
	var tasks []model.Task

	// read the json file
	fileRead, err := os.ReadFile("tasks.json")
	if err != nil {
		panic(err)
	}

	// decode the read json file
	err = json.Unmarshal(fileRead, &tasks)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(tasks); i++ {
		if strings.EqualFold(tasks[i].Name, taskName) {
			tasks[i].Done = true
			fmt.Println("\nTask Completed!")
			b, erro := json.Marshal(tasks)
			if erro != nil {
				panic(erro)
			}
			err = os.WriteFile("tasks.json", b, 0644)
			if err != nil {
				panic(err)
			}
			return
		}
	}

	fmt.Println("\nTask doesn't exist.")
}

func DeleteTask(task string) {
	taskName := task[7:]

	var tasks []model.Task
	// read the json file
	fileRead, err := os.ReadFile("tasks.json")
	if err != nil {
		panic(err)
	}

	// decode the read json file
	err = json.Unmarshal(fileRead, &tasks)
	if err != nil {
		panic(err)
	}

	// decode the read json file
	err = json.Unmarshal(fileRead, &tasks)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(tasks); i++ {
		if strings.EqualFold(tasks[i].Name, taskName) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("\nTask Deleted!")
			b, erro := json.Marshal(tasks)
			if erro != nil {
				panic(erro)
			}
			err = os.WriteFile("tasks.json", b, 0644)
			if err != nil {
				panic(err)
			}
			return
		}
	}

	fmt.Println("\nTask doesn't exist.")
}
