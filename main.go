package main

import (
    "context"
    "log"
    "testEntGo/ent"
    // _ "github.com/jackc/pgx/v4/stdlib"
    _ "github.com/lib/pq"
)

var (
	pgDriver = "postgres" // "pgx" //This is the name of the postgres driver registered by jackc/pgx
)

func main() {
    client, err := ent.Open(pgDriver, "host=localhost port=5432 user=testuser dbname=test_entgo password=testpswd")
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
