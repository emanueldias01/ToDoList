package database

import (
	"log"
	"time"

	"github.com/emanueldias01/todolist/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	DB *gorm.DB
	err error
)

func DBConect(){

	//wait 10 seconds for start connection to db
	time.Sleep(10 * time.Second)
	strCon := "host=postgres user=root password=root dbname=ToDoList port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strCon))

	if err != nil{
		log.Fatal(err.Error())
	}

	DB.AutoMigrate(&model.Task{})
}