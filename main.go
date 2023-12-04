package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testEntGo/ent"
	"testEntGo/ent/clothe"
    "testEntGo/ent/group"
	"testEntGo/ent/people"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	// _ "github.com/lib/pq"
)

var (
	pgDriver = "pgx" //This is the name of the postgres driver registered by jackc/pgx
    clotheData = []struct{
        clotheType string
        color string
        date time.Time
    }{
        {
            clotheType: "tshirt",
            color: "white",
            date: time.Date(2021,02,14,11,0,0,0, time.UTC),
        },
        {
            clotheType: "pants",
            color: "blue",
            date: time.Now(),
        },
    }
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

// creates clothes and a new people who owns them
func createClothe(ctx context.Context, cli *ent.Client) (*ent.People, error) {
    var clothes []*ent.Clothe
    for _, v := range clotheData {
        c, err := cli.Clothe.Create().
            SetType(v.clotheType).
            SetColor(v.color).
            SetBuyDate(v.date).
            Save(ctx)
        if err != nil {
            return nil, fmt.Errorf("while creating clothe %s: %w", v.clotheType, err)
        }
        log.Printf("clothe %s created", v.clotheType)
        clothes = append(clothes, c)
    }

    res, err := cli.People.Create().
        SetName("Otra").
        SetLastName("Gente").
        SetAge(23).
        AddClothes(clothes...).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("while creating new people: %w", err)
    }
    log.Printf("new people created")
    return res, nil
}

// creates a new clothe and adds it to person 1
func otherClothe(ctx context.Context, cli *ent.Client) error {
    c, err := cli.Clothe.Create().
        SetType("hat").
        SetColor("brown").
        SetBuyDate(time.Now()).
        Save(ctx)
    if err != nil {
        return fmt.Errorf("while creating third clothe: %w", err)
    }
    log.Printf("new clothe created")

    p, err := cli.People.Query().
        Where(people.ID(1)).
        Only(ctx)
    if err != nil {
        return fmt.Errorf("while getting people id=1: %w", err)
    }

    if _, err := p.Update().AddClothes(c).Save(ctx); err != nil {
        return fmt.Errorf("while updating people id=1: %w", err)
    }
    log.Printf("new clothe added to people id=1")
    return nil
}

// query clothes of a people record
func queryClothes(ctx context.Context, pe *ent.People) error{
    clothes, err := pe.QueryClothes().All(ctx)
    if err != nil {
        return fmt.Errorf("while querying clothes from people id=2: %w", err)
    }

    log.Println("clothes: ", clothes)

    // query just one clothe by color
    clotheType, err := pe.QueryClothes().Where(clothe.Color("white")).Only(ctx)
    if err != nil {
        return fmt.Errorf("while querying blue clothe: %w", err)
    }
    log.Println("clothe: ", clotheType)
    return nil
}

// query clothes from a people and owner for each clothe
func queryClotheOwner(ctx context.Context, cli *ent.Client) error {
    clothes, err := cli.Clothe.Query().All(ctx)
    if err != nil {
        return fmt.Errorf("while getting clothes to query owner: %w", err)
    }

    for _, v := range clothes {
        owner, err := v.QueryOwner().Only(ctx)
        if err != nil {
            fmt.Errorf("while querying owner for %q: %w", v.Type, err)
        }
        log.Printf("owner of %q: %q %q", v.Type, owner.Name, owner.LastName)
    }
    return nil
}

// create groups
func createGroups(ctx context.Context, cli *ent.Client) error {
    peoples, err := cli.People.Query().All(ctx)
    if err != nil {
        return fmt.Errorf("while getting peoples to create groups: %w", err)
    }

    if err := cli.Group.Create().SetName("unused").Exec(ctx); err != nil {
        return fmt.Errorf("while creating unused group: %w", err)
    }

    if err := cli.Group.Create().SetName("students").AddPeoples(peoples...).Exec(ctx); err != nil {
        return fmt.Errorf("while creating students group: %w", err)
    }

    if err := cli.Group.Create().SetName("artists").AddPeoples(peoples[0]).Exec(ctx); err != nil {
        return fmt.Errorf("while creating artists group: %w", err)
    }

    if err := cli.Group.Create().SetName("scientists").AddPeoples(peoples[1]).Exec(ctx); err != nil {
        return fmt.Errorf("while creating scientists group: %w", err)
    }
    log.Printf("groups created")
    return nil
}

// get clothes owned by peoples on a certain group
func queryGroups(ctx context.Context, cli *ent.Client) error {
    // get clothes of artists
    clothes, err := cli.Group.Query(). //Initializes query
        Where(group.Name("artists")). //Condition on groups
        QueryPeoples(). // querying peoples on previous result
        QueryClothes(). // querying clothes on previuos result
        All(ctx)
    if err != nil {
        return fmt.Errorf("while querying clothes of artists: %w", err)
    }

    log.Println("clothes of artists: ", clothes)
    return nil
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

    // Create records of clothes
    pe, err := createClothe(ctx, client)
    if err != nil {
        log.Fatalf("failed creating clothes: %v", err)
    }
    // Create other clothe
    if err := otherClothe(ctx, client); err != nil {
        log.Fatalf("failed creating clothes: %v", err)
    }
    // query clothes from a people record
    if err := queryClothes(ctx, pe); err != nil {
        log.Fatalf("failed querying clothes: %v", err)
    }

    // querying clothe owner
    if err := queryClotheOwner(ctx, client); err != nil {
        log.Fatal("failed querying owner: %v", err)
    }

    // creating groups
    if err := createGroups(ctx, client); err != nil {
        log.Fatal("failed creating groups: %v", err)
    }

    // getting clothes of artists
    if err := queryGroups(ctx, client); err != nil {
        log.Fatal("failed querying groups: %v", err)
    }
}
