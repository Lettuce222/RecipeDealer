package recipe

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	Quantity
	ProcedureID uint
}
