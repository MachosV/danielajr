package storage

import (
	"database/sql"
	"log"
	//initializing the driver
	_ "github.com/mattn/go-sqlite3"
)

var dbConnection *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Println(err)
		log.Fatal("Database connection failed")
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		log.Fatal("Database ping failed")
	}
	dbConnection = db
}

/*
GetDb returns a db connection
*/
func GetDb() *sql.DB {
	return dbConnection
}
