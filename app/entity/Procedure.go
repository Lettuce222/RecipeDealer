package entity

type Procedure struct {
	Id     int
	Action string
	Inputs []Food
	Output Food
}
