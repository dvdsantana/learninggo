package main

import (
	"log"

	"github.com/dvdsantana/learninggo/go-microservices/internal/database"
	"github.com/dvdsantana/learninggo/go-microservices/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatal("Error creating database client:", err)
		return
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
