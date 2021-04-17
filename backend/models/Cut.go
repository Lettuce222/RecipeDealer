package recipe

type Cut struct {
	*Procedure
	Kind Kind
	Quantity
}

type Kind string
