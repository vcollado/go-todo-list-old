package main

import (
	"fmt"
	"lib"
	"log"
)

func addTask(tasks *chan *lib.Task, name string, description string, done bool) {
	fmt.Println("adding the task, please wait...")
	*tasks <- &lib.Task{Name: name, Description: description, Done: done}
}

func main() {

	tasks := make(chan *lib.Task)

	go addTask(&tasks, "first task", "this is the description of the first task", true)
	go addTask(&tasks, "second task", "today we are programming in go", false)

	for i := 0; i <= 1; i++ {
		select {
		case task := <-tasks:
			log.Print("task added: ", task)
			if task.IsDone() {
				log.Print("and its done")
			}
		}
	}

}
