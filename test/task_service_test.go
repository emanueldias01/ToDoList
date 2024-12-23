package service_test

import (
	"strconv"
	"testing"

	"github.com/emanueldias01/todolist/database"
	"github.com/emanueldias01/todolist/model"
	"github.com/emanueldias01/todolist/service"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDB(){
	database.DBConect()
}

func MockTask() model.Task{
	var task model.Task = model.Task{Name : "task name", Description: "task description mock", State: "Pending"}
	database.DB.Save(&task)
	ID = int(task.ID)
	return task
}

func MockDeleteTask(){
	var task model.Task
	database.DB.Delete(&task, ID)
}

func TestFindAllTaskSucess(t *testing.T){
	SetupDB()
	task := MockTask()
	defer MockDeleteTask()

	list := service.FindAllTasks()

	assert.Contains(t, list, task, "Return list, in list, must contain the task")
}

func TestFindTaskByIdSucess(t *testing.T){
	SetupDB()
	expected := MockTask()
	defer MockDeleteTask()

	result, _  := service.FindTaskById(strconv.Itoa(ID))

	assert.Equal(t, expected.Name, result.Name, "Name task equal")
	assert.Equal(t, expected.Description, result.Description, "Description task equal")
	assert.Equal(t, expected.State, result.State, "Created task now, so you state is PENDING")
}

func TestCreateTaskSucess(t *testing.T){
	SetupDB()
	var task model.Task = model.Task{Name : "task name", Description: "task description mock", State: "PENDING"}
	result,_  := service.CreateTask(task)

	assert.Equal(t, task.Name, result.Name)
	assert.Equal(t, task.Description, result.Description)
	assert.Equal(t, task.State, result.State)

	defer database.DB.Delete(&task, task.ID)
}

func TestFindTaskByIdErrorNotFound(t *testing.T){
	SetupDB()

	result,err := service.FindTaskById("123123123")
	idExpected := uint(0)

	assert.Equal(t, idExpected, result.ID, "If find non-existent task, your ID is 0")
	assert.Equal(t, "Task not found", err.Error(), "If find non-existent task, have error")
}

func TestCreateTaskErrorValidate(t *testing.T){
	SetupDB()
	var task model.Task = model.Task{Name : "",
	Description: "task description mock", State: "PENDING"}
	_, err := service.CreateTask(task)
	
	assert.Equal(t, "Name is empty", err.Error(), "If name is empty, throw error")
}

func TestUpdateTask(t *testing.T){
	SetupDB()
	MockTask()
	defer MockDeleteTask()
	var taskBody model.Task = model.Task{Name : "Task name update",
	 Description: "change description"}

	taskResult, err := service.UpdateTask(taskBody, strconv.Itoa(ID))

	assert.Equal(t, nil, err)
	assert.Equal(t, "Task name update", taskResult.Name)
	assert.Equal(t, "change description", taskResult.Description)
}
func TestDeleteTask(t *testing.T){
	SetupDB()
	MockTask()
	defer MockDeleteTask()
	service.DeteleteTask(strconv.Itoa(ID))

	result, err := service.FindTaskById(strconv.Itoa(ID))

	assert.Equal(t, uint(0), result.ID)
	assert.Equal(t, "Task not found", err.Error())
}