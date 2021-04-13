package main

import (
    "database/sql"
    "context"

    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=example dbname=postgres sslmode=disable")
    defer db.Close()

    if err != nil {
        panic(err)
    }

    ctx := context.Background()
    if err := db.PingContext(ctx); err != nil {
        panic(err)
    }
}