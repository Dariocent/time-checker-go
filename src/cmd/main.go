package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Dariocent/time-checker-go/cmd/api"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	fmt.Println("The current time is: ", time.Now())

	APIServer := api.NewAPIServer(os.Getenv("ADDR"), nil)
	err := APIServer.Run()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
