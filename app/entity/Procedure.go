package entity

type Procedure struct {
	ID     int
	Action string
	Inputs []Food
	Output Food
}
