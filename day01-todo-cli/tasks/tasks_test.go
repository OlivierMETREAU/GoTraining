package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAnEmptyTaskList(t *testing.T) {
	taskList := New()
	assert.Equal(t, 0, taskList.TaskCount, "TaskCount shall be 0 after creation.")
	assert.Empty(t, taskList.TaskList)
}
