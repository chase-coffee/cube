package task

import (
	// stdlib
	"time"

	// external
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

// Tasks are the smallest unit of work that can be scheduled and executed
// In the context of cube, they will be run as docker containers

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

// Task represents the task to be scheduled and executed
type Task struct {
	ID            uuid.UUID
	Name          string
	State         State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	EndTime       time.Time
}

// Task even represents a way to control a task?
type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
