package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dbinfo := fmt.Sprintf("dbname=%s sslmode=disable", "golang")
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("id  |    name  |   email |")
	for rows.Next() {
		var id int
		var name string
		var email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%3v | %8v | %6v\n", id, name, email)

	}
}
