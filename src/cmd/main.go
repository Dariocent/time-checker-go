package main

import (
	"log"
	"os"

	"github.com/Dariocent/time-checker-go/cmd/api"
	"github.com/Dariocent/time-checker-go/db"

	_ "github.com/golang-migrate/migrate/v4/source/file" // Import the file source driver
	_ "github.com/lib/pq"                                // PostgreSQL driver
)

func main() {
	// Database connection
	dataSourceName := "postgres://postgres:password@localhost:5432?sslmode=disable"
	db := db.NewDB(dataSourceName)

	APIServer := api.NewAPIServer(os.Getenv("ADDR"), db)
	err := APIServer.Run()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
