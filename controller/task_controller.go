package controller

import (
	"fmt"
	"net/http"

	"github.com/emanueldias01/todolist/dto"
	"github.com/emanueldias01/todolist/model"
	"github.com/emanueldias01/todolist/service"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context){
	var taskDto dto.TaskCreateDTO

	if err := c.ShouldBindJSON(&taskDto); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
	}

	var(
		task model.Task
		err error
	)

	task = model.Task{Name: taskDto.Name, Description: task.Description, Goal: taskDto.Goal}

	task, err = service.CreateTask(task)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	taskURI := fmt.Sprintf("%s/tasks/%d", c.Request.Host, task.ID)

	c.JSON(http.StatusCreated, gin.H{
		"task": task,
		"uri":  taskURI,
	})
}

func FindTaskById(c *gin.Context){
	id := c.Params.ByName("id")

	task, err := service.FindTaskById(id)

	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func FindAllTasks(c *gin.Context){
	c.JSON(http.StatusOK, service.FindAllTasks())
}

func UpdateTask(c *gin.Context){
	var taskBody dto.TaskUpdateDTO
	id := c.Params.ByName("id")
	if err := c.ShouldBindJSON(&taskBody); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	taskBodyModel := model.Task{
		Name : taskBody.Name,
		Description: taskBody.Description,
	}

	task, err := service.UpdateTask(taskBodyModel, id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context){
	id := c.Params.ByName("id")
	service.DeteleteTask(id)
	c.JSON(http.StatusNoContent, nil)
}

func MarkTaskDone(c *gin.Context){
	id := c.Params.ByName("id")
	task, err := service.MarkTaskDone(id)

	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}