package service

import (
	"errors"

	"github.com/emanueldias01/todolist/model"
	"github.com/emanueldias01/todolist/repository"
)

func CreateTask(task model.Task) (model.Task, error){

	var err error

	err = model.ValidadeTask(&task)
	
	//create task with state pending
	task.State = "PENDING"

	repository.CreateTask(task)

	return task, err
}

func FindTaskById(id string) (model.Task, error){
	var err error
	task := repository.FindTaskById(id)
	if task.ID == 0{
		err = errors.New("Task not found")
		return model.Task{}, err
	}
	return task, err
}

func FindAllTasks() []model.Task{
	return repository.FindAllTasks()
}

func UpdateTask(taskBody model.Task, id string)(model.Task,error){
	var err error
	taskRef := repository.FindTaskById(id)
	if taskRef.ID == 0{
		err = errors.New("task not found")
		return model.Task{}, err
	}
	repository.Update(taskRef, taskBody)
	taskReturn := repository.FindTaskById(id)
	return taskReturn, nil
}

func DeteleteTask(id string){
	repository.Delete(id)
}

func MarkTaskDone(id string)(model.Task, error){
	var err error
	task := repository.FindTaskById(id)
	
	if task.ID == 0{
		err = errors.New("Task not found")
		return model.Task{}, err
	}

	task.State = "DONE"
	repository.Save(task)
	return task, nil
}