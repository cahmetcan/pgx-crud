package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func query(
	dbPool *pgxpool.Pool,
	query string,
) {
	rows, err := dbPool.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Query successful")
	for rows.Next() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(rows.Values())
	}
}
