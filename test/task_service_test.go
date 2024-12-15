package service_test

import (
	"strconv"
	"testing"

	"github.com/emanueldias01/todolist/controller"
	"github.com/emanueldias01/todolist/database"
	"github.com/emanueldias01/todolist/model"
	"github.com/emanueldias01/todolist/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutes() *gin.Engine{
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

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
	r := SetupRoutes()
	r.GET("/all", controller.FindAllTasks)

	SetupDB()
	task := MockTask()
	defer MockDeleteTask()

	list := service.FindAllTasks()

	assert.Contains(t, list, task, "Return list, in list, must contain the task")
}

func TestFindTaskByIdSucess(t *testing.T){
	r := SetupRoutes()
	r.GET("/all", controller.FindAllTasks)

	SetupDB()
	expected := MockTask()
	defer MockDeleteTask()

	result, _  := service.FindTaskById(strconv.Itoa(ID))

	assert.Equal(t, expected.Name, result.Name, "Name task equal")
	assert.Equal(t, expected.Description, result.Description, "Description task equal")
	assert.Equal(t, expected.State, result.State, "Created task now, so you state is PENDING")
}

func TestCreateTaskSucess(t *testing.T){
	r := SetupRoutes()
	r.GET("/all", controller.FindAllTasks)

	SetupDB()
	var task model.Task = model.Task{Name : "task name", Description: "task description mock", State: "PENDING"}
	result,_  := service.CreateTask(task)

	assert.Equal(t, task.Name, result.Name)
	assert.Equal(t, task.Description, result.Description)
	assert.Equal(t, task.State, result.State)

	defer database.DB.Delete(&task, task.ID)
}