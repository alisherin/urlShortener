package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func GetDb() (*sql.DB, error) {
	connStr := "postgres://" + os.Getenv("DB_USERNAME")
	connStr += ":" + os.Getenv("DB_PASSWORD")
	connStr += "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	connStr += "?sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conn, nil
}

func CheckConnection() {
	conn, err := GetDb()
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Connected to DB")
}
