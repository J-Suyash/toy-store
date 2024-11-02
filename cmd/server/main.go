package main

import (
	"log"
	"net/http"
	"os"

	"github.com/J-Suyash/toy-store/internal/api"
	"github.com/J-Suyash/toy-store/internal/config"
	"github.com/J-Suyash/toy-store/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.New()

	db, err := database.ConnectMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect()

	router := api.SetupRoutes(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
