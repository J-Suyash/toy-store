package validators

import (
	"errors"

	"github.com/J-Suyash/toy-store/internal/models"
)

func ValidateToy(toy *models.Toy) error {
	if toy.Name == "" {
		return errors.New("name is required")
	}
	if toy.Description == "" {
		return errors.New("description is required")
	}
	if toy.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if toy.Quantity < 0 {
		return errors.New("quantity must be non-negative")
	}
	return nil
}
