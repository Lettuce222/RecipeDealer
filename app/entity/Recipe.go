package entity

type Recipe struct {
	ID uint
	RecipeBody
}

type RecipeBody struct {
	Name       string
	Procedures []Procedure
}
