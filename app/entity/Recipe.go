package entity

type Recipe struct {
	ID ID
	RecipeBody
}

type RecipeBody struct {
	Name       string
	Procedures []Procedure
}
