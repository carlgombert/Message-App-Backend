package main

import (
	"log"
	"message-app-backend/web-service-gin/db"
)

func main() {
	_, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("couldn't init db connection: %s", err)
	}
}
