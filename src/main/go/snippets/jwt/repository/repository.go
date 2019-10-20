package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Pwera/Playground/src/main/go/snippets/jwt/model"
)

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}
func InsertUser(db *sql.DB, user model.User) model.User {
	statement := "insert into users (email, password) values($1, $2) RETURNING id;"
	err := db.QueryRow(statement, user.Email, user.Password).Scan(&user.ID)
	logFatal(err)
	return user
}

func FindUserByEmail(db *sql.DB, user model.User) (model.User, error) {
	row := db.QueryRow("select * from users where email=$1", user.Email)
	if row != nil {
		fmt.Printf("User exist in database\n")
	}
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}
