package modules

import (
	"time"

	"github.com/Biskwit/qo-q/pkg/models"
	"github.com/surrealdb/surrealdb.go"
)

func NewTask(client *surrealdb.DB, timestamp time.Time, payload string) {
	client.Create("task", models.Task{
		Timestamp: timestamp,
		Payload:   payload,
	})
}
