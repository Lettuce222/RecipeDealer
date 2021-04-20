package entity

type Ingredient struct {
	ID uint
	IngredientBody
}

type IngredientBody struct {
	Name string
	Quantity
}

func NewIngredient(ID uint, Name string, Number float32, Unit string) *Ingredient {
	return &Ingredient{
		ID,
		IngredientBody{
			Name,
			Quantity{
				Number,
				Unit,
			},
		},
	}
}
