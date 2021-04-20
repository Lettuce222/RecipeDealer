package usecase

import "github.com/Lettuce222/RecipeDealer/entity"

type IngredientRepository interface {
	Create(ingredient entity.Ingredient) error
	Update(ingredient entity.Ingredient) error
	Delete(identifier uint) error

	Show() ([]*entity.Ingredient, error)
	Find(identifier uint) (*entity.Ingredient, error)
}
