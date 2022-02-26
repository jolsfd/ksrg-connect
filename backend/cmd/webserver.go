package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func StartServer() {
	var err error

	// Open database
	DB, err = sql.Open("sqlite3", PathDB)
	checkError(err)
	defer DB.Close()

	// Check connection
	err = DB.Ping()
	checkError(err)

	// Create table
	err = CreateTable()
	checkError(err)

	// Webserver
	router := Router()

	// Enable cors for security
	handler := cors.New(cors.Options{AllowedOrigins: []string{Origins}, AllowCredentials: true}).Handler(router)

	// Start server
	fmt.Printf("Server is running on Port%s. End the server with [CTRL + C]\n", PORT)

	log.Fatal(http.ListenAndServe(PORT, handler))
}
