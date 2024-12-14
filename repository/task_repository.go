package repository

import (
	"errors"

	"github.com/emanueldias01/todolist/database"
	"github.com/emanueldias01/todolist/model"
)

func CreateTask(task model.Task) model.Task{
	database.DB.Save(&task)
	return task
}

func FindTaskById(id string) (model.Task, error){
	var task model.Task
	var err error

	database.DB.Find(&task, id)

	if task.ID == 0{
		err = errors.New("Task not found")
	}

	return task, err
}

func FindAllTasks() []model.Task{
	var list []model.Task
	database.DB.Find(&list)
	return list
}