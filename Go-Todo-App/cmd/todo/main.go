package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Name string
	Done bool
}

func main() {
	fmt.Println(`
___________        .___        ____    .__          __   
\__    ___/___   __| _/____   |    |   |__| _______/  |_ 
  |    | /  _ \ / __ |/  _ \  |    |   |  |/  ___/\   __\
  |    |(  <_> ) /_/ (  <_> ) |    |___|  |\___ \  |  |  
  |____| \____/\____ |\____/  |_______ \__/____  > |__|  
                    \/                \/       \/         `)

	fmt.Print("Welcome to the Todo List App! ")

	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println("It seems that you dont have any tasks currently. To add a task, use the \"add task <task name> \" command!")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		// check until add task has properly been written
		for !(strings.Contains(userInput, "add task ")) {
			fmt.Println("Invalid command. Please try again.")
			scanner.Scan()
			userInput = scanner.Text()
		}
		createJson(userInput)
	} else {
		fmt.Println("To continue, use one of the following commands: \n 1. add task <task name> \n 2. list tasks \n 3. quit")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		for !(strings.Contains(userInput, "add task ")) {
			fmt.Println("Invalid command. Please try again.")
			scanner.Scan()
			userInput = scanner.Text()
		}
		addTask(userInput)
	}
	loop(*file, err)
}

// creates a json file containing the task details.
func createJson(task string) {
	taskName := task[9:]
	var tasks []Task
	tasks = append(tasks, Task{taskName, false})
	b, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("tasks.json", b, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Task Added!")
}

// function used for add tasks
func addTask(task string) {
	taskName := task[9:]
	// slice of tasks
	var tasks []Task

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
	tasks = append(tasks, Task{taskName, false})
	b, erro := json.Marshal(tasks)
	if erro != nil {
		panic(erro)
	}

	// write the updated slice of Tasks and turn it into a json file
	err = os.WriteFile("tasks.json", b, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Task Added!")
}

func loop(file os.File, err error) {
	fmt.Println("Entered Loop")
}
