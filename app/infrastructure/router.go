package infrastructure

import (
	"github.com/Lettuce222/RecipeDealer/interface/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	controller := controller.NewIngredientController(NewDbHandler())

	router.POST("/ingredients", func(c *gin.Context) {
		controller.Create(c)
	})
	router.GET("/ingredients", func(c *gin.Context) {
		controller.Index(c)
	})
	router.GET("/ingredients/:id", func(c *gin.Context) {
		controller.Find(c)
	})
	router.PUT("/ingredients", func(c *gin.Context) {
		controller.Update(c)
	})
	router.DELETE("/ingredients/:id", func(c *gin.Context) {
		controller.Remove(c)
	})

	router.Run(":5000")
}
