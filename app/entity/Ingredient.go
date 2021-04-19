package entity

type Ingredient struct {
	ID uint
	IngredientBody
}

type IngredientBody struct {
	Name string
	Quantity
}
