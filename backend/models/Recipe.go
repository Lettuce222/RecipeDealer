package recipe

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Procedures  []Procedure
}