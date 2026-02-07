package tasks

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
