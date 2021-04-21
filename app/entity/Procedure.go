package entity

type Procedure struct {
	ID uint
	ProcedureBody
}

type ProcedureBody struct {
	Action          string
	Ingredients     []Ingredient
	InputProcedures []Procedure
}
