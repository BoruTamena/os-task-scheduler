package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	id       int
	duration time.Duration
}

func excuteTask(task Task, mu_map map[int]*sync.Mutex, done chan<- bool) {

	mu_map[task.id].Lock() // locking for each task with a unique id

	fmt.Printf("starting task with id :%v\n", task.id)

	time.Sleep(task.duration) // doing task

	fmt.Printf(" task id :%v completed with in a %v time\n", task.id, task.duration)

	mu_map[task.id].Unlock()

	done <- true // task completed

}

func scheduleTasks(tasks []Task) {

	mapper := make(map[int]*sync.Mutex)

	for _, tsk := range tasks {

		mapper[tsk.id] = &sync.Mutex{} // creating
	}

	fmt.Println(mapper)
	// creating channel

	done_ch := make(chan bool)

	for _, tsk := range tasks {

		go excuteTask(tsk, mapper, done_ch) // excuting each tasks
	}

	for range tasks {

		<-done_ch // receving the signal
	}

}

func main() {

	tasks := []Task{
		{
			id:       1,
			duration: 1 * time.Second,
		}, {
			id:       2,
			duration: 1 * time.Second,
		},

		{

			id:       1,
			duration: 1 * time.Second,
		},
	}

	scheduleTasks(tasks) // schedule and exc the task

}
