package tasks

import (
	"testing"

	"example.com/day01-todo-cli/jsonmanager"
	"github.com/stretchr/testify/assert"
)

func TestCreateAnEmptyTaskList(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	assert.Equal(t, 0, tasks.TaskCount, "TaskCount shall be 0 after creation.")
	assert.Empty(t, tasks.TaskList)
}

func TestAddATaskIncrementTheTaskCounter(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("New task")
	assert.Equal(t, 1, tasks.TaskCount, "TaskCount shall be 1 after adding the first task.")
	assert.NotEmpty(t, tasks.TaskList)
}

func TestAddTaskAppendTheTaskList(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	expectedTaskDescription := "New task description."
	tasks.Add(expectedTaskDescription)
	assert.Equal(t, StateTodo, tasks.TaskList[0].State, "State shall be toDo after task add.")
	assert.Equal(t, expectedTaskDescription, tasks.TaskList[0].Description)
}

func TestAddTaskWithValidTaskReturnsNoError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	expectedTaskDescription := "New task description."
	i, err := tasks.Add(expectedTaskDescription)
	assert.Equal(t, 0, i, "Add shall return no error when adding a valid task.")
	assert.Equal(t, nil, err, "Add shall return no error when adding a valid task.")
}

func TestAddTaskWithNonValidTaskReturnsAnError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	expectedTaskDescription := ""
	i, err := tasks.Add(expectedTaskDescription)
	assert.NotEqual(t, 0, i, "Add shall return an error when adding a non valid task.")
	assert.NotEqual(t, nil, err, "Add shall return an error when adding a non valid task.")
}

func TestAddTaskWithNonValidTaskDoesNotAddAnyTask(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	expectedTaskDescription := ""
	tasks.Add(expectedTaskDescription)
	assert.Equal(t, 0, tasks.TaskCount, "Add shall not add a non valid task.")
	assert.Empty(t, tasks.TaskList, "Add shall not add a non valid task.")
}

func TestAddingMultipleTasks(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("First task")
	tasks.Add("Second task")
	assert.Equal(t, 2, tasks.TaskCount)
	assert.Equal(t, "First task", tasks.TaskList[0].Description)
	assert.Equal(t, "Second task", tasks.TaskList[1].Description)
}

func TestDeleteAnUnknownTaskReturnsAnError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	err := tasks.Delete(0)
	assert.NotEqual(t, nil, err)
}

func TestDeleteAKnownTaskReturnsNoError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("New Task")
	err := tasks.Delete(0)
	assert.Equal(t, nil, err)
}

func TestDeleteAKnownTaskRemoveTheTaskFromTheList(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("New Task")
	tasks.Delete(0)
	assert.Empty(t, tasks.TaskList)
}

func TestDeleteAKnowTaskRemoveTheCorrectTask(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("First task")
	i, _ := tasks.Add("Second task")
	tasks.Add("Third task")
	tasks.Delete(i)
	assert.Equal(t, true, tasks.IsPresent(0))
	assert.Equal(t, false, tasks.IsPresent(1))
	assert.Equal(t, true, tasks.IsPresent(2))
}

func TestListWithExistingTasks(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("First task")
	tasks.Add("Second task")
	output := tasks.List()
	assert.Equal(t, "0 - First task - toDo\r\n1 - Second task - toDo\r\nNumber of created tasks : 2\r\n", output)
}

func TestListAnEmptyTaskList(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	output := tasks.List()
	assert.Equal(t, "Number of created tasks : 0\r\n", output)
}

func TestProgressWithUnknownTaskReturnsAnError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	err := tasks.Progress(0)
	assert.NotEqual(t, nil, err)
}

func TestProgressFromTodoReturnsNoError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("First task.")
	err := tasks.Progress(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, StateInProgress, tasks.TaskList[0].State)
}

func TestProgressFromProgressReturnsAnError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("First task.")
	tasks.Progress(0)
	err := tasks.Progress(0)
	assert.NotEqual(t, nil, err)
}

func TestProgressFromDoneReturnsAnError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("First task.")
	tasks.Done(0)
	err := tasks.Progress(0)
	assert.NotEqual(t, nil, err)
}

func TestDoneWithUnknownTaskReturnsAnError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	err := tasks.Done(0)
	assert.NotEqual(t, nil, err)
}

func TestDoneFromTodoReturnsNoError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("First task.")
	err := tasks.Done(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, StateDone, tasks.TaskList[0].State)
}

func TestDoneFromProgressReturnsNoError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("First task.")
	tasks.Progress(0)
	err := tasks.Done(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, StateDone, tasks.TaskList[0].State)
}

func TestSaveReturnsNoError(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("First task.")
	tasks.Progress(0)
	tasks.Add("Second task.")
	tasks.Save()
}

func TestReadFromAFile(t *testing.T) {
	tasks := New(jsonmanager.New("taskList.json"))
	tasks.Add("First task")
	tasks.Progress(0)
	tasks.Add("Second task")
	tasks.Save()

	tasksRead := New(jsonmanager.New("taskList.json"))
	tasksRead.ReadFromFile()
	output := tasksRead.List()
	assert.Equal(t, "0 - First task - inProgress\r\n1 - Second task - toDo\r\nNumber of created tasks : 2\r\n", output)
}
