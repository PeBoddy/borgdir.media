package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	"log"
)

// Db handle
var Db *sql.DB
var err error

func InitSQLiteDB() {

	Db, err = sql.Open("sqlite3", "./data/borgdirmedia")
	if err != nil {
		panic(err)
	}
}

func InitPostgresDB() {
	connStr := "user=borgdirmedia dbname=borgdirmedia password=borgdirmedia host=localhost port=5431 sslmode=disable"
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)

	}
}
