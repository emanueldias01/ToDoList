package service

import (
	"github.com/emanueldias01/todolist/database"
	"github.com/emanueldias01/todolist/model"
	"github.com/gin-gonic/gin"
)

var ID int

func SetupRoutes() *gin.Engine{
	r := gin.Default()
	return r
}

func SetupDB(){
	database.DBConect()
}

func MockTask(){
	var task model.Task = model.Task{Name : "task name", Description: "task description mock", State: "Pending"}
	database.DB.Save(&task)
	ID = int(task.ID)
}

func MockDeleteTask(){
	var task model.Task
	database.DB.Delete(&task, ID)
}

