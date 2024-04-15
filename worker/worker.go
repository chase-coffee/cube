package worker

import (
	// stdlib
	"fmt"
	// external
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"

	// internal
	"cube/task"
)

// The requirements for the worker are as follows:
// 1. Run tasks as docker containers
// 2. Accept tasks to run from a manager
// 3. Provide relevant stats to the manager for the purpose of scheduling
// Keep track of tasks and their states

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

// CollectStats will perieodically collect stats about the worker
func (worker *Worker) CollectStats() {
	fmt.Println("I will collect stats")
}

// RunTask will run a task as a docker container
func (worker *Worker) RunTask() {
	fmt.Println("I will run a task")
}

// Start Task will start a task
func (worker *Worker) StartTask() {
	fmt.Println("I will start a task")
}

// StopTask will stop a task
func (worker *Worker) StopTask() {
	fmt.Println("I will stop a task")
}
