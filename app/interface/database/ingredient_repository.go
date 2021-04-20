package database

import (
	"reflect"

	"github.com/Lettuce222/RecipeDealer/entity"
)

type IngredientRepository struct {
	DbHandler DbHandler
}

func (repo *IngredientRepository) Create(ingredient entity.Ingredient) error {
	err := repo.DbHandler.Create(
		NewIngredientRecord(ingredient.Name, ingredient.Number, ingredient.Unit),
	)
	return err
}

func (repo *IngredientRepository) Update(ingredient entity.Ingredient) error {
	err := repo.DbHandler.Update(
		ingredient.ID,
		NewIngredientRecord(ingredient.Name, ingredient.Number, ingredient.Unit),
		[]string{"Name", "Number", "Unit"},
	)
	return err
}

func (repo *IngredientRepository) Delete(identifier uint) error {
	err := repo.DbHandler.Delete(identifier, reflect.TypeOf(IngredientRecord{}))
	return err
}

func (repo *IngredientRepository) Show() ([]*entity.Ingredient, error) {
	ingredientRecords := []IngredientRecord{}
	_, err := repo.DbHandler.Show(reflect.TypeOf(IngredientRecord{}))
	if err != nil {
		return nil, err
	}

	ingredients := []*entity.Ingredient{}
	for _, ingredientRecord := range ingredientRecords {
		ingredients = append(
			ingredients,
			entity.NewIngredient(ingredientRecord.ID, ingredientRecord.Name, ingredientRecord.Number, ingredientRecord.Unit),
		)
	}
	return ingredients, nil
}

func (repo *IngredientRepository) Find(identifier uint) (*entity.Ingredient, error) {
	ingredientRecord := IngredientRecord{}
	_, err := repo.DbHandler.Find(identifier, reflect.TypeOf(IngredientRecord{}))
	if err != nil {
		return nil, err
	}

	ingredient := entity.NewIngredient(ingredientRecord.ID, ingredientRecord.Name, ingredientRecord.Number, ingredientRecord.Unit)

	return ingredient, nil
}

type IngredientRecord struct {
	Model
	entity.IngredientBody
}

func NewIngredientRecord(Name string, Number float32, Unit string) *IngredientRecord {
	return &IngredientRecord{
		IngredientBody: entity.IngredientBody{
			Name: Name,
			Quantity: entity.Quantity{
				Number: Number,
				Unit:   Unit,
			}}}
}
