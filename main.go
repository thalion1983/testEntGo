package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testEntGo/ent"
    "testEntGo/ent/people"

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

// creates a test record of people
func createPeople(ctx context.Context, cli *ent.Client) (*ent.People, error) {
    res, err := cli.People.Create().
        SetName("Primera").
        SetLastName("Persona").
        SetAge(23).
        Save(ctx)

    if err != nil {
        return nil, fmt.Errorf("while creating new people: %w", err)
    }
    log.Printf("record in people created")
    return res, nil
}

// queries user, fails if not found
func queryPeople(ctx context.Context, cli *ent.Client) (*ent.People, error) {
    res, err := cli.People.Query().
        Where(people.Name("Primera")).
        Only(ctx)
    if err != nil {
        return nil, fmt.Errorf("while querying people: %w", err)
    }
    log.Printf("people queried")
    return res, nil
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

    ctx := context.Background() // context in use

    // Create a record on people
    if _, err := createPeople(ctx, client); err != nil {
        log.Fatalf("failed creating new people: %v", err)
    }
    // Query people
    if _, err := queryPeople(ctx, client); err != nil {
        log.Fatalf("failed querying people: %v", err)
    }
}
