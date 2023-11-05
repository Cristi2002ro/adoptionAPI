package dal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "roundhouse.proxy.rlwy.net"
	port     = 27626
	user     = "postgres"
	password = "CB3d2-f-a-aCb-GBdBa154cefe66c-C3"
	dbname   = "railway"
)

var db *sql.DB

func init() { //function called init will always be executed by go
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
	return db
}
