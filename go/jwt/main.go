package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/pwera/jwt/controller"
	"github.com/pwera/jwt/driver"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
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
