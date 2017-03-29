package main

import (
	"fmt"
	"log"
)

type task struct {
	name        string
	description string
	isDone      bool
}

func addTask(tasks *chan *task, name string, description string, isDone bool) {
	fmt.Println("adding the task, please wait...")
	*tasks <- &task{name, description, isDone}
}

func main() {

	tasks := make(chan *task)

	go addTask(&tasks, "first task", "this is the description of the first task", false)
	go addTask(&tasks, "second task", "today we are programming in go", false)

	for i := 0; i <= 1; i++ {
		select {
		case task := <-tasks:
			log.Print("task added: ", task)
		}
	}
}
