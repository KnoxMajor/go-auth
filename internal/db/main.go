package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	_, er := os.Stat(".env")
	if os.IsNotExist(er) {
		log.Fatal(".env file doesnt exist")
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("db: Error loading .env file")
	} else {
		log.Print("Loaded env file for db connection")
	}

	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	connStr := "user=" + user + " dbname=" + dbName + " sslmode=" + sslMode
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	} else {
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Connected to DB")
	}

	return db
}
