package main

import (
	"database/sql"
	"log"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM usuarios WHERE id > ?", 2)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		log.Println(u)
	}
}
