package dto

import (
	"time"
)

type TaskCreateDTO struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Goal *time.Time `json:"goal"`
}

type TaskUpdateDTO struct{
	Name string `json:"name"`
	Description string `json:"description"`
	Goal *time.Time `json:"goal"`
}