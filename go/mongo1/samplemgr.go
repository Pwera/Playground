package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pwera/mongo/service"

	"github.com/jmoiron/sqlx"
	"github.com/urfave/negroni"
)

func initDb() (err error) {
	db, err := sqlx.Connect("sqlite3", "_database.db")
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Duration(30) * time.Minute)

	return
}

func main() {

	// mux router
	router := service.InitRouter()
	// init web server
	server := negroni.Classic()
	server.UseHandler(router)

	port := 33505
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	server.Run(addr)
}
