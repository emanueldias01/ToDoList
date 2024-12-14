package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero max=30"`
	Description string `json:"description" validate:"nonzero max=350"`
	State string `json:"state"`
	Goal time.Time `json:"goal"`
}