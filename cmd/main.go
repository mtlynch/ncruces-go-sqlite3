package main

import (
	"fmt"
	"log"

	"github.com/ncruces/go-sqlite3"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func main() {
	db, err := sqlite3.Open(":memory:")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id int, name varchar(10))`)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Exec(`INSERT INTO users(id, name) VALUES(0, 'go'), (1, 'zig'), (2, 'whatever')`)
	if err != nil {
		log.Fatal(err)
	}

	stmt, _, err := db.Prepare(`SELECT id, name FROM users`)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if row, err := stmt.Step(); err != nil {
			log.Fatal(err)
		} else if !row {
			break
		}
		fmt.Println(stmt.ColumnInt(0), stmt.ColumnText(1))
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
