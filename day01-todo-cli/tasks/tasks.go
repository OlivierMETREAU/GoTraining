package tasks

import (
	"errors"
	"fmt"
	"strings"
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
// Returns a slice containing
//   - the id of the added task if successfull, -1 if not
//   - the error if not successfully added, nil otherwise
func (tasks *Tasks) Add(description string) (int, error) {
	if description == "" {
		return -1, errors.New("Cannot add a task with an empty description.")
	}
	tasks.TaskList[tasks.TaskCount] = Task{
		Description: description,
		State:       StateTodo,
	}
	tasks.TaskCount++
	return tasks.TaskCount - 1, nil
}

// Delete a task from the task list
// Returns an error if the 'id' is not found
func (tasks *Tasks) Delete(id int) error {
	_, ok := tasks.TaskList[id]
	if !ok {
		return errors.New("Cannot find the task Id in the task list.")
	}
	delete(tasks.TaskList, id)
	return nil
}

// Check if a task id is present in the task list
// Returns true if the task id is present, false otherwise
func (tasks *Tasks) IsPresent(id int) bool {
	_, ok := tasks.TaskList[id]
	return ok
}

// List the tasks in the task list
// Return a string containing the list
func (tasks *Tasks) List() string {
	var sb strings.Builder
	for id := range tasks.TaskCount {
		task, ok := tasks.TaskList[id]
		if ok {
			sb.WriteString(fmt.Sprintf("%d - %s - %s\r\n", id, task.Description, stateName[task.State]))
		}
	}
	sb.WriteString(fmt.Sprintf("Number of created tasks : %d\r\n", tasks.TaskCount))
	return sb.String()
}

// Make the task in progress
// Return an error if the task is already inProgress or Done, or is task id is not found.
func (tasks *Tasks) Progress(id int) error {
	task, ok := tasks.TaskList[id]
	if !ok {
		return errors.New("Cannot find the task Id in the task list.")
	}
	if task.State == StateTodo {
		tasks.TaskList[id] = Task{
			Description: task.Description,
			State:       StateInProgress,
		}
		return nil
	}
	return errors.New(fmt.Sprintf("Transition to InProgress is not allowed from %s state.", stateName[task.State]))
}

// Make the task done
// Return an error if the task is already Done, or is task id is not found.
func (tasks *Tasks) Done(id int) error {
	task, ok := tasks.TaskList[id]
	if !ok {
		return errors.New("Cannot find the task Id in the task list.")
	}
	if task.State != StateDone {
		tasks.TaskList[id] = Task{
			Description: task.Description,
			State:       StateDone,
		}
		return nil
	}
	return errors.New(fmt.Sprintf("Transition to done is not allowed from %s state.", stateName[task.State]))
}
