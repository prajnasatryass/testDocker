package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	db, err := sql.Open("postgres",
		"host=db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		panic(err)
	}
}
