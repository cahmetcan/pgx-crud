package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	// docker run --name postgre -p 5432:5432 -e POSTGRES_PASSWORD=123123 -d postgres
	host     = "0.0.0.0:5432"
	port     = 5432
	user     = "postgres"
	password = "123123"
	dbname   = "postgres"
)

func pgConnect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
