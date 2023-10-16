package tasks

import (
	"graph-service/db"
	"graph-service/models"
	"graph-service/services"
)

const (
	JobTypeClaimRecord = iota + 1
)

func JobClaimRecord() {
	var taskStatus models.TaskStatus
	db.Mysql.Model(&models.TaskStatus{}).Where("task_type=?", JobTypeClaimRecord).First(&taskStatus)
	if taskStatus.Status != 1 {
		return
	}
	services.GetClaimRecord()
}
