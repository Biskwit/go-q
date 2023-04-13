package main

import (
	"time"

	"github.com/Biskwit/qo-q/pkg/db"
	"github.com/Biskwit/qo-q/pkg/modules"
	"github.com/surrealdb/surrealdb.go"
)

func main() {
	var client *surrealdb.DB = db.SurrealDb()
	modules.NewTask(client, time.Now(), "Hello, world")
	modules.ExecuteTask(client)
}
