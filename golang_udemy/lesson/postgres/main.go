package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	//C
	cmd := "INSERT INTO persons (name, age) VALUES ($1, $2)"
	_, err := Db.Exec(cmd, "Nancy", 20)
	if err != nil {
		log.Fatalln(err)
	}
}
