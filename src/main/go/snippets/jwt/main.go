package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

var db *sql.DB

func main() {

	pgUrl, err := pq.ParseURL("postgres://wrrlykkg:lR-a3bWALvwx1C6NIv5HgrWtjjdr7VuI@manny.db.elephantsql.com:5432/wrrlykkg")
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

	r := mux.NewRouter()
	r.HandleFunc("/signup", signup).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Problem Decode User")
	}

	spew.Dump(user)
	if user.Email == "" {
		error := Error{"Email not provided"}
		reposndWithError(w, http.StatusBadRequest, error)
		return
	}
	if user.Password == "" {
		error := Error{"Password not provided"}
		reposndWithError(w, http.StatusBadRequest, error)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal("Problem during hashing password")
	}
	user.Password = string(hash)

	fmt.Printf("Inserting user into database\n")
	statement := "insert into users (email, password) values($1, $2) RETURNING id;"
	err = db.QueryRow(statement, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		log.Fatalln(err)
		error := Error{"Server error"}
		reposndWithError(w, http.StatusBadRequest, error)
		return
	}
	fmt.Printf("Inserted user %v into database\n", user.ID)
	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	responseJSON(w, user)
}

func reposndWithError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup login")
	var user User
	var error Error

	json.NewDecoder(r.Body).Decode(&user)
	spew.Dump(user)

	if user.Email == "" {
		error = Error{"Email not provided"}
		reposndWithError(w, http.StatusBadRequest, error)
		return
	}
	if user.Password == "" {
		error = Error{"Password not provided"}
		reposndWithError(w, http.StatusBadRequest, error)
		return
	}
	password := user.Password

	row := db.QueryRow("select * from users where email=$1", user.Email)
	if row != nil {
		fmt.Printf("User exist in database\n")
	}
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = "The user does not exist"
			reposndWithError(w, http.StatusBadRequest, error)
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
		reposndWithError(w, http.StatusBadRequest, error)
		return
	}

	tok, err := GenerateToken(user)
	if err != nil {
		fmt.Printf("Problem with generated token %v\n", tok)
	}
	fmt.Println(tok)

	jwt2 := JWT{Token: tok}
	responseJSON(w, jwt2)

}
func protected(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protected started")

}
func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint started")

}

func GenerateToken(user User) (string, error) {
	// var err error
	secret := "secret"

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
func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
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
			return []byte("secret"), nil
		})
		if error != nil {
			reposndWithError(w, http.StatusBadRequest, Error{})
			return
		}

		spew.Dump(token) // token is now validated

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			reposndWithError(w, http.StatusBadRequest, Error{})
			return
		}
	})

}
func responseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
