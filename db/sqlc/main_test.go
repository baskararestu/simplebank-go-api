package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/baskararestu/simplebank-go-api/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:root@localhost:5436/db_simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(dbDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
