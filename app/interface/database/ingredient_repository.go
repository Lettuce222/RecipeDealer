package database

import (
	"github.com/Lettuce222/RecipeDealer/entity"
)

type IngredientRepository struct {
	DbHandler DbHandler
}

var ingredientTableName = "ingredient_records"

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
	err := repo.DbHandler.Delete(identifier, ingredientTableName)
	return err
}

func (repo *IngredientRepository) Show() ([]*entity.Ingredient, error) {
	rows, err := repo.DbHandler.Show(ingredientTableName, "id, name, number, unit")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredients := []*entity.Ingredient{}
	for rows.Next() {
		var id uint
		var name string
		var number float32
		var unit string

		err = rows.Scan(&id, &name, &number, &unit)
		if err != nil {
			return nil, err
		}

		ingredients = append(
			ingredients,
			entity.NewIngredient(id, name, number, unit),
		)
	}

	return ingredients, nil
}

func (repo *IngredientRepository) Find(identifier uint) (*entity.Ingredient, error) {
	row := repo.DbHandler.Find(identifier, ingredientTableName, "id, name, number, unit")

	var id uint
	var name string
	var number float32
	var unit string

	err := row.Scan(&id, &name, &number, &unit)
	if err != nil {
		return nil, err
	}

	ingredient := entity.NewIngredient(id, name, number, unit)

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
