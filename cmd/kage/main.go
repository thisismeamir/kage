package main

import (
	"github.com/thisismeamir/kage/internal/server"
	"log"
)

func main() {
	log.Println("Starting Kage...")

	srv := server.New()
	if err := srv.Start(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
