package database

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func Connect(connectionString string) {
    var err error
    db, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("Error pinging the database: %v", err)
    }
}

func GetDB() *sql.DB {
    return db
}

func Close() {
    if err := db.Close(); err != nil {
        log.Fatalf("Error closing the database: %v", err)
    }
}