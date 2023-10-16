package main

import (
	"graph-service/db"
	"graph-service/tasks"
)

func main() {
	db.InitMysql()
	tasks.Start()
}
