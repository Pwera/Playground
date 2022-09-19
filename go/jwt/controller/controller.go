package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/pwera/jwt/model"
	"github.com/pwera/jwt/repository"
	"github.com/pwera/jwt/util"

	"golang.org/x/crypto/bcrypt"
)

type Controller struct{}

func (c Controller) Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Fatal("Problem Decode User")
		}

		spew.Dump(user)
		if user.Email == "" {
			util.ReposndWithError(w, http.StatusBadRequest, model.Error{"Email not provided"})
			return
		}
		if user.Password == "" {
			util.ReposndWithError(w, http.StatusBadRequest, model.Error{"Password not provided"})
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			log.Fatal("Problem during hashing password")
		}
		user.Password = string(hash)

		fmt.Printf("Inserting user into database\n")
		// statement := "insert into users (email, password) values($1, $2) RETURNING id;"
		// err = db.QueryRow(statement, user.Email, user.Password).Scan(&user.ID)
		repository.InsertUser(db, user)

		if err != nil {
			log.Fatalln(err)
			util.ReposndWithError(w, http.StatusBadRequest, model.Error{"Server error"})
			return
		}
		fmt.Printf("Inserted user %v into database\n", user.ID)
		user.Password = ""
		w.Header().Set("Content-Type", "application/json")
		util.ResponseJSON(w, user)
	}
}
func (c Controller) Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("signup login")
		var user model.User
		var error model.Error

		json.NewDecoder(r.Body).Decode(&user)
		spew.Dump(user)

		if user.Email == "" {
			util.ReposndWithError(w, http.StatusBadRequest, model.Error{"Email not provided"})
			return
		}
		if user.Password == "" {
			util.ReposndWithError(w, http.StatusBadRequest, model.Error{"Password not provided"})
			return
		}
		password := user.Password

		user, err := repository.FindUserByEmail(db, user)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "The user does not exist"
				util.ReposndWithError(w, http.StatusBadRequest, error)
				return
			} else {
				log.Fatal(err)
			}
		}

		spew.Dump(user)

		hashedPassword := user.Password

		// Compare password provided by user with database
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			passwordNotValidMessage := "Password is not valid"
			fmt.Println(passwordNotValidMessage)
			error.Message = passwordNotValidMessage
			util.ReposndWithError(w, http.StatusBadRequest, error)
			return
		}

		tok, err := GenerateToken(user)
		if err != nil {
			fmt.Printf("Problem with generated token %v\n", tok)
		}
		fmt.Println(tok)

		jwt2 := model.JWT{Token: tok}
		util.ResponseJSON(w, jwt2)

	}
}

func protected(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protected started")

}

func (c Controller) ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint started")

}

func GenerateToken(user model.User) (string, error) {
	// var err error
	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})
	spew.Dump(token)
	tokeStrin, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	return tokeStrin, nil
}

/* Validate token from client to the server
and give access to protected endpoint */
func (c Controller) TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var errorObjct Error
		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		// fmt.Println("authHeader")
		token, error := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			spew.Dump(token) //token is not valid
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if error != nil {
			util.ReposndWithError(w, http.StatusBadRequest, model.Error{})
			return
		}

		spew.Dump(token) // token is now validated

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			util.ReposndWithError(w, http.StatusBadRequest, model.Error{})
			return
		}
	})

}
