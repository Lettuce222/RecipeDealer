package recipe

import (
	"gorm.io/gorm"
)

type Procedure struct {
	gorm.Model
	Motion  Motion
	Targets []Ingredient
	RecipeID uint
}

type Motion string