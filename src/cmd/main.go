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
	// dataSourceName := "postgres://postgres:password@localhost:5432?sslmode=disable"
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_url := os.Getenv("DB_URL")
	db_name := os.Getenv("DB_NAME")
	dataSourceName := "postgres://" + db_user + ":" + db_password + "@" + db_url + "/" + db_name + "?sslmode=disable"

	db := db.NewDB(dataSourceName)

	APIServer := api.NewAPIServer(os.Getenv("ADDR"), db)
	err := APIServer.Run()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
