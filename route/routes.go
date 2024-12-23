package route

import (
	"github.com/emanueldias01/todolist/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequest(){
	r := gin.Default()
	defer r.Run()

	r.POST("/createTask", controller.CreateTask)
	r.GET("/:id", controller.FindTaskById)
	r.GET("/all", controller.FindAllTasks)
	r.PUT("/:id", controller.UpdateTask)
}