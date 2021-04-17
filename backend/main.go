package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	recipe "github.com/Lettuce222/RecipeDealer/models"
	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "host=postgres port=5432 user=postgres password=example dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&recipe.Recipe{}, &recipe.Ingredient{}, &recipe.Cut{}, &recipe.Burn{}, &recipe.Mix{}, &recipe.Steam{})

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello!!!!!",
		})
	})
	err = r.Run(":5000")

	if err != nil {
		panic(err)
	}

}
