package usecase

import "github.com/Lettuce222/RecipeDealer/entity"

type IngredientRepository interface {
	Create(*entity.Ingredient) error
	Update(*entity.Ingredient) error
	Delete(uint) error

	Show() ([]*entity.Ingredient, error)
	Find(uint) (*entity.Ingredient, error)
}
