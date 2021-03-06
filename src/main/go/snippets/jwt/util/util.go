package util

import (
	"encoding/json"
	"net/http"

	"github.com/pwera/jwt/model"
)

func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func ReposndWithError(w http.ResponseWriter, status int, error model.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}
