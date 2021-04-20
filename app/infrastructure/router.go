package infrastructure

import (
	"github.com/Lettuce222/RecipeDealer/interface/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	controller := controller.NewIngredientController(NewDbHandler())

	router.POST("/users", func(c *gin.Context) {
		controller.Create(c)
	})
	router.GET("/users", func(c *gin.Context) {
		controller.Index(c)
	})
	router.GET("/users/:id", func(c *gin.Context) {
		controller.Find(c)
	})

	router.Run(":5000")
}
