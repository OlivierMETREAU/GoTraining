package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAnEmptyTaskList(t *testing.T) {
	tasks := New()
	assert.Equal(t, 0, tasks.TaskCount, "TaskCount shall be 0 after creation.")
	assert.Empty(t, tasks.TaskList)
}

func TestAddATaskIncrementTheTaskCounter(t *testing.T) {
	tasks := New()
	tasks.Add("New task")
	assert.Equal(t, 1, tasks.TaskCount, "TaskCount shall be 1 after adding the first task.")
	assert.NotEmpty(t, tasks.TaskList)
}

func TestAddTaskAppendTheTaskList(t *testing.T) {
	tasks := New()
	expectedTaskDescription := "New task description."
	tasks.Add(expectedTaskDescription)
	assert.Equal(t, 0, tasks.TaskList[0].Id, "Id shall be 0 for the first added task.")
	assert.Equal(t, StateTodo, tasks.TaskList[0].State, "State shall be toDo after task add.")
	assert.Equal(t, expectedTaskDescription, tasks.TaskList[0].Description)
}
