package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Aparneesh-Patil/GO-Todo-CLI/internal/store"
)

func Run() {
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
			store.CreateJson(userInput)
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
			LoadTasks()
		} else if strings.EqualFold(userInput, "quit") {
			fmt.Println("\nSee you soon!")
			defer file.Close()
			return
		} else if strings.Contains(userInput, "complete ") {
			CompleteTasks(userInput)
		} else if strings.Contains(userInput, "delete ") {
			DeleteTask(userInput)
		} else {
			AddTask(userInput)
		}
	}
	loop(*file, err)
}

func loop(file os.File, err error) {
	userInput := ""

	for !strings.EqualFold(userInput, "quit") {
		fmt.Print("\nWhat would you like to do? \n 1. add task <task name> \n 2. list tasks \n 3. complete <task name> \n 4. delete <task name> \n 5. quit\n> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput = scanner.Text()

		if strings.Contains(userInput, "add task ") {
			AddTask(userInput)
		} else if strings.EqualFold(userInput, "list tasks") {
			LoadTasks()
		} else if strings.Contains(userInput, "complete ") {
			CompleteTasks(userInput)
		} else if strings.Contains(userInput, "delete ") {
			DeleteTask(userInput)
		} else if strings.EqualFold(userInput, "quit") {
			break
		} else {
			fmt.Println("Invalid command. ")
		}
	}

	fmt.Print("See you soon!")
	defer file.Close()
}
