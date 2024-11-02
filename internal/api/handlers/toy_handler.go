package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/J-Suyash/toy-store/internal/models"
	"github.com/J-Suyash/toy-store/internal/services"
	"github.com/J-Suyash/toy-store/internal/validators"
	"github.com/J-Suyash/toy-store/pkg/apierror"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ToyHandler struct {
	DB           *mongo.Database
	ImageService *services.ImageService
}

func (h *ToyHandler) GetAllToys(w http.ResponseWriter, r *http.Request) {
	toys, err := models.GetAllToys(h.DB)
	if err != nil {
		apierror.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch toys")
		return
	}

	json.NewEncoder(w).Encode(toys)
}

func (h *ToyHandler) GetToyByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	toy, err := models.GetToyByID(h.DB, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			apierror.RespondWithError(w, http.StatusNotFound, "Toy not found")
		} else {
			apierror.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch toy")
		}
		return
	}

	json.NewEncoder(w).Encode(toy)
}

func (h *ToyHandler) CreateToy(w http.ResponseWriter, r *http.Request) {
	var toy models.Toy
	if err := json.NewDecoder(r.Body).Decode(&toy); err != nil {
		apierror.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validators.ValidateToy(&toy); err != nil {
		apierror.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	thumbnail, err := h.ImageService.GetRandomImage(toy.Name)
	if err != nil {
		apierror.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch thumbnail")
		return
	}
	toy.Thumbnail = thumbnail

	if err := toy.Create(h.DB); err != nil {
		apierror.RespondWithError(w, http.StatusInternalServerError, "Failed to create toy")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(toy)
}

func (h *ToyHandler) UpdateToy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var toy models.Toy
	if err := json.NewDecoder(r.Body).Decode(&toy); err != nil {
		apierror.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validators.ValidateToy(&toy); err != nil {
		apierror.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	toy.ID = id
	if err := toy.Update(h.DB); err != nil {
		apierror.RespondWithError(w, http.StatusInternalServerError, "Failed to update toy")
		return
	}

	json.NewEncoder(w).Encode(toy)
}

func (h *ToyHandler) DeleteToy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	if err := models.DeleteToy(h.DB, id); err != nil {
		apierror.RespondWithError(w, http.StatusInternalServerError, "Failed to delete toy")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
