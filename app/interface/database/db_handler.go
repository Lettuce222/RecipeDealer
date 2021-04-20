package database

import (
	"reflect"

	"gorm.io/gorm"
)

type DbHandler interface {
	Create(value interface{}) error
	Update(identifier uint, value interface{}, columns []string) error
	Delete(identifier uint, model reflect.Type) error

	Show(model reflect.Type) (*gorm.DB, error)
	Find(identifier uint, model reflect.Type) (*gorm.DB, error)
}
