package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Replace DSN with ones own credentials
	dsn := ""

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB open error:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	fmt.Println("Connected to MySQL (inventorydb)")
}
