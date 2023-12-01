package main

import (
    "context"
    "database/sql"
    "log"
    "testEntGo/ent"

    "entgo.io/ent/dialect"
    entsql "entgo.io/ent/dialect/sql"
    _ "github.com/jackc/pgx/v5/stdlib"
    // _ "github.com/lib/pq"
)

var (
	pgDriver = "pgx" //This is the name of the postgres driver registered by jackc/pgx
)

func getClient(dbURL string) (*ent.Client, error) {
    // Open Database
    db, err := sql.Open(pgDriver, dbURL)
    if err != nil {
        return nil, err
    }

    // Create driver and return
    driver := entsql.OpenDB(dialect.Postgres, db)
    return ent.NewClient(ent.Driver(driver)), nil
}

func main() {
    // client, err := ent.Open(pgDriver, "host=localhost port=5432 user=testuser dbname=test_entgo password=testpswd")
    client, err := getClient("host=localhost port=5432 user=testuser dbname=test_entgo password=testpswd")
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
