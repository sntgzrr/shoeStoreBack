package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	usr      = "postgres"
	psw      = "root"
	host     = "localhost"
	port     = "5432"
	database = "store"
)

func GetConnection() *sql.DB {
	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", usr, psw, host, port, database)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
