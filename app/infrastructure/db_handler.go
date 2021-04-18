package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Lettuce222/RecipeDealer/interface/database"
)

type DbHandler struct {
	Db *gorm.DB
}

func NewDbHandler() *DbHandler {
	dsn := "host=postgres port=5432 user=postgres password=example dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &DbHandler{Db: db}
}

func (hander *DbHandler) Create(value interface{}, args ...interface{}) (*database.Result, error) {
	result := hander.Db.Create(&value)
	return &database.Result{Value: value, Count: result.RowsAffected}, result.Error
}
