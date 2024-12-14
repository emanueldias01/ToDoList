package service

import (
	"github.com/emanueldias01/todolist/model"
	"github.com/emanueldias01/todolist/repository"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context){
	var task model.Task

	if err := c.ShouldBindJSON(task); err != nil{
		c.JSON(400, gin.H{
			"message" : "error in should bind JSON",
		})
	}



	if err := model.ValidadeTask(&task); err != nil{
		c.JSON(400, err.Error())
	}

	//create task with state pending
	task.State = "PENDING"

	repository.CreateTask(task)
}

func FindTaskById(c *gin.Context){
	id := c.Params.ByName("id")

	task, err := repository.FindTaskById(id)

	if err != nil{
		c.JSON(404, gin.H{
			"message" : err.Error(),
		})

		return
	}

	c.JSON(200, task)
}

func FindAllTasks(c *gin.Context){
	c.JSON(200, repository.FindAllTasks())
}