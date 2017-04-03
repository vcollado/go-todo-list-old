package main

import (
	"fmt"
	"lib"
	"time"
)

func taskListener(listenerId int, tasksToListenOver <-chan lib.Task, tasksEvents chan<- lib.Task) {
	for task := range tasksToListenOver {
		fmt.Println("listener", listenerId, "says: the following task has been added:", task)
		// let the others listeners to read it
		time.Sleep(time.Second * 1)
		// respond to the event with the recently added task
		// this will be listened from receiveTaskListener
		tasksEvents <- task
	}
}

func receiveTaskListener(tasksEventsToReceive chan lib.Task, tasksAddedCount int) {
	// receive the values of the listener channel
	for t := 0; t < tasksAddedCount; t++ {
		<-tasksEventsToReceive
	}
}

func addTask(tasks chan<- lib.Task, name, description string, done bool) {
	taskToAdd := lib.Task{Name: name, Description: description, Done: done}
	tasks <- taskToAdd
}

func main() {

	var maxTasksAllowed int = 3
	tasks := make(chan lib.Task, maxTasksAllowed)
	tasksEvents := make(chan lib.Task)

	// create 2 parallel foreground listeners that will read the tasks added
	go taskListener(0, tasks, tasksEvents)
	go taskListener(1, tasks, tasksEvents)

	addTask(tasks, "first task", "this is the description of the first task", true)
	addTask(tasks, "second task", "today we are programming in go", false)
	addTask(tasks, "last task", "some goroutines are made", true)

	// todo.. mutex to access the tasks between goroutines
	receiveTaskListener(tasksEvents, len(tasks))
}

// the ouput is the following:
// listener 1 says: the following task has been added: {first task this is the description of the first task true}
// listener 0 says: the following task has been added: {second task today we are programming in go false}
// listener 1 says: the following task has been added: {last task some goroutines are made true}
