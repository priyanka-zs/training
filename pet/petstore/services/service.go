package services

import (
	"context"
	"errors"
	"petstore/models"

	"github.com/google/uuid"
)

type Service struct {
	store Store
}

func New(s Store) Service {
	return Service{store: s}
}

// ServicePost takes a row from store layer and returns it to handler layer
func (s Service) ServicePost(ctx context.Context, pet models.Pet) (models.Pet, error) {
	if pet.Status != "Available" && pet.Status != "Pending" && pet.Status != "Sold" {
		return models.Pet{}, errors.New("invalid status")
	}

	if pet.Species == "Bird" {
		if pet.Age != 1 {
			return models.Pet{}, errors.New("invalid age for Bird")
		}
	}

	if pet.Species == "Dog" {
		if pet.Age < 1 || pet.Age > 4 {
			return models.Pet{}, errors.New("invalid age for Dog")
		}
	}

	if pet.Species == "Cat" {
		if pet.Age < 1 || pet.Age > 3 {
			return models.Pet{}, errors.New("invalid age for Cat")
		}
	}

	post, err := s.store.Post(ctx, pet)
	if err != nil {
		return models.Pet{}, err
	}

	return post, nil
}

// ServiceGetByID takes id as input and returns a row to the handler layer
func (s Service) ServiceGetByID(ctx context.Context, id uuid.UUID) (models.Pet, error) {
	pet, err := s.store.GetByID(ctx, id)
	if err != nil {
		return models.Pet{}, errors.New("error from store layer")
	}

	return pet, nil
}

func (s Service) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
