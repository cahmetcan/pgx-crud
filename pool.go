package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"

	"github.com/joho/godotenv"

	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println(os.Getenv("DB"))

	startConnection := time.Now()
	databaseUrl := "postgresql://postgres:123123@0.0.0.0:5432/postgres"

	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to database")
	fmt.Println("Time to connect:", time.Since(startConnection))

	randomId := strconv.Itoa(rand.Intn(100))
	startQuery := time.Now()

	fmt.Print("Inserting ", randomId, " into test table... ")
	query(dbPool, fmt.Sprintf(`
			insert into test (id) values (%s)
		`, randomId))
	fmt.Println("Time to query:", time.Since(startQuery))

	defer dbPool.Close()
	fmt.Println("Closed database connection")
}

// test_table
/*
express üstünde

2 3



*/
