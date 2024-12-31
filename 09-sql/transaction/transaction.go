package main

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")
	stmt.Exec(2000, "Maria")
	stmt.Exec(2001, "João")

	_, err = stmt.Exec(1, "Pedro") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
