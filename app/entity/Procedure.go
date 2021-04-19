package entity

type Procedure struct {
	ID ID
	ProcedureBody
}

type ProcedureBody struct {
	Action string
	Inputs []Food
	Output Food
}
