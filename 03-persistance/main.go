package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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
	/*
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
	*/

	//Inserting a new user
	/*
		username := "testuser"
		password := "testpassword"
		created_at := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, created_at)
		if err != nil {
			log.Fatalln(err)
		}
		id, err := result.LastInsertId()
		fmt.Printf("User with id %d is created\n", id)
	*/

	//Selecting a user
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)
	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err = db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id, username, password, createdAt)

}
