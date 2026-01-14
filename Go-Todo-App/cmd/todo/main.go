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
		fmt.Print("It seems that you dont have any tasks currently. To continue, use one of the following commands: \n 1. add task <task name> \n 2. quit \n> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		// check until add task or quit has properly been written
		for !(strings.Contains(userInput, "add task ")) && !(strings.EqualFold(userInput, "quit")) {
			fmt.Print("Invalid command. Please try again.\n> ")
			scanner.Scan()
			userInput = scanner.Text()
		}
		switch userInput {
		case "quit":
			fmt.Println("See you soon!")
			defer file.Close()
			return
		default:
			createJson(userInput)
			file, err = os.Open("tasks.json")
		}
	} else {
		fmt.Print("To continue, use one of the following commands: \n 1. add task <task name> \n 2. list tasks \n 3. complete <task name> \n 4. delete <task name> \n 5. quit\n> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		// check until one of the commands is written correcty
		for !(strings.Contains(userInput, "add task ")) && !(strings.EqualFold(userInput, "list tasks")) && !(strings.EqualFold(userInput, "quit")) && !(strings.Contains(userInput, "complete ")) && !(strings.Contains(userInput, "delete ")) {
			fmt.Print("\nInvalid command. Please try again.\n> ")
			scanner.Scan()
			userInput = scanner.Text()
		}

		// Different commands lead to different functions
		if strings.EqualFold(userInput, "list tasks") {
			loadTasks()
		} else if strings.EqualFold(userInput, "quit") {
			fmt.Println("\nSee you soon!")
			defer file.Close()
			return
		} else if strings.Contains(userInput, "complete ") {
			completeTasks(userInput)
		} else if strings.Contains(userInput, "delete ") {
			deleteTask(userInput)
		} else {
			addTask(userInput)
		}
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
	fmt.Println("\nTask Added!")
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

	fmt.Println("\nTask Added!")
}

// loads tasks from the json file
func loadTasks() {
	var tasks []Task
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

func completeTasks(task string) {
	taskName := task[9:]
	var tasks []Task

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

func deleteTask(task string) {
	taskName := task[7:]

	var tasks []Task
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

func loop(file os.File, err error) {
	userInput := ""

	for !strings.EqualFold(userInput, "quit") {
		fmt.Print("\nWhat would you like to do? \n 1. add task <task name> \n 2. list tasks \n 3. complete <task name> \n 4. delete <task name> \n 5. quit\n> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput = scanner.Text()

		if strings.Contains(userInput, "add task ") {
			addTask(userInput)
		} else if strings.EqualFold(userInput, "list tasks") {
			loadTasks()
		} else if strings.Contains(userInput, "complete ") {
			completeTasks(userInput)
		} else if strings.Contains(userInput, "delete ") {
			deleteTask(userInput)
		} else if strings.EqualFold(userInput, "quit") {
			break
		} else {
			fmt.Println("Invalid command. ")
		}
	}

	fmt.Print("See you soon!")
	defer file.Close()
}
