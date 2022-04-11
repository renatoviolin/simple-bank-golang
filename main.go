package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/renatoviolin/simplebank/api"
	db "github.com/renatoviolin/simplebank/db/sqlc"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAdress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect do DB: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("Server cannot start ", err.Error())
		os.Exit(2)
	}
}
