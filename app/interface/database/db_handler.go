package database

import (
	"database/sql"
	"reflect"
	"time"
)

type DbHandler interface {
	Create(value interface{}) error
	Update(identifier uint, value interface{}, columns []string) error
	Delete(identifier uint, model reflect.Type) error

	Show(model reflect.Type) (reflect.Value, error)
	Find(identifier uint, model reflect.Type) (reflect.Value, error)
}

type Model struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
