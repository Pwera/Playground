package controller

import (
	"github.com/pwera/Playground/src/main/go/snippets/_new/ddd/domain"
	"net/http"
)

type UserController struct{
	BaseController
	UserService domain.UserService
}

func (c UserController) List(w http.ResponseWriter, r *http.Request){
	users, err := c.UserService.Users()
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.MarshalAndWriteHeaders(users, w)
}

func (c UserController) Show(w http.ResponseWriter, r *http.Request){}
func (c UserController) Create(w http.ResponseWriter, r *http.Request){}
func (c UserController) Delete(w http.ResponseWriter, r *http.Request){}
