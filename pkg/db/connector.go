package db

import (
	"github.com/surrealdb/surrealdb.go"
)

func SurrealDb() *surrealdb.DB {
	db, err := surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(map[string]string{
		"user": "root",
		"pass": "root",
	}); err != nil {
		panic(err)
	}
	// Select namespace and database
	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}

	return db
}
