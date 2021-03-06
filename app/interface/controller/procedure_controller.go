package controller

import (
	"fmt"
	"net/http"

	"github.com/Lettuce222/RecipeDealer/entity"
	"github.com/Lettuce222/RecipeDealer/interface/database"
	"github.com/Lettuce222/RecipeDealer/usecase"
)

type ProcedureController struct {
	Interactor usecase.ProcedureInteractor
}

func NewProcedureController(dbHandler database.DbHandler) *ProcedureController {
	return &ProcedureController{
		Interactor: usecase.ProcedureInteractor{
			ProcedureRepository: &database.ProcedureRepository{
				DbHandler: dbHandler,
			},
		},
	}
}

func (controller *ProcedureController) Create(c Context) {
	procedure := &entity.Procedure{}
	err := c.Bind(procedure)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("Failed to Create Procedure: msg... %w", err))
		return
	}

	err = controller.Interactor.Add(procedure)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("Failed to Create Ingredient: msg... %w", err))
		return
	}

	c.JSON(http.StatusCreated, fmt.Sprintf("Procedure(%s) is Created", procedure.Action))
}
