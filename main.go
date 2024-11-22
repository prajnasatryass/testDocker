package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {
	db, err := sql.Open("postgres",
		"host=host.docker.internal port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		if err := db.Ping(); err != nil {
			fmt.Fprintf(writer, "Error: %s\n", err.Error())
			return
		}
		fmt.Fprintf(writer, "Pong!\n")
		log.Printf("Ping took %dms\n", time.Since(start).Milliseconds())
	})

	port := "8080"
	log.Println("Listening on :" + port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
