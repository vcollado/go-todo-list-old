package main

import (
	"fmt"
)

type task struct {
	name        string
	description string
	isDone      bool
	// todo... add field issuedAt
}

var tasksLimitCount int = 10
var tasksCount int

func addTask(tasks *chan task, name string, description string, isDone bool) {
	tasksCount++
	*tasks <- task{name: name, description: description, isDone: isDone}

	if tasksCount == tasksLimitCount {
		// close(tasks)
	}
}

func main() {

	tasks := make(chan task, tasksLimitCount)

	addTask(&tasks, "this is the name of the task", "this is the description of the task", false)

	for task := range tasks {
		fmt.Println(task)
	}

}
