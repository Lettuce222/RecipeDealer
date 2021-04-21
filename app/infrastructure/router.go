package infrastructure

import (
	"github.com/Lettuce222/RecipeDealer/interface/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	// Ingredient
	ingredientController := controller.NewIngredientController(NewDbHandler())
	router.POST("/ingredients", func(c *gin.Context) {
		ingredientController.Create(c)
	})
	router.GET("/ingredients", func(c *gin.Context) {
		ingredientController.Index(c)
	})
	router.GET("/ingredients/:id", func(c *gin.Context) {
		ingredientController.Find(c)
	})
	router.PUT("/ingredients", func(c *gin.Context) {
		ingredientController.Update(c)
	})
	router.DELETE("/ingredients/:id", func(c *gin.Context) {
		ingredientController.Remove(c)
	})

	// Procedure
	procedureController := controller.NewProcedureController(NewDbHandler())
	router.POST("/procedures", func(c *gin.Context) {
		procedureController.Create(c)
	})

	router.Run(":5000")
}
