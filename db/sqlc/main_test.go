package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/renatoviolin/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, errConfig := util.LoadConfig("../../", "DEV")
	if errConfig != nil {
		log.Fatal("cannot load config: ", errConfig)
		return
	}

	var err error

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect do DB: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
