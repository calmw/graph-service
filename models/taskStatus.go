package models

type TaskStatus struct {
	Id       int   `gorm:"column:id;primaryKey"`
	TaskType int64 `json:"task_type" gorm:"column:task_type"`
	Status   int64 `json:"status" gorm:"column:status"`
}
