package model

import (
	"time"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero max=30"`
	Description string `json:"description" validate:"nonzero max=350"`
	State string `json:"state"`
	Goal time.Time `json:"goal"`
}

func ValidadeTask(t *Task) error{
	if err := validator.Validate(t); err != nil{
		return err
	}
	return nil
}