package entity

type Procedure struct {
	ID uint
	ProcedureBody
}

type ProcedureBody struct {
	Action string
	Inputs []Food
	Output Food
}
