package infrastructure

import (
	"reflect"

	"github.com/Lettuce222/RecipeDealer/interface/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	db.AutoMigrate(&database.IngredientRecord{})

	return &DbHandler{Db: db}
}

func (hander *DbHandler) Create(value interface{}) error {
	result := hander.Db.Create(value)
	return result.Error
}

func (hander *DbHandler) Update(identifier uint, value interface{}, columns []string) error {
	result := hander.Db.Model(&value).Where("ID = ?", identifier).Updates(value)
	return result.Error
}

func (hander *DbHandler) Delete(identifier uint, model reflect.Type) error {
	result := hander.Db.Delete(model, identifier)
	return result.Error
}

func (hander *DbHandler) Show(model reflect.Type) (reflect.Value, error) {
	records := reflect.MakeSlice(reflect.SliceOf(model), 1, 1)
	result := hander.Db.Model(reflect.New(model)).Find(&records)
	return records, result.Error
}

func (hander *DbHandler) Find(identifier uint, model reflect.Type) (reflect.Value, error) {
	record := reflect.New(model)
	result := hander.Db.First(record, identifier)
	return record, result.Error
}
