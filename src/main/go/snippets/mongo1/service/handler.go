package service

import (
	"encoding/json"
	"github.com/Pwera/Playground/src/main/go/snippets/mongo1/db"
	"net/http"
	"time"
)

type Pingrepsonse struct {
	Message string `json:"message"`
}

func pingHandler(res http.ResponseWriter, req *http.Request) {
	response := Pingrepsonse{Message: "pong"}
	status := http.StatusOK
	respBytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
		status = http.StatusInternalServerError
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(respBytes)
	sweat := db.Sweat{CreatedAt:time.Now()}
	sweat.Create()

}


//db.createUser({user:"admin", pwd:"pass123", roles: [{role: "readWrite", db: "sweatdb"}]})