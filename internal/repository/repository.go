package repository

import "github.com/reaper47/recipya/internal/models"

// Repository is the database repository
type Repository interface {
	GetAllRecipes() ([]models.Recipe, error)
}