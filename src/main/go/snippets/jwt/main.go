package main

import (
	"database/sql"
	"log"
	"net/http"

	controller "github.com/pwera/jwt/controller"
	driver "github.com/pwera/jwt/driver"

	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {
	gotenv.Load()
}
func main() {
	db = driver.ConnectDBB()
	controller := controller.Controller{}
	r := mux.NewRouter()
	r.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	r.HandleFunc("/login", controller.Login(db)).Methods("POST")
	r.HandleFunc("/protected", controller.TokenVerifyMiddleWare(controller.ProtectedEndpoint)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
