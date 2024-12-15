package service

import (
	"github.com/emanueldias01/todolist/model"
	"github.com/emanueldias01/todolist/repository"
)

func CreateTask(task model.Task ) (model.Task, error){

	var err error

	err = model.ValidadeTask(&task)
	
	//create task with state pending
	task.State = "PENDING"

	repository.CreateTask(task)

	return task, err
}

func FindTaskById(id string) (model.Task, error){

	task, err := repository.FindTaskById(id)

	return task, err
}

func FindAllTasks() []model.Task{
	return repository.FindAllTasks()
}