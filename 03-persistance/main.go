package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:rootuser@(127.0.0.1:3306)/todos?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	//Pinging the db
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	//Creating a new table
	query := `
		CREATE TABLE users (
			id INT NOT NULL AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("table created")
}
