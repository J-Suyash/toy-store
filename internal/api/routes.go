package api

import (
	"github.com/gorilla/mux"
	"github.com/yourusername/toy-store-go/internal/api/handlers"
	"github.com/yourusername/toy-store-go/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(db *mongo.Database) *mux.Router {
	r := mux.NewRouter()

	imageService := services.NewImageService()
	toyHandler := &handlers.ToyHandler{
		DB:           db,
		ImageService: imageService,
	}

	r.HandleFunc("/api/toys", toyHandler.GetAllToys).Methods("GET")
	r.HandleFunc("/api/toys/{id}", toyHandler.GetToyByID).Methods("GET")
	r.HandleFunc("/api/toys", toyHandler.CreateToy).Methods("POST")
	r.HandleFunc("/api/toys/{id}", toyHandler.UpdateToy).Methods("PUT")
	r.HandleFunc("/api/toys/{id}", toyHandler.DeleteToy).Methods("DELETE")

	return r
}
