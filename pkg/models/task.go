package models

import "time"

type TaskResult struct {
	Result []Task `json:"result"`
	Status string `json:"status"`
	Time   string `json:"time"`
}

type Task struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   string    `json:"payload"`
	Priority  int       `json:"priority"`
	Locked    bool      `json:"locked"`
}
