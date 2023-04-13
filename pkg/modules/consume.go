package modules

import (
	"encoding/json"

	"github.com/Biskwit/qo-q/pkg/models"
	"github.com/surrealdb/surrealdb.go"
)

func ExecuteTask(client *surrealdb.DB) []models.Task {
	data, err := client.Query("SELECT * FROM (SELECT * FROM task WHERE locked = FALSE ORDER BY timestamp DESC LIMIT 10) ORDER BY priority DESC", nil)
	if err != nil {
		panic(err)
	}
	var tasks []models.TaskResult
	jsonByte, error := json.Marshal(data)
	if error != nil {
		panic(error)
	}
	json.Unmarshal(jsonByte, &tasks)

	for _, row := range tasks[0].Result {
		println("payload: " + row.Timestamp.Format("2006-01-02 15:04:05"))
		// goroutine to execute task
		client.Change(row.ID, map[string]interface{}{
			"locked": true,
		})
	}
	return tasks[0].Result
}
