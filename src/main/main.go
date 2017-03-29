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

func addTask(tasks *chan task, name string, description string, isDone bool) {
	*tasks <- task{name, description, isDone}
}

func main() {

	tasks := make(chan task, 2)

	addTask(&tasks, "first task", "this is the description of the first task", false)
	addTask(&tasks, "second task", "today we are programming in go", false)
	close(tasks)

	for task := range tasks {
		fmt.Println(task)
	}

}
