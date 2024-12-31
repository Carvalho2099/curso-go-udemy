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

	// Update
	stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	stmt.Exec("Uóxiton Istive", 1)
	stmt.Exec("Sheristone Uasleska", 2)

	// Delete
	stmt2, _ := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	stmt2.Exec(3)
}
