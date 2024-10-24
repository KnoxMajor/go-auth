package user

import (
	"database/sql"
	"log"

	"errors"
	"github.com/knoxmajor/go-auth/config"
	"golang.org/x/crypto/bcrypt"
)

func Signup(email string, password string) (sql.Result, error) {
	tx, err := config.Database.Begin()
	if err != nil {
		log.Println(err)
		return nil, errors.New("Internal Server Error")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Internal Server Error")
	}

	stmt, err := tx.Prepare("INSERT INTO users(email, password) VALUES($1, $2)")
	if err != nil {
		log.Println(err)
		return nil, errors.New("Internal Server Error")
	}

	defer stmt.Close()
	result, err := stmt.Exec(email, hashedPassword)
	err = tx.Commit()
	return result, err
}
