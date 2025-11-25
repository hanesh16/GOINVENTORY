package main

import (
	"database/sql"
	"fmt"
	"goinventory/api"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dataSourceName := "root:Honey@007@tcp(127.0.0.1:3306)/inventorydb?parseTime=true"

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	fmt.Println("Connected to MySQL (inventorydb)")

	mux := http.NewServeMux()
	api.RegisterRoutes(mux, db, nil)

	fmt.Println("Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("server error:", err)
	}
}
