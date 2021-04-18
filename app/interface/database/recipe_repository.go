package database

import "github.com/Lettuce222/RecipeDealer/entity"

type RecipeRepository struct {
	DbHandler DbHandler
}

func (repo *RecipeRepository) Create(recipe []entity.Recipe) (int, error) {
	result, err := repo.DbHandler.Create(recipe)
}

type Result struct {
	Value interface{}
	Count int64
}
