package infrastructure

import (
	"reflect"

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

func (hander *DbHandler) Show(model reflect.Type) (*gorm.DB, error) {
	result := hander.Db.Find(model)
	return result, result.Error
}

func (hander *DbHandler) Find(identifier uint, model reflect.Type) (*gorm.DB, error) {
	result := hander.Db.First(model, identifier)
	return result, result.Error
}
