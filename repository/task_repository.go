package repository

import (
	"github.com/emanueldias01/todolist/database"
	"github.com/emanueldias01/todolist/model"
)

func CreateTask(task model.Task) model.Task{
	database.DB.Save(&task)
	return task
}

func FindTaskById(id string) (model.Task){
	var task model.Task
	database.DB.Find(&task, id)
	return task
}

func FindAllTasks() []model.Task{
	var list []model.Task
	database.DB.Find(&list)
	return list
}

func Update(taskRef model.Task, taskBody model.Task){
	database.DB.Model(&taskRef).UpdateColumns(taskBody)
}