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

func main() {
	var amt int
	fmt.Println("Task Management Program")
	fmt.Print("How many tasks? ")
	fmt.Scan(&amt)

	fmt.Println("Task Amount:", taskAmount(amt))

	tasks := make([]string, 0, amt) // Initialize an empty slice with capacity

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // **Consume the leftover newline from fmt.Scan**

	fmt.Println("Enter tasks:")
	for i := 0; i < amt; i++ {
		fmt.Printf("Task %d: ", i+1)
		task, _ := reader.ReadString('\n')
		//insert the task to the slice and clean the whitespace
		tasks = append(tasks, strings.TrimSpace(task))
	}

	fmt.Println("\nTasks Entered:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}
