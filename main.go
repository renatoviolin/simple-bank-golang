package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/renatoviolin/simplebank/api"
	db "github.com/renatoviolin/simplebank/db/sqlc"
	"github.com/renatoviolin/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect do DB: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot start server")
	}

	err = server.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("Server cannot start ", err.Error())
		os.Exit(2)
	}
}
