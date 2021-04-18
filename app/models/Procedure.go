package recipe

import (
	"gorm.io/gorm"
)

type Procedure struct {
	gorm.Model
	Action   string
	Inputs   []Food `gorm:"foreignKey:FromProcedureID"`
	Output   Food   `gorm:"foreignKey:ToProcedureID"`
	RecipeID uint
}
