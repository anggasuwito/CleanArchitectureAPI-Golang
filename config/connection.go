package config

import (
	"database/sql"
	"log"

	//driver for mysql
	_ "github.com/go-sql-driver/mysql"
)

//Connection Connection
func Connection() *sql.DB {
	db, _ := sql.Open("mysql", "root:angga123@tcp(localhost:3306)/enigmaschool_api")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
