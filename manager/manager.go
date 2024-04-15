package manager

import (
	"cube/task"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// Manager requirements:
// 1.  Accept requests from users to start and stop tasks
// 2. Schedule tasks onto worker machines
// 3. Keep track of tasks, their states, and the machines they are running on

type Manager struct {
	Pending        queue.Queue
	TaskDb         map[string][]*task.Task
	EventDb        map[string][]*task.TaskEvent
	Workers        []string
	WorkerTaskMap  map[string]uuid.UUID
	TasksWorkerMap map[uuid.UUID]string
}

// SelectWorker will schedule a task onto a worker
func (manager *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker")
}

// UpdateTasks will update the task database within information about tasks state
func (manager *Manager) UpdateTasks() {
	fmt.Println("I will update the task database")
}

// SendWork will send work to a workers
func (manager *Manager) SendWork() {
	fmt.Println("I will send work to a worker")
}
