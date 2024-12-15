package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name string `json:"name"`
	Description string `json:"description"`
	State string `json:"state"`
	Goal *time.Time `json:"goal"`
}

func ValidadeTask(t *Task) error{

	var (
		err error
	)
	if t.Name == ""{
		err = errors.New("Name is empty")
		return err
	}
	return nil
}