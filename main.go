package main

import (
	"github.com/emanueldias01/todolist/database"
	"github.com/emanueldias01/todolist/route"
)

func main() {
	database.DBConect()
	route.HandleRequest()
}