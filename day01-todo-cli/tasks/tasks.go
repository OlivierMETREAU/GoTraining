package tasks

import (
	"errors"
)

// Task state - can be Todo, InProgress or Done
type TaskState int

const (
	StateTodo TaskState = iota
	StateInProgress
	StateDone
)

var stateName = map[TaskState]string{
	StateTodo:       "toDo",
	StateInProgress: "inProgress",
	StateDone:       "done",
}

// Task - include Id, Description, State
type Task struct {
	Description string
	Id          int
	State       TaskState
}

type Tasks struct {
	TaskCount int
	TaskList  map[int]Task
}

// Create an empty list of tasks
func New() Tasks {
	return Tasks{
		TaskCount: 0,
		TaskList:  map[int]Task{},
	}
}

// Add a task to the task list
func (tasks *Tasks) Add(description string) (int, error) {
	if description == "" {
		return -1, errors.New("Cannot add a task with an empty description.")
	}
	tasks.TaskList[tasks.TaskCount] = Task{
		Description: description,
		Id:          tasks.TaskCount,
		State:       StateTodo,
	}
	tasks.TaskCount++
	return tasks.TaskCount - 1, nil
}

func (tasks *Tasks) Delete(id int) error {
	_, ok := tasks.TaskList[id]
	if !ok {
		return errors.New("Cannot find the task Id in the task list.")
	}
	delete(tasks.TaskList, id)
	return nil
}

func (tasks *Tasks) IsPresent(id int) bool {
	_, ok := tasks.TaskList[id]
	return ok
}
