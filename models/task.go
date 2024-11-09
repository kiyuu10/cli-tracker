package models

import (
	"time"
)

type TaskStatus int8

const (
	TaskStatusTodo       TaskStatus = 1
	TaskStatusInProgress TaskStatus = 2
	TaskStatusDone       TaskStatus = 3
)

type (
	TaskList struct {
		Tasks     []TaskInfo `json:"tasks"`
		CurrentId int32      `json:"current_id"`
	}

	TaskInfo struct {
		ID          int32      `json:"id"`
		Description string     `json:"description"`
		Status      TaskStatus `json:"status"`
		CreateAt    time.Time  `json:"create_at"`
		UpdateAt    time.Time  `json:"update_at"`
	}
)
