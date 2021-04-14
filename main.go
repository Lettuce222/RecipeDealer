package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"

    "github.com/gin-gonic/gin"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func main() {

    dsn := "host=postgres port=5432 user=postgres password=example dbname=postgres sslmode=disable"
    _, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(err)
    }
    
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":5000") 
}