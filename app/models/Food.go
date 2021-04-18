package recipe

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	FromProcedureID uint
	ToProcedureID   uint
}
