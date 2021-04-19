package entity

type Procedure struct {
	ID     ID
	Action string
	Inputs []Food
	Output Food
}
