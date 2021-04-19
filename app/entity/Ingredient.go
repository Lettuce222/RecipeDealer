package entity

type Ingredient struct {
	ID ID
	IngredientBody
}

type IngredientBody struct {
	Name string
	Quantity
}
