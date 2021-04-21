package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Lettuce222/RecipeDealer/entity"
	"github.com/Lettuce222/RecipeDealer/interface/database"
	"github.com/Lettuce222/RecipeDealer/usecase"
)

type IngredientController struct {
	Interactor usecase.IngredientInteractor
}

func NewIngredientController(dbHandler database.DbHandler) *IngredientController {
	return &IngredientController{
		Interactor: usecase.IngredientInteractor{
			IngredientRepository: &database.IngredientRepository{
				DbHandler: dbHandler,
			},
		},
	}
}

func (controller *IngredientController) Create(c Context) {
	ingredient := &entity.Ingredient{}
	c.Bind(ingredient)
	err := controller.Interactor.Add(ingredient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("Failed to Create Ingredient: msg... %w", err))
		return
	}

	c.JSON(http.StatusCreated, fmt.Sprintf("Please Specify ID: %s is Created", ingredient.Name))
}

func (controller *IngredientController) Index(c Context) {
	ingredients, err := controller.Interactor.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("Failed to Index Ingredients: msg... %w", err))
		return
	}

	c.JSON(http.StatusOK, ingredients)
}

func (controller *IngredientController) Find(c Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Errorf("Please Specify ID: msg... %w", err))
		return
	}

	ingredient, err := controller.Interactor.FindById(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("Failed to Find Ingredient: msg... %w", err))
	}

	c.JSON(200, ingredient)
}

func (controller *IngredientController) Update(c Context) {
	ingredient := &entity.Ingredient{}
	c.Bind(ingredient)
	err := controller.Interactor.Update(ingredient)

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("Failed to Update Ingredient: msg... %w", err))
	}

	c.JSON(200, ingredient)
}

func (controller *IngredientController) Remove(c Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Errorf("Please Specify ID: msg... %w", err))
		return
	}

	ingredient, err := controller.Interactor.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("Failed to Find Ingredient: msg... %w", err))
	}

	err = controller.Interactor.Remove(ingredient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("Failed to Remove Ingredient: msg... %w", err))
	}

	c.JSON(200, ingredient)
}
