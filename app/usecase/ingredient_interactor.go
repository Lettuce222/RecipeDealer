package usecase

import "github.com/Lettuce222/RecipeDealer/entity"

type IngredientInteractor struct {
	IngredientRepository IngredientRepository
}

func (interactor *IngredientInteractor) Add(ingredient entity.Ingredient) (err error) {
	err = interactor.IngredientRepository.Create(ingredient)
	return
}

func (interactor *IngredientInteractor) Update(ingredient entity.Ingredient) (err error) {
	err = interactor.IngredientRepository.Update(ingredient)
	return
}

func (interactor *IngredientInteractor) Remove(ingredient entity.Ingredient) (err error) {
	err = interactor.IngredientRepository.Delete(ingredient.ID)
	return
}

func (interactor *IngredientInteractor) FindAll() (ingredients []*entity.Ingredient, err error) {
	ingredients, err = interactor.IngredientRepository.Show()
	return
}

func (interactor *IngredientInteractor) FindById(identifier uint) (ingredient *entity.Ingredient, err error) {
	ingredient, err = interactor.IngredientRepository.Find(identifier)
	return
}
