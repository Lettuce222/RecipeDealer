package usecase

import "github.com/Lettuce222/RecipeDealer/entity"

type ProcedureRepository interface {
	Create(*entity.Procedure) error
	Update(*entity.Procedure) error
	Delete(uint) error

	Show() ([]*entity.Procedure, error)
	Find(uint) (*entity.Procedure, error)
}
