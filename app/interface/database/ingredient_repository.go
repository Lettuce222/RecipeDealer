package database

import (
	"reflect"

	"github.com/Lettuce222/RecipeDealer/entity"
	"gorm.io/gorm"
)

type IngredientRepository struct {
	DbHandler DbHandler
}

func (repo *IngredientRepository) Create(ingredient entity.Ingredient) error {
	err := repo.DbHandler.Create(
		NewGormIngredient(ingredient.Name, ingredient.Number, ingredient.Unit),
	)
	return err
}

func (repo *IngredientRepository) Update(identifier uint, ingredient entity.Ingredient) error {
	err := repo.DbHandler.Update(
		identifier,
		NewGormIngredient(ingredient.Name, ingredient.Number, ingredient.Unit),
		[]string{"Name", "Number", "Unit"},
	)
	return err
}

func (repo *IngredientRepository) Delete(identifier uint) error {
	err := repo.DbHandler.Delete(identifier, reflect.TypeOf(GormIngredient{}))
	return err
}

func (repo *IngredientRepository) Show() ([]*entity.Ingredient, error) {
	rows, err := repo.DbHandler.Show(reflect.TypeOf(GormIngredient{}))
	if err != nil {
		return nil, err
	}

	gormIngredients := []GormIngredient{}
	rows.Scan(&gormIngredients)

	ingredients := []*entity.Ingredient{}
	for _, gormIngredient := range gormIngredients {
		ingredients = append(
			ingredients,
			entity.NewIngredient(gormIngredient.ID, gormIngredient.Name, gormIngredient.Number, gormIngredient.Unit),
		)
	}
	return ingredients, nil
}

func (repo *IngredientRepository) Find(identifier uint) (*entity.Ingredient, error) {
	row, err := repo.DbHandler.Find(identifier, reflect.TypeOf(GormIngredient{}))
	if err != nil {
		return nil, err
	}

	gormIngredient := GormIngredient{}
	row.Scan(&gormIngredient)

	ingredient := entity.NewIngredient(gormIngredient.ID, gormIngredient.Name, gormIngredient.Number, gormIngredient.Unit)

	return ingredient, nil
}

type GormIngredient struct {
	gorm.Model
	entity.IngredientBody
}

func NewGormIngredient(Name string, Number float32, Unit string) *GormIngredient {
	return &GormIngredient{
		IngredientBody: entity.IngredientBody{
			Name,
			entity.Quantity{
				Number,
				Unit,
			}}}
}
