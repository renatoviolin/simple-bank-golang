package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/renatoviolin/simplebank/api"
	db "github.com/renatoviolin/simplebank/db/sqlc"
	"github.com/renatoviolin/simplebank/util"
)

func main() {
	start_mode := os.Getenv("ENV_MODE") // PROD or DEV

	config, err := util.LoadConfig(".", start_mode)
	if err != nil {
		log.Fatal("cannot load config: ", err)
		return
	}

	fmt.Printf("\n")
	fmt.Printf("DB URI........: %s\n", config.DBSource)
	fmt.Printf("DB DRIVER.....: %s\n", config.DBDriver)
	fmt.Printf("SERVER ADDRESS: %s\n", config.ServerAdress)
	fmt.Printf("TOKEN SYMETRIC: %s\n", config.TokenSymmetricKey)
	fmt.Printf("TOKEN DURATION: %s\n", config.AccessTokenDuration)
	fmt.Printf("\n")

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
