package usecase

import "github.com/Lettuce222/RecipeDealer/entity"

type ProcedureInteractor struct {
	ProcedureRepository ProcedureRepository
}

func (interactor *ProcedureInteractor) Add(ingredient *entity.Procedure) (err error) {
	err = interactor.ProcedureRepository.Create(ingredient)
	return
}

func (interactor *ProcedureInteractor) Update(ingredient *entity.Procedure) (err error) {
	err = interactor.ProcedureRepository.Update(ingredient)
	return
}

func (interactor *ProcedureInteractor) Remove(ingredient *entity.Procedure) (err error) {
	err = interactor.ProcedureRepository.Delete(ingredient.ID)
	return
}

func (interactor *ProcedureInteractor) FindAll() (ingredients []*entity.Procedure, err error) {
	ingredients, err = interactor.ProcedureRepository.Show()
	return
}

func (interactor *ProcedureInteractor) FindById(id uint) (ingredient *entity.Procedure, err error) {
	ingredient, err = interactor.ProcedureRepository.Find(id)
	return
}
