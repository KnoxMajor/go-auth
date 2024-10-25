package user

import (
	"database/sql"
	"log"

	"errors"

	"github.com/knoxmajor/go-auth/config"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID    int
	Email string
}

func Signup(email string, password string) error {
	var foundUser User

	tx, err := config.Database.Begin()
	if err != nil {
		log.Println(err)
		return errors.New("Internal Server Error")
	}

	findStmt, _ := tx.Prepare("SELECT * FROM users WHERE email = $1")
	defer findStmt.Close()
	err = findStmt.QueryRow(email).Scan(&foundUser)
	if err != nil && err != sql.ErrNoRows {
		log.Println("A user with that email already exists")
		return errors.New("A user with that email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return errors.New("Internal Server Error")
	}

	stmt, err := tx.Prepare("INSERT INTO users(email, password) VALUES($1, $2)")
	if err != nil {
		log.Println(err)
		return errors.New("Internal Server Error")
	}

	defer stmt.Close()
	_, err = stmt.Exec(email, hashedPassword)
	err = tx.Commit()
	return err
}
