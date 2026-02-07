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

// Add a task to the task list
func (tasks *Tasks) Add(description string) {
	tasks.TaskList[tasks.TaskCount] = Task{
		Description: description,
		Id:          tasks.TaskCount,
		State:       StateTodo,
	}
	tasks.TaskCount++
}
