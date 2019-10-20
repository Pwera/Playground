package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func ConnectDBB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Problem with postgres url")
	}
	db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal("Problem with connect to db")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Problem db ping")
	}
	return db
}
