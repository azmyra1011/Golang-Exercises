package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func taskAmount(amt int) int {
	return amt
}

// Function to display tasks with status
func read(tasks []string, completedTasks []bool) {
	fmt.Println("\nTasks Left:")
	for i, task := range tasks {
		status := "❌"
		if completedTasks[i] {
			status = "✔"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task)
	}
}

// Function to mark a task as done
func markTaskAsDone(tasks []string, completedTasks []bool) {
	var taskNum int
	fmt.Print("Enter task number to mark as done: ")
	fmt.Scan(&taskNum)

	if taskNum > 0 && taskNum <= len(tasks) {
		if completedTasks[taskNum-1] {
			fmt.Println("Task is already marked as done! ✔")
		} else {
			completedTasks[taskNum-1] = true
			fmt.Println("Task marked as done! ✔")
		}
	} else {
		fmt.Println("Invalid task number.")
	}
}

func main() {
	var amt int
	var action = []string{"Read Task List", "Mark Task as Done", "Update Task", "Delete Task", "Done"}
	var user_action int

	fmt.Println("Task Management Program")
	fmt.Print("How many tasks? ")
	fmt.Scan(&amt)

	fmt.Println("Task Amount:", taskAmount(amt))

	tasks := make([]string, 0, amt)
	completedTasks := make([]bool, amt)

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	fmt.Println("Enter tasks:")
	for i := 0; i < amt; i++ {
		fmt.Printf("Task %d: ", i+1)
		task, _ := reader.ReadString('\n')
		tasks = append(tasks, strings.TrimSpace(task))
	}

	// Loop to handle user actions
	for {
		read(tasks, completedTasks)

		fmt.Println("\nWhat do you want to do?")
		fmt.Println("Actions:")
		for i, act := range action {
			fmt.Println(i+1, "-", act)
		}

		fmt.Print("Your action: ")
		fmt.Scan(&user_action)

		switch user_action {
		case 1:
			read(tasks, completedTasks)
		case 2:
			markTaskAsDone(tasks, completedTasks)
		case 3:
			fmt.Println("Update feature not implemented yet.")
		case 4:
			fmt.Println("Delete feature not implemented yet.")
		case 5:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
