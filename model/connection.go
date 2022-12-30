package model

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Connection est une structure de données pour la connexion à la base de données
type Connection struct {
	DB *sql.DB
}

var ModelInstance Connection

// Connect établit une connexion à la base de données et renvoie une Connection
func Connect() Connection {
	var (
		DB_DRIVER   string = os.Getenv("DB_DRIVER")
		DB_USER     string = os.Getenv("DB_USER")
		DB_PASSWORD string = os.Getenv("DB_PASSWORD")
		DB_NAME     string = os.Getenv("DB_NAME")
		DB_HOST     string = os.Getenv("DB_HOST")
		DB_PORT     string = os.Getenv("DB_PORT")
	)

	connectionString := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?parseTime=true"

	db, err := sql.Open(DB_DRIVER, connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return Connection{db}
}
