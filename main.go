package main

import (
	"database/sql"
	"log"
	"net/http"
	"nimblestack/database"
	"nimblestack/router"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()
	queries := database.New(db)

	jwtSercet := os.Getenv("API_TOKEN")
	route := router.NewRouter(queries, []byte(jwtSercet))

	log.Println("NimbleStack server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", route.Handler()))
}
